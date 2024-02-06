package member

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/mrexmelle/connect-org/internal/config"
	"github.com/mrexmelle/connect-org/internal/dto/dtorespwithdata"
	"github.com/mrexmelle/connect-org/internal/localerror"
	"github.com/mrexmelle/connect-org/internal/membership"
)

type Controller struct {
	ConfigService     *config.Service
	LocalErrorService *localerror.Service
	MembershipService *membership.Service
}

func NewController(
	cfg *config.Service,
	les *localerror.Service,
	ms *membership.Service,
) *Controller {
	return &Controller{
		ConfigService:     cfg,
		LocalErrorService: les,
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
	info := c.LocalErrorService.Map(err)
	dtorespwithdata.New[[]membership.ViewEntity](
		&data,
		info.ServiceErrorCode,
		info.ServiceErrorMessage,
	).RenderTo(w, info.HttpStatusCode)
}

// Get History : HTTP endpoint to get the history of organization nodes
// @Tags Members
// @Description Get a member's current organization nodes
// @Produce json
// @Param ehid path string true "Employee Hash ID"
// @Param sort query string false "Start Date's sorting direction (asc or desc)"
// @Success 200 {object} GetHistoryResponseDto "Success Response"
// @Failure 400 "BadRequest"
// @Failure 500 "InternalServerError"
// @Router /members/{ehid}/history [GET]
func (c *Controller) GetHistory(w http.ResponseWriter, r *http.Request) {
	data, err := c.MembershipService.RetrieveByEhidOrderByStartDate(
		chi.URLParam(r, "ehid"),
		strings.ToUpper(r.URL.Query().Get("sort")),
	)
	info := c.LocalErrorService.Map(err)
	dtorespwithdata.New[[]membership.ViewEntity](
		&data,
		info.ServiceErrorCode,
		info.ServiceErrorMessage,
	).RenderTo(w, info.HttpStatusCode)
}
