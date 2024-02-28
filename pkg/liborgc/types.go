package liborgc

import (
	"github.com/mrexmelle/connect-org/internal/dto"
	"github.com/mrexmelle/connect-org/internal/member"
	"github.com/mrexmelle/connect-org/internal/membership"
)

type GetMembershipHistoryResponseDto = member.GetHistoryResponseDto
type MembershipViewEntity = membership.ViewEntity
type ServiceError = dto.ServiceError
