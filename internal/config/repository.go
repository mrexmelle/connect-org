package config

import (
	"os"

	"github.com/spf13/viper"
)

type Repository interface {
	GetProfile() string
	GetDsn() string
	GetDbUser() string
	GetDbPassword() string
	GetPort() int
}

type RepositoryImpl struct {
	Profile    string
	Dsn        string
	DbUser     string
	DbPassword string
	Port       int
}

func NewRepository() Repository {
	profile := os.Getenv("APP_PROFILE")
	if profile == "" {
		profile = "local"
	}
	viper.SetConfigName("application-" + profile)
	viper.SetConfigType("yml")
	for _, cp := range []string{
		"/etc/conf",
		"./config",
	} {
		viper.AddConfigPath(cp)
	}
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	dsn := "bolt://" + viper.GetString("app.datasource.host") + ":" + viper.GetString("app.datasource.port")

	dbUser := viper.GetString("app.datasource.user")
	dbPassword := viper.GetString("app.datasource.password")

	port := viper.GetInt("app.server.port")

	return &RepositoryImpl{
		Profile:    profile,
		Dsn:        dsn,
		DbUser:     dbUser,
		DbPassword: dbPassword,
		Port:       port,
	}
}

func (r *RepositoryImpl) GetProfile() string {
	return r.Profile
}

func (r *RepositoryImpl) GetDsn() string {
	return r.Dsn
}

func (r *RepositoryImpl) GetDbUser() string {
	return r.DbUser
}

func (r *RepositoryImpl) GetDbPassword() string {
	return r.DbPassword
}

func (r *RepositoryImpl) GetPort() int {
	return r.Port
}
