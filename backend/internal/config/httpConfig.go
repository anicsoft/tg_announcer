package config

import (
	"net"
	"os"
)

const (
	httpHostEnv = "COMPANIES_HTTP_HOST"
	httpPortEnv = "COMPANIES_HTTP_PORT"
)

type HTTPConfig interface {
	Address() string
}

type httpConfig struct {
	host string
	port string
}

func NewHTTPConfig() HTTPConfig {
	host := os.Getenv(httpHostEnv)
	port := os.Getenv(httpPortEnv)
	return &httpConfig{
		host: host,
		port: port,
	}
}

func (c *httpConfig) Address() string {
	return net.JoinHostPort(c.host, c.port)
}
