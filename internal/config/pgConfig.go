package config

import "os"

const (
	dsnEnvName = "DATABASE_URL"
)

type PGConfig interface {
	DNS() string
}

type pgConfig struct {
	dns string
}

func NewPGConfig() PGConfig {
	dns := os.Getenv(driverEnv)
	return &pgConfig{dns: dns}
}

func (p *pgConfig) DNS() string {
	return p.dns
}
