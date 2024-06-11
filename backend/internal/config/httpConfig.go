package config

import (
	"net"
	"os"
)

const (
	backendHost = "BACKEND_HOST"
	backendPort = "BACKEND_PORT"
)

type HTTPConfig interface {
	Address() string
}

type httpConfig struct {
	host string
	port string
}

func NewHTTPConfig() HTTPConfig {
	host := os.Getenv(backendHost)
	port := os.Getenv(backendPort)

	return &httpConfig{
		host: host,
		port: port,
	}
}

func (c *httpConfig) Address() string {
	return net.JoinHostPort(c.host, c.port)
}
