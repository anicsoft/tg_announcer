package config

import "os"

const (
	pathEnv       = "STORAGE_PATH"
	migrationsEnv = "MIGRATIONS_PATH"
)

type SQLiteConfig interface {
	Driver() string
	Migrations() string
	Path() string
}

type sqliteConf struct {
	driver     string
	migrations string
	path       string
}

func NewSQLiteConfig() SQLiteConfig {
	driver := "postgres"
	migrations := os.Getenv(migrationsEnv)
	path := os.Getenv(pathEnv)
	return &sqliteConf{
		driver:     driver,
		migrations: migrations,
		path:       path,
	}
}

func (c *sqliteConf) Driver() string {
	return c.driver
}

func (c *sqliteConf) Migrations() string {
	return c.migrations
}

func (c *sqliteConf) Path() string {
	return c.path
}
