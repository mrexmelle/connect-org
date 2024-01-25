package designation

import (
	"github.com/mrexmelle/connect-org/internal/config"
	"github.com/mrexmelle/connect-org/internal/localerror"
	"github.com/mrexmelle/connect-org/internal/role"
)

type Service struct {
	ConfigService       *config.Service
	PlacementRepository Repository
	RoleRepository      role.Repository
}

func NewService(
	cfg *config.Service,
	r Repository,
	rr role.Repository,
) *Service {
	return &Service{
		ConfigService:       cfg,
		PlacementRepository: r,
		RoleRepository:      rr,
	}
}

func (s *Service) Create(req PostRequestDto) (*Entity, error) {
	existingPlacements, err := s.PlacementRepository.FindByNodeIdAndRoleId(
		req.NodeId,
		req.RoleId,
	)
	if err != nil {
		existingPlacements = []Entity{}
	}

	newEntity := &Entity{
		NodeId: req.NodeId,
		RoleId: req.RoleId,
		Ehid:   req.Ehid,
	}

	for i := range existingPlacements {
		if existingPlacements[i].Ehid == req.Ehid &&
			existingPlacements[i].NodeId == req.NodeId &&
			existingPlacements[i].RoleId == req.RoleId {
			newEntity.Id = existingPlacements[i].Id
			return newEntity, nil
		}
	}

	maxCount := s.RoleRepository.CountById(req.RoleId)
	if (int64)(len(existingPlacements)) >= maxCount {
		return nil, localerror.ErrAlreadyMax
	}

	result, err := s.PlacementRepository.Create(newEntity)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *Service) RetrieveById(id string) (*Entity, error) {
	result, err := s.PlacementRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *Service) RetrieveByNodeId(nodeId string) ([]Entity, error) {
	result, err := s.PlacementRepository.FindByNodeId(nodeId)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *Service) RetrieveByNodeIdAndRoleId(nodeId string, roleId string) ([]Entity, error) {
	result, err := s.PlacementRepository.FindByNodeIdAndRoleId(nodeId, roleId)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *Service) UpdateById(fields map[string]interface{}, id string) error {
	return s.PlacementRepository.UpdateById(fields, id)
}

func (s *Service) DeleteById(id string) error {
	err := s.PlacementRepository.DeleteById(id)
	return err
}
