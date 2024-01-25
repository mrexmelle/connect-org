package node

import (
	"github.com/mrexmelle/connect-org/internal/designation"
	"github.com/mrexmelle/connect-org/internal/dto"
	"github.com/mrexmelle/connect-org/internal/tree"
)

type PostRequestDto struct {
	Id           string `json:"id"`
	Hierarchy    string `json:"hierarchy"`
	Name         string `json:"name"`
	EmailAddress string `json:"email_address"`
}

type PatchRequestDto struct {
	Fields map[string]interface{} `json:"fields"`
}

type GetResponseDto = dto.HttpResponseWithData[Entity]
type PostResponseDto = dto.HttpResponseWithData[Entity]
type PatchResponseDto = dto.HttpResponseWithoutData
type DeleteResponseDto = dto.HttpResponseWithoutData
type GetChildrenResponseDto = dto.HttpResponseWithData[[]Entity]
type GetLineageResponseDto = dto.HttpResponseWithData[tree.Node[Entity]]
type GetLineagelSiblingsResponseDto = dto.HttpResponseWithData[tree.Node[Entity]]
type GetOfficersResponseDto = dto.HttpResponseWithData[[]designation.Entity]
