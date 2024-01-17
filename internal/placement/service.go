package placement

import (
	"github.com/mrexmelle/connect-orgs/internal/config"
)

type Service struct {
	ConfigService     *config.Service
	OrgRoleRepository Repository
}

func NewService(
	cfg *config.Service,
	r Repository,
) *Service {
	return &Service{
		ConfigService:     cfg,
		OrgRoleRepository: r,
	}
}

func (s *Service) Create(req PostRequestDto) (*Entity, error) {
	result, err := s.OrgRoleRepository.Create(&Entity{
		OrganizationId: req.OrganizationId,
		RoleId:         req.RoleId,
		Ehid:           req.Ehid,
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *Service) RetrieveById(id string) (*Entity, error) {
	result, err := s.OrgRoleRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *Service) RetrieveByOrganizationId(organizationId string) ([]Entity, error) {
	result, err := s.OrgRoleRepository.FindByOrganizationId(organizationId)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *Service) UpdateById(fields map[string]interface{}, id string) error {
	return s.OrgRoleRepository.UpdateById(fields, id)
}

func (s *Service) DeleteById(id string) error {
	err := s.OrgRoleRepository.DeleteById(id)
	return err
}
