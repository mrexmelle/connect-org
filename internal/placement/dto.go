package placement

import (
	"github.com/mrexmelle/connect-orgs/internal/dto"
)

type PostRequestDto struct {
	OrganizationId string `json:"organization_id"`
	RoleId         string `json:"role_id"`
	Ehid           string `json:"ehid"`
}

type PatchRequestDto struct {
	Fields map[string]interface{} `json:"fields"`
}

type GetResponseDto = dto.HttpResponseWithData[Entity]
type PostResponseDto = dto.HttpResponseWithData[Entity]
type PatchResponseDto = dto.HttpResponseWithoutData
type DeleteResponseDto = dto.HttpResponseWithoutData
