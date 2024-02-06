package role

import (
	"github.com/mrexmelle/connect-org/internal/dto/dtorespwithdata"
	"github.com/mrexmelle/connect-org/internal/dto/dtorespwithoutdata"
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

type GetResponseDto = dtorespwithdata.Class[Entity]
type PostResponseDto = dtorespwithdata.Class[Entity]
type PatchResponseDto = dtorespwithoutdata.Class
type DeleteResponseDto = dtorespwithoutdata.Class
