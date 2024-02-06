package designation

import (
	"github.com/mrexmelle/connect-org/internal/dto/dtorespwithdata"
	"github.com/mrexmelle/connect-org/internal/dto/dtorespwithoutdata"
)

type PostRequestDto struct {
	NodeId string `json:"node_id"`
	RoleId string `json:"role_id"`
	Ehid   string `json:"ehid"`
}

type PatchRequestDto struct {
	Fields map[string]interface{} `json:"fields"`
}

type GetResponseDto = dtorespwithdata.Class[Entity]
type PostResponseDto = dtorespwithdata.Class[Entity]
type PatchResponseDto = dtorespwithoutdata.Class
type DeleteResponseDto = dtorespwithoutdata.Class
