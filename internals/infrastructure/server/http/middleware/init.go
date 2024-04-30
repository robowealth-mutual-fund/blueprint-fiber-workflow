package middleware

import "github.com/robowealth-mutual-fund/blueprint-roa-fiber-golang/internals/config"

type Middleware struct {
	config config.Config
}

func New(config config.Config) *Middleware {
	return &Middleware{
		config: config,
	}
}
