package designation

import (
	"github.com/mrexmelle/connect-org/internal/config"
	"github.com/mrexmelle/connect-org/internal/localerror"
	"github.com/mrexmelle/connect-org/internal/role"
)

type Service struct {
	ConfigService         *config.Service
	DesignationRepository Repository
	RoleRepository        role.Repository
}

func NewService(
	cfg *config.Service,
	r Repository,
	rr role.Repository,
) *Service {
	return &Service{
		ConfigService:         cfg,
		DesignationRepository: r,
		RoleRepository:        rr,
	}
}

func (s *Service) Create(req PostRequestDto) (*Entity, error) {
	data, err := s.DesignationRepository.FindByNodeIdAndRoleId(
		req.NodeId,
		req.RoleId,
	)
	if err != nil {
		data = []Entity{}
	}

	newEntity := &Entity{
		NodeId: req.NodeId,
		RoleId: req.RoleId,
		Ehid:   req.Ehid,
	}

	for i := range data {
		if data[i].Ehid == req.Ehid &&
			data[i].NodeId == req.NodeId &&
			data[i].RoleId == req.RoleId {
			newEntity.Id = data[i].Id
			return newEntity, nil
		}
	}

	maxCount := s.RoleRepository.CountById(req.RoleId)
	if (int64)(len(data)) >= maxCount {
		return nil, localerror.ErrAlreadyMax
	}

	result, err := s.DesignationRepository.Create(newEntity)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *Service) RetrieveById(id string) (*Entity, error) {
	result, err := s.DesignationRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *Service) RetrieveByNodeId(nodeId string) ([]Entity, error) {
	result, err := s.DesignationRepository.FindByNodeId(nodeId)
	if err != nil {
		return []Entity{}, err
	}
	return result, nil
}

func (s *Service) RetrieveByNodeIdAndRoleId(nodeId string, roleId string) ([]Entity, error) {
	result, err := s.DesignationRepository.FindByNodeIdAndRoleId(nodeId, roleId)
	if err != nil {
		return []Entity{}, err
	}
	return result, nil
}

func (s *Service) UpdateById(fields map[string]interface{}, id string) error {
	return s.DesignationRepository.UpdateById(fields, id)
}

func (s *Service) DeleteById(id string) error {
	err := s.DesignationRepository.DeleteById(id)
	return err
}
