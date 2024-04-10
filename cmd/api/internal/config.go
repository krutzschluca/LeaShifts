package config

import (
	"fmt"

	"github.com/alexflint/go-arg"
)

type Config struct {
	Listen             string `arg:"--listen,env:LISTEN"`
	DBConnectionString string `arg:"--dbconnectionstring,env:DBCONNECTIONSTRING"`
	DBDriver           string `arg:"--dbdriver,env:DBDRIVER"`
}

func New() (*Config, error) {
	c := &Config{
		Listen:   ":8080",
		DBDriver: "pgx",
	}

	err := arg.Parse(c)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}
	if c.DBConnectionString == "" {
		return nil, fmt.Errorf("database connection string is missing")
	}

	return c, nil
}
