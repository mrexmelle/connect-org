package config

import (
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Service struct {
	ConfigRepository Repository
	ReadDb           *gorm.DB
	WriteDb          *gorm.DB
}

func NewService(
	cr Repository,
) *Service {
	readDb, err := gorm.Open(
		postgres.Open(strings.TrimSpace(cr.GetReadDsn())),
		&gorm.Config{
			Logger:         logger.Default.LogMode(logger.Info),
			TranslateError: true,
		},
	)
	if err != nil {
		panic(err)
	}

	writeDb, err := gorm.Open(
		postgres.Open(strings.TrimSpace(cr.GetWriteDsn())),
		&gorm.Config{
			Logger:         logger.Default.LogMode(logger.Info),
			TranslateError: true,
		},
	)
	if err != nil {
		panic(err)
	}

	return &Service{
		ConfigRepository: cr,
		ReadDb:           readDb,
		WriteDb:          writeDb,
	}
}

func (s *Service) GetProfile() string {
	return s.ConfigRepository.GetProfile()
}

func (s *Service) GetPort() int {
	return s.ConfigRepository.GetPort()
}
