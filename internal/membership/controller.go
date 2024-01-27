package membership

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/mrexmelle/connect-org/internal/config"
	"github.com/mrexmelle/connect-org/internal/dto/dtobuilderwithdata"
	"github.com/mrexmelle/connect-org/internal/dto/dtobuilderwithoutdata"
	"github.com/mrexmelle/connect-org/internal/localerror"
)

type Controller struct {
	ConfigService     *config.Service
	MembershipService *Service
}

func NewController(cfg *config.Service, svc *Service) *Controller {
	return &Controller{
		ConfigService:     cfg,
		MembershipService: svc,
	}
}

// Get Memberships : HTTP endpoint to get memberships
// @Tags Memberships
// @Description Get a memberhsip
// @Produce json
// @Param id path string true "Membership ID"
// @Success 200 {object} GetResponseDto "Success Response"
// @Failure 400 "BadRequest"
// @Failure 500 "InternalServerError"
// @Router /memberships/{id} [GET]
func (c *Controller) Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		dtobuilderwithdata.New[Entity](nil, localerror.ErrIdNotInteger).RenderTo(w)
		return
	}
	data, err := c.MembershipService.RetrieveById(id)
	dtobuilderwithdata.New[ViewEntity](data, err).RenderTo(w)
}

// Post Memberships : HTTP endpoint to post new memberships
// @Tags Memberships
// @Description Post a new memberships
// @Accept json
// @Produce json
// @Param data body PostRequestDto true "Membership Request"
// @Success 200 {object} PostResponseDto "Success Response"
// @Failure 400 "BadRequest"
// @Failure 500 "InternalServerError"
// @Router /memberships [POST]
func (c *Controller) Post(w http.ResponseWriter, r *http.Request) {
	var requestBody PostRequestDto
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		dtobuilderwithdata.New[Entity](nil, localerror.ErrBadJson).RenderTo(w)
		return
	}

	data, err := c.MembershipService.Create(requestBody)
	dtobuilderwithdata.New[ViewEntity](data, err).RenderTo(w)
}

// Patch Memberships : HTTP endpoint to patch a membership
// @Tags Memberships
// @Description Patch a membership
// @Accept json
// @Produce json
// @Param id path string true "Membership ID"
// @Param data body PatchRequestDto true "Membership Patch Request"
// @Success 200 {object} PatchResponseDto "Success Response"
// @Failure 400 "BadRequest"
// @Failure 500 "InternalServerError"
// @Router /memberships/{id} [PATCH]
func (c *Controller) Patch(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		dtobuilderwithoutdata.New(localerror.ErrIdNotInteger).RenderTo(w)
		return
	}

	var requestBody PatchRequestDto
	err = json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		dtobuilderwithoutdata.New(localerror.ErrBadJson).RenderTo(w)
		return
	}

	err = c.MembershipService.UpdateById(requestBody.Fields, id)
	dtobuilderwithoutdata.New(err).RenderTo(w)
}

// Delete Memberships : HTTP endpoint to delete memberships
// @Tags Memberships
// @Description Delete a membership
// @Produce json
// @Param id path string true "Membership ID"
// @Success 200 {object} DeleteResponseDto "Success Response"
// @Failure 400 "BadRequest"
// @Failure 500 "InternalServerError"
// @Router /memberships/{id} [DELETE]
func (c *Controller) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		dtobuilderwithoutdata.New(localerror.ErrIdNotInteger).RenderTo(w)
		return
	}

	err = c.MembershipService.DeleteById(id)
	dtobuilderwithoutdata.New(err).RenderTo(w)
}
