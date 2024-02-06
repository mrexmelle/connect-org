package node

import (
	"github.com/mrexmelle/connect-org/internal/designation"
	"github.com/mrexmelle/connect-org/internal/dto/dtorespwithdata"
	"github.com/mrexmelle/connect-org/internal/dto/dtorespwithoutdata"
	"github.com/mrexmelle/connect-org/internal/membership"
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

type GetResponseDto = dtorespwithdata.Class[Entity]
type PostResponseDto = dtorespwithdata.Class[Entity]
type PatchResponseDto = dtorespwithoutdata.Class
type DeleteResponseDto = dtorespwithoutdata.Class
type GetChildrenResponseDto = dtorespwithdata.Class[[]Entity]
type GetLineageResponseDto = dtorespwithdata.Class[tree.Node[Entity]]
type GetLineagelSiblingsResponseDto = dtorespwithdata.Class[tree.Node[Entity]]
type GetOfficersResponseDto = dtorespwithdata.Class[[]designation.Entity]
type GetMembersResponseDto = dtorespwithdata.Class[[]membership.ViewEntity]
