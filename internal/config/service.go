package config

type Service struct {
	ConfigRepository Repository
}

func NewService(
	cr Repository,
) *Service {
	return &Service{
		ConfigRepository: cr,
	}
}

func (s *Service) GetProfile() string {
	return s.ConfigRepository.GetProfile()
}

func (s *Service) GetPort() int {
	return s.ConfigRepository.GetPort()
}
