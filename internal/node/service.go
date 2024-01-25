package node

import (
	"github.com/mrexmelle/connect-org/internal/config"
	"github.com/mrexmelle/connect-org/internal/tree"
)

type Service struct {
	ConfigService  *config.Service
	NodeRepository Repository
}

func NewService(
	cfg *config.Service,
	r Repository,
) *Service {
	return &Service{
		ConfigService:  cfg,
		NodeRepository: r,
	}
}

func (s *Service) Create(req PostRequestDto) (*Entity, error) {
	result, err := s.NodeRepository.Create(&Entity{
		Id:           req.Id,
		Hierarchy:    req.Hierarchy,
		Name:         req.Name,
		EmailAddress: req.EmailAddress,
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *Service) RetrieveById(id string) (*Entity, error) {
	result, err := s.NodeRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *Service) UpdateById(fields map[string]interface{}, id string) error {
	return s.NodeRepository.UpdateById(fields, id)
}

func (s *Service) DeleteById(id string) error {
	err := s.NodeRepository.DeleteById(id)
	return err
}

func (s *Service) RetrieveChildrenById(id string) ([]Entity, error) {
	result, err := s.NodeRepository.FindChildrenById(id)
	if err != nil {
		return []Entity{}, err
	}
	return result, nil
}

func (s *Service) RetrieveLineageById(id string) (*tree.Node[Entity], error) {
	orgs, err := s.NodeRepository.FindLineageById(id)
	if err != nil {
		return nil, err
	}

	orgTree := tree.New[Entity](".")
	for i := 0; i < len(orgs); i++ {
		orgTree.AssignEntity(orgs[i].Hierarchy, &orgs[i])
	}
	return orgTree.Root, nil
}

func (s *Service) RetrieveLineageSiblingsById(id string) (*tree.Node[Entity], error) {
	orgs, err := s.NodeRepository.FindLineageSiblingsById(id)
	if err != nil {
		return nil, err
	}

	orgTree := tree.New[Entity](".")
	for i := 0; i < len(orgs); i++ {
		orgTree.AssignEntity(orgs[i].Hierarchy, &orgs[i])
	}
	return orgTree.Root, nil
}
