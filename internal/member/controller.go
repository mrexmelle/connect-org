package member

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/mrexmelle/connect-org/internal/config"
	"github.com/mrexmelle/connect-org/internal/dto/dtobuilderwithdata"
	"github.com/mrexmelle/connect-org/internal/membership"
)

type Controller struct {
	ConfigService     *config.Service
	MembershipService *membership.Service
}

func NewController(cfg *config.Service, ms *membership.Service) *Controller {
	return &Controller{
		ConfigService:     cfg,
		MembershipService: ms,
	}
}

// Get Nodes : HTTP endpoint to get current organization nodes
// @Tags Members
// @Description Get a member's current organization nodes
// @Produce json
// @Param ehid path string true "Employee Hash ID"
// @Success 200 {object} GetNodesResponseDto "Success Response"
// @Failure 400 "BadRequest"
// @Failure 500 "InternalServerError"
// @Router /members/{ehid}/nodes [GET]
func (c *Controller) GetNodes(w http.ResponseWriter, r *http.Request) {
	data, err := c.MembershipService.RetrieveCurrentByEhid(chi.URLParam(r, "ehid"))
	dtobuilderwithdata.New[[]membership.ViewEntity](&data, err).RenderTo(w)
}

// Get History : HTTP endpoint to get the history of organization nodes
// @Tags Members
// @Description Get a member's current organization nodes
// @Produce json
// @Param ehid path string true "Employee Hash ID"
// @Success 200 {object} GetHistoryResponseDto "Success Response"
// @Failure 400 "BadRequest"
// @Failure 500 "InternalServerError"
// @Router /members/{ehid}/history [GET]
func (c *Controller) GetHistory(w http.ResponseWriter, r *http.Request) {
	data, err := c.MembershipService.RetrieveByEhid(chi.URLParam(r, "ehid"))
	dtobuilderwithdata.New[[]membership.ViewEntity](&data, err).RenderTo(w)
}
