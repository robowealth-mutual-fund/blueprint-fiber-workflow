package todo

import (
	"github.com/robowealth-mutual-fund/blueprint-roa-fiber-golang/internals/config"
	"github.com/robowealth-mutual-fund/blueprint-roa-fiber-golang/internals/services/blueprint/todo/wrapper"
)

type Controller struct {
	config  config.Config
	todoSvc wrapper.Wrapper
}

func New(
	config config.Config,
	todoSvc wrapper.Wrapper,
) *Controller {
	return &Controller{
		config:  config,
		todoSvc: todoSvc,
	}
}
