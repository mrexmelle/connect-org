package role

import (
	"github.com/mrexmelle/connect-orgs/internal/config"
)

type Service struct {
	ConfigService  *config.Service
	RoleRepository Repository
}

func NewService(
	cfg *config.Service,
	r Repository,
) *Service {
	return &Service{
		ConfigService:  cfg,
		RoleRepository: r,
	}
}

func (s *Service) Create(req PostRequestDto) (*Entity, error) {
	result, err := s.RoleRepository.Create(&Entity{
		Id:       req.Id,
		Name:     req.Name,
		Rank:     req.Rank,
		MaxCount: req.MaxCount,
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *Service) RetrieveById(id string) (*Entity, error) {
	result, err := s.RoleRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *Service) UpdateById(fields map[string]interface{}, id string) error {
	return s.RoleRepository.UpdateById(fields, id)
}

func (s *Service) DeleteById(id string) error {
	err := s.RoleRepository.DeleteById(id)
	return err
}
