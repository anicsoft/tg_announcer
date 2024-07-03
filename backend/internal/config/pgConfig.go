package config

import (
	"fmt"
	"os"
)

const (
	postgresHost     = "POSTGRES_HOST"
	postgresPort     = "POSTGRES_PORT"
	postgresUser     = "POSTGRES_USER"
	postgresPassword = "POSTGRES_PASSWORD"
	postgresDBName   = "POSTGRES_DB"
)

type PGConfig interface {
	DNS() string
}

type pgConfig struct {
	dns string
}

func NewPGConfig() PGConfig {
	dns := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv(postgresHost),
		os.Getenv(postgresPort),
		os.Getenv(postgresUser),
		os.Getenv(postgresPassword),
		os.Getenv(postgresDBName),
	)
	return &pgConfig{dns: dns}
}

func (p *pgConfig) DNS() string {
	return p.dns
}
