package opts

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/mrexmelle/connect-org/internal/config"
	"github.com/mrexmelle/connect-org/internal/designation"
	"github.com/mrexmelle/connect-org/internal/localerror"
	"github.com/mrexmelle/connect-org/internal/member"
	"github.com/mrexmelle/connect-org/internal/membership"
	"github.com/mrexmelle/connect-org/internal/node"
	"github.com/mrexmelle/connect-org/internal/role"
	"github.com/spf13/cobra"
	httpSwagger "github.com/swaggo/http-swagger"
	"go.uber.org/dig"
)

func EnableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func Serve(cmd *cobra.Command, args []string) {
	container := dig.New()

	container.Provide(config.NewRepository)
	container.Provide(designation.NewRepository)
	container.Provide(membership.NewRepository)
	container.Provide(node.NewRepository)
	container.Provide(role.NewRepository)

	container.Provide(config.NewService)
	container.Provide(localerror.NewService)
	container.Provide(designation.NewService)
	container.Provide(membership.NewService)
	container.Provide(node.NewService)
	container.Provide(role.NewService)

	container.Provide(designation.NewController)
	container.Provide(member.NewController)
	container.Provide(membership.NewController)
	container.Provide(node.NewController)
	container.Provide(role.NewController)

	process := func(
		configService *config.Service,
		designationController *designation.Controller,
		memberController *member.Controller,
		membershipController *membership.Controller,
		nodeController *node.Controller,
		roleController *role.Controller,
	) {
		r := chi.NewRouter()

		r.Use(cors.Handler(cors.Options{
			AllowedOrigins:   []string{"https://*", "http://localhost:3000"},
			AllowedMethods:   []string{"GET", "PATCH", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			AllowCredentials: true,
			MaxAge:           300, // Maximum value not ignored by any of major browsers
		}))

		if configService.GetProfile() == "local" {
			r.Mount("/swagger", httpSwagger.Handler(
				httpSwagger.URL(fmt.Sprintf("http://localhost:%d/swagger/doc.json", configService.GetPort())),
				httpSwagger.UIConfig(map[string]string{
					"defaultModelsExpandDepth": "-1",
				}),
			))
		}

		r.Route("/nodes", func(r chi.Router) {
			r.Post("/", nodeController.Post)
			r.Get("/{id}", nodeController.Get)
			r.Patch("/{id}", nodeController.Patch)
			r.Delete("/{id}", nodeController.Delete)
			r.Get("/{id}/children", nodeController.GetChildren)
			r.Get("/{id}/lineage", nodeController.GetLineage)
			r.Get("/{id}/officers", nodeController.GetOfficers)
			r.Get("/{id}/lineage-siblings", nodeController.GetLineageSiblings)
			r.Get("/{id}/members", nodeController.GetMembers)
		})

		r.Route("/roles", func(r chi.Router) {
			r.Post("/", roleController.Post)
			r.Get("/{id}", roleController.Get)
			r.Patch("/{id}", roleController.Patch)
			r.Delete("/{id}", roleController.Delete)
		})

		r.Route("/designations", func(r chi.Router) {
			r.Post("/", designationController.Post)
			r.Get("/{id}", designationController.Get)
			r.Patch("/{id}", designationController.Patch)
			r.Delete("/{id}", designationController.Delete)
		})

		r.Route("/memberships", func(r chi.Router) {
			r.Post("/", membershipController.Post)
			r.Get("/{id}", membershipController.Get)
			r.Patch("/{id}", membershipController.Patch)
			r.Delete("/{id}", membershipController.Delete)
		})

		r.Route("/members", func(r chi.Router) {
			r.Get("/{ehid}/nodes", memberController.GetNodes)
			r.Get("/{ehid}/history", memberController.GetHistory)
		})

		err := http.ListenAndServe(fmt.Sprintf(":%d", configService.GetPort()), r)

		if err != nil {
			panic(err)
		}
	}

	if err := container.Invoke(process); err != nil {
		panic(err)
	}
}

var ServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start connect-org server",
	Run:   Serve,
}
