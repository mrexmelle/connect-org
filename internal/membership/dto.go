package membership

import (
	"github.com/mrexmelle/connect-org/internal/dto/dtorespwithdata"
	"github.com/mrexmelle/connect-org/internal/dto/dtorespwithoutdata"
)

type PostRequestDto struct {
	Ehid      string `json:"ehid"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	NodeId    string `json:"node_id"`
}

type PatchRequestDto struct {
	Fields map[string]interface{} `json:"fields"`
}

type GetResponseDto = dtorespwithdata.Class[ViewEntity]
type PostResponseDto = dtorespwithdata.Class[ViewEntity]
type PatchResponseDto = dtorespwithoutdata.Class
type DeleteResponseDto = dtorespwithoutdata.Class
