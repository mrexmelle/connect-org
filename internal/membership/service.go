package membership

import (
	"database/sql"
	"time"

	"github.com/mrexmelle/connect-org/internal/config"
)

type Service struct {
	ConfigService        *config.Service
	MembershipRepository Repository
}

func NewService(
	cfg *config.Service,
	r Repository,
) *Service {
	return &Service{
		ConfigService:        cfg,
		MembershipRepository: r,
	}
}

func (s *Service) Create(req PostRequestDto) (*ViewEntity, error) {
	sd, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		return nil, err
	}

	var ed sql.NullTime
	if req.EndDate == "" {
		ed.Valid = false
	} else {
		ed.Time, err = time.Parse("2006-01-02", req.EndDate)
		ed.Valid = (err == nil)
	}

	result, err := s.MembershipRepository.Create(&Entity{
		Ehid:      req.Ehid,
		StartDate: sd,
		EndDate:   ed,
		NodeId:    req.NodeId,
	})
	if err != nil {
		return nil, err
	}
	return toViewEntity(result), nil
}

func (s *Service) RetrieveById(id int) (*ViewEntity, error) {
	result, err := s.MembershipRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	return toViewEntity(result), nil
}

func (s *Service) UpdateById(fields map[string]interface{}, id int) error {
	return s.MembershipRepository.UpdateById(fields, id)
}

func (s *Service) DeleteById(id int) error {
	err := s.MembershipRepository.DeleteById(id)
	return err
}

func (s *Service) RetrieveByEhid(ehid string) ([]ViewEntity, error) {
	result, err := s.MembershipRepository.FindByEhid(ehid)
	if err != nil {
		return []ViewEntity{}, err
	}
	return toViewEntitySlice(result), nil
}

func (s *Service) RetrieveCurrentByNodeId(nodeId string) ([]ViewEntity, error) {
	result, err := s.MembershipRepository.FindCurrentByNodeId(nodeId)
	if err != nil {
		return []ViewEntity{}, err
	}
	return toViewEntitySlice(result), nil
}

func (s *Service) RetrieveCurrentByEhid(ehid string) ([]ViewEntity, error) {
	result, err := s.MembershipRepository.FindCurrentByEhid(ehid)
	if err != nil {
		return []ViewEntity{}, err
	}
	return toViewEntitySlice(result), nil
}
