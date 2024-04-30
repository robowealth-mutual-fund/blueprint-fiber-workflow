package wrapper

import (
	service "github.com/robowealth-mutual-fund/blueprint-roa-fiber-golang/internals/services/blueprint/todo"
	"go.uber.org/dig"
)

type Wrapper struct {
	dig.In  `name:"wrapperCreateTodo"`
	Service service.Interface
}
