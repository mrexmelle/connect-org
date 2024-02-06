package node

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/mrexmelle/connect-org/internal/config"
	"github.com/mrexmelle/connect-org/internal/designation"
	"github.com/mrexmelle/connect-org/internal/dto/dtorespwithdata"
	"github.com/mrexmelle/connect-org/internal/dto/dtorespwithoutdata"
	"github.com/mrexmelle/connect-org/internal/localerror"
	"github.com/mrexmelle/connect-org/internal/membership"
	"github.com/mrexmelle/connect-org/internal/tree"
)

type Controller struct {
	ConfigService      *config.Service
	LocalErrorService  *localerror.Service
	NodeService        *Service
	DesignationService *designation.Service
	MembershipService  *membership.Service
}

func NewController(
	cfg *config.Service,
	les *localerror.Service,
	svc *Service,
	ds *designation.Service,
	ms *membership.Service,
) *Controller {
	return &Controller{
		ConfigService:      cfg,
		LocalErrorService:  les,
		NodeService:        svc,
		DesignationService: ds,
		MembershipService:  ms,
	}
}

// Get Nodes : HTTP endpoint to get a node
// @Tags Nodes
// @Description Get a node
// @Produce json
// @Param id path string true "Node ID"
// @Success 200 {object} GetResponseDto "Success Response"
// @Failure 400 "BadRequest"
// @Failure 500 "InternalServerError"
// @Router /nodes/{id} [GET]
func (c *Controller) Get(w http.ResponseWriter, r *http.Request) {
	data, err := c.NodeService.RetrieveById(
		chi.URLParam(r, "id"),
	)
	info := c.LocalErrorService.Map(err)
	dtorespwithdata.New[Entity](
		data,
		info.ServiceErrorCode,
		info.ServiceErrorMessage,
	).RenderTo(w, info.HttpStatusCode)
}

// Post Nodes : HTTP endpoint to post new nodes
// @Tags Nodes
// @Description Post a new node
// @Accept json
// @Produce json
// @Param data body PostRequestDto true "Node Request"
// @Success 200 {object} PostResponseDto "Success Response"
// @Failure 400 "BadRequest"
// @Failure 500 "InternalServerError"
// @Router /nodes [POST]
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

	data, err := c.NodeService.Create(requestBody)
	info := c.LocalErrorService.Map(err)
	dtorespwithdata.New[Entity](
		data,
		info.ServiceErrorCode,
		info.ServiceErrorMessage,
	).RenderTo(w, info.HttpStatusCode)
}

// Patch Nodes : HTTP endpoint to patch a node
// @Tags Nodes
// @Description Patch a node
// @Accept json
// @Produce json
// @Param id path string true "Node ID"
// @Param data body PatchRequestDto true "Node Patch Request"
// @Success 200 {object} PatchResponseDto "Success Response"
// @Failure 400 "BadRequest"
// @Failure 500 "InternalServerError"
// @Router /nodes/{id} [PATCH]
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
	err = c.NodeService.UpdateById(requestBody.Fields, chi.URLParam(r, "id"))
	info := c.LocalErrorService.Map(err)
	dtorespwithoutdata.New(
		info.ServiceErrorCode,
		info.ServiceErrorMessage,
	).RenderTo(w, info.HttpStatusCode)
}

// Delete Nodes : HTTP endpoint to delete nodes
// @Tags Nodes
// @Description Delete a node
// @Produce json
// @Param id path string true "Node ID"
// @Success 200 {object} DeleteResponseDto "Success Response"
// @Failure 400 "BadRequest"
// @Failure 500 "InternalServerError"
// @Router /nodes/{id} [DELETE]
func (c *Controller) Delete(w http.ResponseWriter, r *http.Request) {
	err := c.NodeService.DeleteById(chi.URLParam(r, "id"))
	info := c.LocalErrorService.Map(err)
	dtorespwithoutdata.New(
		info.ServiceErrorCode,
		info.ServiceErrorMessage,
	).RenderTo(w, info.HttpStatusCode)
}

// Get Children of Nodes : HTTP endpoint to get the children of a node
// @Tags Nodes
// @Description Get children of a node
// @Produce json
// @Param id path string true "Node ID"
// @Success 200 {object} GetChildrenResponseDto "Success Response"
// @Failure 400 "BadRequest"
// @Failure 500 "InternalServerError"
// @Router /nodes/{id}/children [GET]
func (c *Controller) GetChildren(w http.ResponseWriter, r *http.Request) {
	data, err := c.NodeService.RetrieveChildrenById(
		chi.URLParam(r, "id"),
	)
	info := c.LocalErrorService.Map(err)
	dtorespwithdata.New[[]Entity](
		&data,
		info.ServiceErrorCode,
		info.ServiceErrorMessage,
	).RenderTo(w, info.HttpStatusCode)
}

// Get Lineage of Nodes : HTTP endpoint to get the lineage of a node
// @Tags Nodes
// @Description Get lineage of a node
// @Produce json
// @Param id path string true "Node ID"
// @Success 200 {object} GetLineageResponseDto "Success Response"
// @Failure 400 "BadRequest"
// @Failure 500 "InternalServerError"
// @Router /nodes/{id}/lineage [GET]
func (c *Controller) GetLineage(w http.ResponseWriter, r *http.Request) {
	data, err := c.NodeService.RetrieveLineageById(
		chi.URLParam(r, "id"),
	)
	info := c.LocalErrorService.Map(err)
	dtorespwithdata.New[tree.Node[Entity]](
		data,
		info.ServiceErrorCode,
		info.ServiceErrorMessage,
	).RenderTo(w, info.HttpStatusCode)
}

// Get Lineage Siblings of Nodes : HTTP endpoint to get the siblings and ancestral siblings of a node
// @Tags Nodes
// @Description Get siblings and ancestral siblings of a node
// @Produce json
// @Param id path string true "Node ID"
// @Success 200 {object} GetLineagelSiblingsResponseDto "Success Response"
// @Failure 400 "BadRequest"
// @Failure 500 "InternalServerError"
// @Router /nodes/{id}/lineage-siblings [GET]
func (c *Controller) GetLineageSiblings(w http.ResponseWriter, r *http.Request) {
	data, err := c.NodeService.RetrieveLineageSiblingsById(
		chi.URLParam(r, "id"),
	)
	info := c.LocalErrorService.Map(err)
	dtorespwithdata.New[tree.Node[Entity]](
		data,
		info.ServiceErrorCode,
		info.ServiceErrorMessage,
	).RenderTo(w, info.HttpStatusCode)
}

// Get Officers within Nodes : HTTP endpoint to get the officers within a node
// @Tags Nodes
// @Description Get officers within a node
// @Produce json
// @Param id path string true "Node ID"
// @Success 200 {object} GetOfficersResponseDto "Success Response"
// @Failure 400 "BadRequest"
// @Failure 500 "InternalServerError"
// @Router /nodes/{id}/officers [GET]
func (c *Controller) GetOfficers(w http.ResponseWriter, r *http.Request) {
	data, err := c.DesignationService.RetrieveByNodeId(
		chi.URLParam(r, "id"),
	)
	info := c.LocalErrorService.Map(err)
	dtorespwithdata.New[[]designation.Entity](
		&data,
		info.ServiceErrorCode,
		info.ServiceErrorMessage,
	).RenderTo(w, info.HttpStatusCode)
}

// Get Members within Nodes : HTTP endpoint to get the current members within a node
// @Tags Nodes
// @Description Get current members within a node
// @Produce json
// @Param id path string true "Node ID"
// @Success 200 {object} GetMembersResponseDto "Success Response"
// @Failure 400 "BadRequest"
// @Failure 500 "InternalServerError"
// @Router /nodes/{id}/members [GET]
func (c *Controller) GetMembers(w http.ResponseWriter, r *http.Request) {
	data, err := c.MembershipService.RetrieveCurrentByNodeId(
		chi.URLParam(r, "id"),
	)
	info := c.LocalErrorService.Map(err)
	dtorespwithdata.New[[]membership.ViewEntity](
		&data,
		info.ServiceErrorCode,
		info.ServiceErrorMessage,
	).RenderTo(w, info.HttpStatusCode)
}
