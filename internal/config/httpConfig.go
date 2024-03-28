package config

import (
	"net"
	"os"
)

const (
	httpHostEnv = "https://obscure-island-92871-9f528af4e4b1.herokuapp.com"
	httpPortEnv = "PORT"
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
	if len(port) == 0 {
		port = os.Getenv("COMPANIES_HTTP_PORT")
	}
	return &httpConfig{
		host: host,
		port: port,
	}
}

func (c *httpConfig) Address() string {
	return net.JoinHostPort(c.host, c.port)
}
