package role

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/mrexmelle/connect-orgs/internal/config"
	"github.com/mrexmelle/connect-orgs/internal/dto/dtobuilderwithdata"
	"github.com/mrexmelle/connect-orgs/internal/dto/dtobuilderwithoutdata"
	"github.com/mrexmelle/connect-orgs/internal/localerror"
)

type Controller struct {
	ConfigService *config.Service
	RoleService   *Service
}

func NewController(cfg *config.Service, svc *Service) *Controller {
	return &Controller{
		ConfigService: cfg,
		RoleService:   svc,
	}
}

// Get Roles : HTTP endpoint to get a role
// @Tags Roles
// @Description Get a role
// @Produce json
// @Param id path string true "Role ID"
// @Success 200 {object} GetResponseDto "Success Response"
// @Failure 400 "BadRequest"
// @Failure 500 "InternalServerError"
// @Router /roles/{id} [GET]
func (c *Controller) Get(w http.ResponseWriter, r *http.Request) {
	data, err := c.RoleService.RetrieveById(
		chi.URLParam(r, "id"),
	)
	dtobuilderwithdata.New[Entity](data, err).RenderTo(w)
}

// Post Roles : HTTP endpoint to post new roles
// @Tags Roles
// @Description Post a new role
// @Accept json
// @Produce json
// @Param data body PostRequestDto true "Role Request"
// @Success 200 {object} PostResponseDto "Success Response"
// @Failure 400 "BadRequest"
// @Failure 500 "InternalServerError"
// @Router /roles [POST]
func (c *Controller) Post(w http.ResponseWriter, r *http.Request) {
	var requestBody PostRequestDto
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		dtobuilderwithdata.New[Entity](nil, localerror.ErrBadJson).RenderTo(w)
		return
	}

	data, err := c.RoleService.Create(requestBody)
	dtobuilderwithdata.New[Entity](data, err).RenderTo(w)
}

// Patch Roles : HTTP endpoint to patch a role
// @Tags Roles
// @Description Patch a role
// @Accept json
// @Produce json
// @Param id path string true "Role ID"
// @Param data body PatchRequestDto true "Role Patch Request"
// @Success 200 {object} PatchResponseDto "Success Response"
// @Failure 400 "BadRequest"
// @Failure 500 "InternalServerError"
// @Router /roles/{id} [PATCH]
func (c *Controller) Patch(w http.ResponseWriter, r *http.Request) {
	var requestBody PatchRequestDto
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		dtobuilderwithoutdata.New(localerror.ErrBadJson).RenderTo(w)
		return
	}
	err = c.RoleService.UpdateById(requestBody.Fields, chi.URLParam(r, "id"))
	dtobuilderwithoutdata.New(err).RenderTo(w)
}

// Delete Roles : HTTP endpoint to delete roles
// @Tags Roles
// @Description Delete a role
// @Produce json
// @Param id path string true "Role ID"
// @Success 200 {object} DeleteResponseDto "Success Response"
// @Failure 400 "BadRequest"
// @Failure 500 "InternalServerError"
// @Router /roles/{id} [DELETE]
func (c *Controller) Delete(w http.ResponseWriter, r *http.Request) {
	err := c.RoleService.DeleteById(chi.URLParam(r, "id"))
	dtobuilderwithoutdata.New(err).RenderTo(w)
}