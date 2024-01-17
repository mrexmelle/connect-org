package organization

import (
	"github.com/mrexmelle/connect-orgs/internal/dto"
	"github.com/mrexmelle/connect-orgs/internal/placement"
	"github.com/mrexmelle/connect-orgs/internal/tree"
)

type PostRequestDto struct {
	Id                  string `json:"id"`
	Hierarchy           string `json:"hierarchy"`
	Name                string `json:"name"`
	EmailAddress        string `json:"email_address"`
	PrivateSlackChannel string `json:"private_slack_channel"`
	PublicSlackChannel  string `json:"public_slack_channel"`
}

type PatchRequestDto struct {
	Fields map[string]interface{} `json:"fields"`
}

type GetResponseDto = dto.HttpResponseWithData[Entity]
type PostResponseDto = dto.HttpResponseWithData[Entity]
type PatchResponseDto = dto.HttpResponseWithoutData
type DeleteResponseDto = dto.HttpResponseWithoutData
type GetChildrenResponseDto = dto.HttpResponseWithData[[]Entity]
type GetLineageResponseDto = dto.HttpResponseWithData[[]Entity]
type GetSiblingsAndAncestralSiblingsResponseDto = dto.HttpResponseWithData[tree.Node[*Entity]]
type GetOfficersResponseDto = dto.HttpResponseWithData[[]placement.Entity]
