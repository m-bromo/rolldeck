package config

import (
	"fmt"

	"github.com/Netflix/go-env"
	"github.com/joho/godotenv"
)

func NewConfig() (*Config, error) {
	if err := godotenv.Load(".env"); err != nil {
		return nil, err
	}

	var environment Config
	_, err := env.UnmarshalFromEnviron(&environment)
	if err != nil {
		return nil, err
	}

	environment.Api.Addr = fmt.Sprintf("%s:%s", environment.Api.Host, environment.Api.Port)
	environment.PostgresDB.DSN = fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s",
		environment.PostgresDB.Host,
		environment.PostgresDB.Port,
		environment.PostgresDB.Name,
		environment.PostgresDB.User,
		environment.PostgresDB.Password,
	)

	return &environment, nil
}
