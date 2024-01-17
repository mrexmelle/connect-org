package role

import (
	"github.com/mrexmelle/connect-orgs/internal/dto"
)

type PostRequestDto struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Rank     int    `json:"rank"`
	MaxCount int    `json:"max_count"`
}

type PatchRequestDto struct {
	Fields map[string]interface{} `json:"fields"`
}

type GetResponseDto = dto.HttpResponseWithData[Entity]
type PostResponseDto = dto.HttpResponseWithData[Entity]
type PatchResponseDto = dto.HttpResponseWithoutData
type DeleteResponseDto = dto.HttpResponseWithoutData
