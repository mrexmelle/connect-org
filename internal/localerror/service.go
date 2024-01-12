package localerror

import (
	"net/http"

	"github.com/mrexmelle/connect-orgs/internal/config"
)

type Service struct {
}

func NewService(cfg *config.Service) *Service {
	return &Service{}
}

func (s *Service) Map(err error) StatusInfo {
	if err == nil {
		return NewStatusInfo(http.StatusOK, ErrSvcCodeNone, "OK")
	}

	codePair, exists := ErrorMap[err]
	if exists {
		return NewStatusInfo(codePair.HttpStatusCode, codePair.ServiceErrorCode, err.Error())
	}

	return NewStatusInfo(http.StatusBadRequest, ErrSvcCodeUnregistered, err.Error())
}
