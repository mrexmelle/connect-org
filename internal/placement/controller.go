package placement

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
	ConfigService  *config.Service
	OrgRoleService *Service
}

func NewController(cfg *config.Service, svc *Service) *Controller {
	return &Controller{
		ConfigService:  cfg,
		OrgRoleService: svc,
	}
}

// Get Placements : HTTP endpoint to get placements
// @Tags Placements
// @Description Get a placement
// @Produce json
// @Param id path string true "Placement ID"
// @Success 200 {object} GetResponseDto "Success Response"
// @Failure 400 "BadRequest"
// @Failure 500 "InternalServerError"
// @Router /placements/{id} [GET]
func (c *Controller) Get(w http.ResponseWriter, r *http.Request) {
	data, err := c.OrgRoleService.RetrieveById(
		chi.URLParam(r, "id"),
	)
	dtobuilderwithdata.New[Entity](data, err).RenderTo(w)
}

// Post Placements : HTTP endpoint to post new placements
// @Tags Placements
// @Description Post a new placement
// @Accept json
// @Produce json
// @Param data body PostRequestDto true "Placement Request"
// @Success 200 {object} PostResponseDto "Success Response"
// @Failure 400 "BadRequest"
// @Failure 500 "InternalServerError"
// @Router /placements [POST]
func (c *Controller) Post(w http.ResponseWriter, r *http.Request) {
	var requestBody PostRequestDto
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		dtobuilderwithdata.New[Entity](nil, localerror.ErrBadJson).RenderTo(w)
		return
	}

	data, err := c.OrgRoleService.Create(requestBody)
	dtobuilderwithdata.New[Entity](data, err).RenderTo(w)
}

// Patch Placements : HTTP endpoint to patch a placement
// @Tags Placements
// @Description Patch a placement
// @Accept json
// @Produce json
// @Param id path string true "Placement ID"
// @Param data body PatchRequestDto true "Placement Patch Request"
// @Success 200 {object} PatchResponseDto "Success Response"
// @Failure 400 "BadRequest"
// @Failure 500 "InternalServerError"
// @Router /placements/{id} [PATCH]
func (c *Controller) Patch(w http.ResponseWriter, r *http.Request) {
	var requestBody PatchRequestDto
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		dtobuilderwithoutdata.New(localerror.ErrBadJson).RenderTo(w)
		return
	}
	err = c.OrgRoleService.UpdateById(requestBody.Fields, chi.URLParam(r, "id"))
	dtobuilderwithoutdata.New(err).RenderTo(w)
}

// Delete Placements : HTTP endpoint to delete placements
// @Tags Placements
// @Description Delete a placement
// @Produce json
// @Param id path string true "Placement ID"
// @Success 200 {object} DeleteResponseDto "Success Response"
// @Failure 400 "BadRequest"
// @Failure 500 "InternalServerError"
// @Router /placements/{id} [DELETE]
func (c *Controller) Delete(w http.ResponseWriter, r *http.Request) {
	err := c.OrgRoleService.DeleteById(chi.URLParam(r, "id"))
	dtobuilderwithoutdata.New(err).RenderTo(w)
}
