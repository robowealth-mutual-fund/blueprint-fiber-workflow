package http

import (
	healthController "github.com/robowealth-mutual-fund/blueprint-roa-fiber-golang/internals/controllers/health"
	todoController "github.com/robowealth-mutual-fund/blueprint-roa-fiber-golang/internals/controllers/todo"
)

type Controller struct {
	health *healthController.Controller
	todo   *todoController.Controller
}

func NewController(
	health *healthController.Controller,
	todo *todoController.Controller,
) *Controller {
	return &Controller{
		health: health,
		todo:   todo,
	}
}
