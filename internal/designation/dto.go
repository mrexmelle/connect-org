package designation

import (
	"github.com/mrexmelle/connect-org/internal/dto"
)

type PostRequestDto struct {
	NodeId string `json:"node_id"`
	RoleId string `json:"role_id"`
	Ehid   string `json:"ehid"`
}

type PatchRequestDto struct {
	Fields map[string]interface{} `json:"fields"`
}

type GetResponseDto = dto.HttpResponseWithData[Entity]
type PostResponseDto = dto.HttpResponseWithData[Entity]
type PatchResponseDto = dto.HttpResponseWithoutData
type DeleteResponseDto = dto.HttpResponseWithoutData
