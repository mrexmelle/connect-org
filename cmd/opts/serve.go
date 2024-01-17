package opts

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/mrexmelle/connect-orgs/internal/config"
	"github.com/mrexmelle/connect-orgs/internal/localerror"
	"github.com/mrexmelle/connect-orgs/internal/organization"
	"github.com/mrexmelle/connect-orgs/internal/placement"
	"github.com/mrexmelle/connect-orgs/internal/role"
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
	container.Provide(organization.NewRepository)
	container.Provide(role.NewRepository)
	container.Provide(placement.NewRepository)

	container.Provide(config.NewService)
	container.Provide(localerror.NewService)
	container.Provide(organization.NewService)
	container.Provide(role.NewService)
	container.Provide(placement.NewService)

	container.Provide(organization.NewController)
	container.Provide(role.NewController)
	container.Provide(placement.NewController)

	process := func(
		configService *config.Service,
		organizationController *organization.Controller,
		roleController *role.Controller,
		placementController *placement.Controller,
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

		r.Route("/organizations", func(r chi.Router) {
			r.Post("/", organizationController.Post)
			r.Get("/{id}", organizationController.Get)
			r.Patch("/{id}", organizationController.Patch)
			r.Delete("/{id}", organizationController.Delete)
			r.Get("/{id}/children", organizationController.GetChildren)
			r.Get("/{id}/lineage", organizationController.GetLineage)
			r.Get("/{id}/officers", organizationController.GetOfficers)
			r.Get("/{id}/siblings-and-ancestral-siblings", organizationController.
				GetSiblingsAndAncestralSiblings)
		})

		r.Route("/roles", func(r chi.Router) {
			r.Post("/", roleController.Post)
			r.Get("/{id}", roleController.Get)
			r.Patch("/{id}", roleController.Patch)
			r.Delete("/{id}", roleController.Delete)
		})

		r.Route("/placements", func(r chi.Router) {
			r.Post("/", placementController.Post)
			r.Get("/{id}", placementController.Get)
			r.Patch("/{id}", placementController.Patch)
			r.Delete("/{id}", placementController.Delete)
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
	Short: "Start connect-orgs server",
	Run:   Serve,
}
