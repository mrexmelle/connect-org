package member

import (
	"github.com/mrexmelle/connect-org/internal/dto"
	"github.com/mrexmelle/connect-org/internal/membership"
)

type GetNodesResponseDto = dto.HttpResponseWithData[[]membership.ViewEntity]
type GetHistoryResponseDto = dto.HttpResponseWithData[[]membership.ViewEntity]
