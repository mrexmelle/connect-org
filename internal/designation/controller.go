package designation

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/mrexmelle/connect-org/internal/config"
	"github.com/mrexmelle/connect-org/internal/dto/dtobuilderwithdata"
	"github.com/mrexmelle/connect-org/internal/dto/dtobuilderwithoutdata"
	"github.com/mrexmelle/connect-org/internal/localerror"
)

type Controller struct {
	ConfigService      *config.Service
	DesignationService *Service
}

func NewController(cfg *config.Service, svc *Service) *Controller {
	return &Controller{
		ConfigService:      cfg,
		DesignationService: svc,
	}
}

// Get Designations : HTTP endpoint to get designations
// @Tags Designations
// @Description Get a designation
// @Produce json
// @Param id path string true "Designation ID"
// @Success 200 {object} GetResponseDto "Success Response"
// @Failure 400 "BadRequest"
// @Failure 500 "InternalServerError"
// @Router /designations/{id} [GET]
func (c *Controller) Get(w http.ResponseWriter, r *http.Request) {
	data, err := c.DesignationService.RetrieveById(
		chi.URLParam(r, "id"),
	)
	dtobuilderwithdata.New[Entity](data, err).RenderTo(w)
}

// Post Designations : HTTP endpoint to post new designations
// @Tags Designations
// @Description Post a new designations
// @Accept json
// @Produce json
// @Param data body PostRequestDto true "Designation Request"
// @Success 200 {object} PostResponseDto "Success Response"
// @Failure 400 "BadRequest"
// @Failure 500 "InternalServerError"
// @Router /designations [POST]
func (c *Controller) Post(w http.ResponseWriter, r *http.Request) {
	var requestBody PostRequestDto
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		dtobuilderwithdata.New[Entity](nil, localerror.ErrBadJson).RenderTo(w)
		return
	}

	data, err := c.DesignationService.Create(requestBody)
	dtobuilderwithdata.New[Entity](data, err).RenderTo(w)
}

// Patch Designations : HTTP endpoint to patch a designation
// @Tags Designations
// @Description Patch a designation
// @Accept json
// @Produce json
// @Param id path string true "Designation ID"
// @Param data body PatchRequestDto true "Designation Patch Request"
// @Success 200 {object} PatchResponseDto "Success Response"
// @Failure 400 "BadRequest"
// @Failure 500 "InternalServerError"
// @Router /designations/{id} [PATCH]
func (c *Controller) Patch(w http.ResponseWriter, r *http.Request) {
	var requestBody PatchRequestDto
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		dtobuilderwithoutdata.New(localerror.ErrBadJson).RenderTo(w)
		return
	}
	err = c.DesignationService.UpdateById(requestBody.Fields, chi.URLParam(r, "id"))
	dtobuilderwithoutdata.New(err).RenderTo(w)
}

// Delete Designations : HTTP endpoint to delete designations
// @Tags Designations
// @Description Delete a designation
// @Produce json
// @Param id path string true "Designation ID"
// @Success 200 {object} DeleteResponseDto "Success Response"
// @Failure 400 "BadRequest"
// @Failure 500 "InternalServerError"
// @Router /designations/{id} [DELETE]
func (c *Controller) Delete(w http.ResponseWriter, r *http.Request) {
	err := c.DesignationService.DeleteById(chi.URLParam(r, "id"))
	dtobuilderwithoutdata.New(err).RenderTo(w)
}
