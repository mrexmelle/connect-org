package organization

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
	ConfigService       *config.Service
	OrganizationService *Service
}

func NewController(cfg *config.Service, svc *Service) *Controller {
	return &Controller{
		ConfigService:       cfg,
		OrganizationService: svc,
	}
}

// Get Organizations : HTTP endpoint to get an organization
// @Tags Organizations
// @Description Get an organization
// @Produce json
// @Param id path string true "Organization ID"
// @Success 200 {object} GetResponseDto "Success Response"
// @Failure 400 "BadRequest"
// @Failure 500 "InternalServerError"
// @Router /organizations/{id} [GET]
func (c *Controller) Get(w http.ResponseWriter, r *http.Request) {
	response, err := c.OrganizationService.RetrieveById(
		chi.URLParam(r, "id"),
	)
	dtobuilderwithdata.New[Entity](response, err).RenderTo(w)
}

// Post Organizations : HTTP endpoint to post new organizations
// @Tags Organizations
// @Description Post a new organization
// @Accept json
// @Produce json
// @Param data body PostRequestDto true "Organization Request"
// @Success 200 {object} PostResponseDto "Success Response"
// @Failure 400 "BadRequest"
// @Failure 500 "InternalServerError"
// @Router /organizations [POST]
func (c *Controller) Post(w http.ResponseWriter, r *http.Request) {
	var requestBody PostRequestDto
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		dtobuilderwithdata.New[Entity](nil, localerror.ErrParsingJson).RenderTo(w)
		return
	}

	response, err := c.OrganizationService.Create(requestBody)
	dtobuilderwithdata.New[Entity](response, err).RenderTo(w)
}

// Delete Organizations : HTTP endpoint to delete organizations
// @Tags Organizations
// @Description Delete an organization
// @Produce json
// @Param id path string true "Organization ID"
// @Success 200 {object} DeleteResponseDto "Success Response"
// @Failure 400 "BadRequest"
// @Failure 500 "InternalServerError"
// @Router /organizations/{id} [DELETE]
func (c *Controller) Delete(w http.ResponseWriter, r *http.Request) {
	err := c.OrganizationService.DeleteById(chi.URLParam(r, "id"))
	dtobuilderwithoutdata.New(err).RenderTo(w)
}

// Get Children of Organizations : HTTP endpoint to get the children of an organization
// @Tags Organizations
// @Description Get children of an organization
// @Produce json
// @Param id path string true "Organization ID"
// @Success 200 {object} GetChildrenResponseDto "Success Response"
// @Failure 400 "BadRequest"
// @Failure 500 "InternalServerError"
// @Router /organizations/{id}/children [GET]
func (c *Controller) GetChildren(w http.ResponseWriter, r *http.Request) {
	response, err := c.OrganizationService.RetrieveChildrenById(
		chi.URLParam(r, "id"),
	)
	dtobuilderwithdata.New[[]Entity](&response, err).RenderTo(w)
}
