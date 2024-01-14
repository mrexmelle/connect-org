package organization

import (
	"github.com/mrexmelle/connect-orgs/internal/dto"
	"github.com/mrexmelle/connect-orgs/internal/tree"
)

type PostRequestDto struct {
	Id                  string `json:"id"`
	Hierarchy           string `json:"hierarchy"`
	Name                string `json:"name"`
	EmailAddress        string `json:"emailAddress"`
	PrivateSlackChannel string `json:"privateSlackChannel"`
	PublicSlackChannel  string `json:"publicSlackChannel"`
}

type GetResponseDto = dto.HttpResponseWithData[Entity]
type PostResponseDto = dto.HttpResponseWithData[Entity]
type DeleteResponseDto = dto.HttpResponseWithoutData
type GetChildrenResponseDto = dto.HttpResponseWithData[[]Entity]
type GetLineageResponseDto = dto.HttpResponseWithData[[]Entity]
type GetSiblingsAndAncestralSiblingsDto dto.HttpResponseWithData[tree.Node[*Entity]]
