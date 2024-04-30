package todo

import (
	"github.com/robowealth-mutual-fund/blueprint-roa-fiber-golang/internals/config"
	todoRepository "github.com/robowealth-mutual-fund/blueprint-roa-fiber-golang/internals/repositories/blueprint/todo"
)

type Service struct {
	config         config.Config
	todoRepository todoRepository.Interface
}

func New(
	config config.Config,
	todoRepository todoRepository.Interface,
) Interface {
	return &Service{
		config:         config,
		todoRepository: todoRepository,
	}
}
