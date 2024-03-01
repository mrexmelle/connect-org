package liborgc

import (
	"github.com/mrexmelle/connect-org/internal/dto"
	"github.com/mrexmelle/connect-org/internal/member"
	"github.com/mrexmelle/connect-org/internal/membership"
)

type GetMemberHistoryResponseDto = member.GetHistoryResponseDto
type GetMemberNodesResponseDto = member.GetNodesResponseDto
type MembershipViewEntity = membership.ViewEntity
type ServiceError = dto.ServiceError
