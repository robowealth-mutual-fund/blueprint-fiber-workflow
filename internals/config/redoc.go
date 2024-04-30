package config

type Redoc struct {
	DocsPath string `env:"REDOC_DOCS_PATH" envDefault:"redoc"`
}
