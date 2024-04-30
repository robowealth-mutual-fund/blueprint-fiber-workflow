package config

type Swagger struct {
	Title       string `env:"SWAGGER_TITLE" envDefault:"BluePrint Workflow"`
	Description string `env:"SWAGGER_DESCRIPTION" envDefault:"BluePrint Workflow API"`
}
