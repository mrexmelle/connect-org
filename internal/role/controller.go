package role

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/mrexmelle/connect-org/internal/config"
	"github.com/mrexmelle/connect-org/internal/dto/dtorespwithdata"
	"github.com/mrexmelle/connect-org/internal/dto/dtorespwithoutdata"
	"github.com/mrexmelle/connect-org/internal/localerror"
)

type Controller struct {
	ConfigService     *config.Service
	LocalErrorService *localerror.Service
	RoleService       *Service
}

func NewController(cfg *config.Service, les *localerror.Service, svc *Service) *Controller {
	return &Controller{
		ConfigService:     cfg,
		LocalErrorService: les,
		RoleService:       svc,
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
	info := c.LocalErrorService.Map(err)
	dtorespwithdata.New[Entity](
		data,
		info.ServiceErrorCode,
		info.ServiceErrorMessage,
	).RenderTo(w, info.HttpStatusCode)
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
		dtorespwithdata.NewError(
			localerror.ErrBadJson.Error(),
			err.Error(),
		).RenderTo(w, http.StatusBadRequest)
		return
	}

	data, err := c.RoleService.Create(requestBody)
	info := c.LocalErrorService.Map(err)
	dtorespwithdata.New[Entity](
		data,
		info.ServiceErrorCode,
		info.ServiceErrorMessage,
	).RenderTo(w, info.HttpStatusCode)
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
		dtorespwithoutdata.New(
			localerror.ErrBadJson.Error(),
			err.Error(),
		).RenderTo(w, http.StatusBadRequest)
		return
	}
	err = c.RoleService.UpdateById(requestBody.Fields, chi.URLParam(r, "id"))
	info := c.LocalErrorService.Map(err)
	dtorespwithoutdata.New(
		info.ServiceErrorCode,
		info.ServiceErrorMessage,
	).RenderTo(w, info.HttpStatusCode)
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
	info := c.LocalErrorService.Map(err)
	dtorespwithoutdata.New(
		info.ServiceErrorCode,
		info.ServiceErrorMessage,
	).RenderTo(w, info.HttpStatusCode)
}
