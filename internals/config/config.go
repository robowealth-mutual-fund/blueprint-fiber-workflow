package config

import (
	"github.com/caarlos0/env/v10"
	_ "github.com/joho/godotenv/autoload"
	log "github.com/robowealth-mutual-fund/stdlog"
)

type Config struct {
	AppName         string `env:"APP_NAME" envDefault:"blueprint-fiber-workflow"`
	BasePath        string `env:"BASE_PATH" envDefault:""`
	APIUrl          string `env:"API_URL" envDefault:"http://localhost:4002"`
	Environment     string `env:"ENV" envDefault:"develop"`
	PlatformName    string `env:"PLATFORM_NAME" envDefault:"Blueprint Fiber Workflow"`
	ServiceVersion  string `env:"SERVICE_VERSION" envDefault:"v0.0.0"`
	ServiceBasePath string `env:"SERVICE_BASE_PATH" envDefault:"blueprint-fiber-workflow"`
	ServiceCode     string `env:"SERVICE_CODE" envDefault:"PFWF"`
	ReadBufferSize  int    `env:"READ_BUFFER_SIZE" envDefault:"40690"`

	Connection Connection
	Swagger    Swagger
	Redoc      Redoc
	Trace      Trace
	Server     Server
}

func New() Config {
	c := Config{}

	if err := env.Parse(&c); err != nil {
		log.Error("Load configuration failed", err)
		panic(err)
	}

	return c
}
