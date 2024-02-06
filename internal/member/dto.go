package member

import (
	"github.com/mrexmelle/connect-org/internal/dto/dtorespwithdata"
	"github.com/mrexmelle/connect-org/internal/membership"
)

type GetNodesResponseDto = dtorespwithdata.Class[[]membership.ViewEntity]
type GetHistoryResponseDto = dtorespwithdata.Class[[]membership.ViewEntity]
