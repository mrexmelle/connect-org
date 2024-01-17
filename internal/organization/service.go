package organization

import (
	"github.com/mrexmelle/connect-orgs/internal/config"
	"github.com/mrexmelle/connect-orgs/internal/tree"
)

type Service struct {
	ConfigService          *config.Service
	OrganizationRepository Repository
}

func NewService(
	cfg *config.Service,
	r Repository,
) *Service {
	return &Service{
		ConfigService:          cfg,
		OrganizationRepository: r,
	}
}

func (s *Service) Create(req PostRequestDto) (*Entity, error) {
	result, err := s.OrganizationRepository.Create(&Entity{
		Id:                  req.Id,
		Hierarchy:           req.Hierarchy,
		Name:                req.Name,
		EmailAddress:        req.EmailAddress,
		PrivateSlackChannel: req.PrivateSlackChannel,
		PublicSlackChannel:  req.PublicSlackChannel,
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *Service) RetrieveById(id string) (*Entity, error) {
	result, err := s.OrganizationRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *Service) UpdateById(fields map[string]interface{}, id string) error {
	return s.OrganizationRepository.UpdateById(fields, id)
}

func (s *Service) DeleteById(id string) error {
	err := s.OrganizationRepository.DeleteById(id)
	return err
}

func (s *Service) RetrieveChildrenById(id string) ([]Entity, error) {
	result, err := s.OrganizationRepository.FindChildrenById(id)
	if err != nil {
		return []Entity{}, err
	}
	return result, nil
}

func (s *Service) RetrieveLineageById(id string) (*tree.Node[Entity], error) {
	orgs, err := s.OrganizationRepository.FindLineageById(id)
	if err != nil {
		return nil, err
	}

	orgTree := tree.New[Entity](".")
	for i := 0; i < len(orgs); i++ {
		orgTree.AssignEntity(orgs[i].Hierarchy, &orgs[i])
	}
	return orgTree.Root, nil
}

func (s *Service) RetrieveSiblingsAndAncestralSiblingsById(id string) (*tree.Node[Entity], error) {
	orgs, err := s.OrganizationRepository.FindSiblingsAndAncestralSiblingsById(id)
	if err != nil {
		return nil, err
	}

	orgTree := tree.New[Entity](".")
	for i := 0; i < len(orgs); i++ {
		orgTree.AssignEntity(orgs[i].Hierarchy, &orgs[i])
	}
	return orgTree.Root, nil
}
