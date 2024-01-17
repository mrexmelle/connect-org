package organization

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/mrexmelle/connect-orgs/internal/config"
	"github.com/mrexmelle/connect-orgs/internal/dto/dtobuilderwithdata"
	"github.com/mrexmelle/connect-orgs/internal/dto/dtobuilderwithoutdata"
	"github.com/mrexmelle/connect-orgs/internal/localerror"
	"github.com/mrexmelle/connect-orgs/internal/placement"
	"github.com/mrexmelle/connect-orgs/internal/tree"
)

type Controller struct {
	ConfigService       *config.Service
	OrganizationService *Service
	PlacementService    *placement.Service
}

func NewController(
	cfg *config.Service,
	svc *Service,
	ps *placement.Service,
) *Controller {
	return &Controller{
		ConfigService:       cfg,
		OrganizationService: svc,
		PlacementService:    ps,
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
	data, err := c.OrganizationService.RetrieveById(
		chi.URLParam(r, "id"),
	)
	dtobuilderwithdata.New[Entity](data, err).RenderTo(w)
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
		dtobuilderwithdata.New[Entity](nil, localerror.ErrBadJson).RenderTo(w)
		return
	}

	data, err := c.OrganizationService.Create(requestBody)
	dtobuilderwithdata.New[Entity](data, err).RenderTo(w)
}

// Patch Organizations : HTTP endpoint to patch an organization
// @Tags Organizations
// @Description Patch an organization
// @Accept json
// @Produce json
// @Param id path string true "Organization ID"
// @Param data body PatchRequestDto true "Organization Patch Request"
// @Success 200 {object} PatchResponseDto "Success Response"
// @Failure 400 "BadRequest"
// @Failure 500 "InternalServerError"
// @Router /organizations/{id} [PATCH]
func (c *Controller) Patch(w http.ResponseWriter, r *http.Request) {
	var requestBody PatchRequestDto
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		dtobuilderwithoutdata.New(localerror.ErrBadJson).RenderTo(w)
		return
	}
	err = c.OrganizationService.UpdateById(requestBody.Fields, chi.URLParam(r, "id"))
	dtobuilderwithoutdata.New(err).RenderTo(w)
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
	data, err := c.OrganizationService.RetrieveChildrenById(
		chi.URLParam(r, "id"),
	)
	dtobuilderwithdata.New[[]Entity](&data, err).RenderTo(w)
}

// Get Lineage of Organizations : HTTP endpoint to get the lineage of an organization
// @Tags Organizations
// @Description Get lineage of an organization
// @Produce json
// @Param id path string true "Organization ID"
// @Success 200 {object} GetLineageResponseDto "Success Response"
// @Failure 400 "BadRequest"
// @Failure 500 "InternalServerError"
// @Router /organizations/{id}/lineage [GET]
func (c *Controller) GetLineage(w http.ResponseWriter, r *http.Request) {
	data, err := c.OrganizationService.RetrieveLineageById(
		chi.URLParam(r, "id"),
	)

	dtobuilderwithdata.New[tree.Node[Entity]](data, err).RenderTo(w)
}

// Get Siblings and Ancestral Siblings of Organizations : HTTP endpoint to get the siblings and ancestral siblings of an organization
// @Tags Organizations
// @Description Get siblings and ancestral siblings of an organization
// @Produce json
// @Param id path string true "Organization ID"
// @Success 200 {object} GetSiblingsAndAncestralSiblingsResponseDto "Success Response"
// @Failure 400 "BadRequest"
// @Failure 500 "InternalServerError"
// @Router /organizations/{id}/siblings-and-ancestral-siblings [GET]
func (c *Controller) GetSiblingsAndAncestralSiblings(w http.ResponseWriter, r *http.Request) {
	data, err := c.OrganizationService.RetrieveSiblingsAndAncestralSiblingsById(
		chi.URLParam(r, "id"),
	)
	dtobuilderwithdata.New[tree.Node[Entity]](data, err).RenderTo(w)
}

// Get Officers within Organizations : HTTP endpoint to get the officers within an organization
// @Tags Organizations
// @Description Get officers within an organization
// @Produce json
// @Param id path string true "Organization ID"
// @Success 200 {object} GetOfficersResponseDto "Success Response"
// @Failure 400 "BadRequest"
// @Failure 500 "InternalServerError"
// @Router /organizations/{id}/officers [GET]
func (c *Controller) GetOfficers(w http.ResponseWriter, r *http.Request) {
	data, err := c.PlacementService.RetrieveByOrganizationId(
		chi.URLParam(r, "id"),
	)
	dtobuilderwithdata.New[[]placement.Entity](&data, err).RenderTo(w)
}
