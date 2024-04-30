package todo

import (
	"context"

	model "github.com/robowealth-mutual-fund/blueprint-roa-fiber-golang/internals/models/todo"
)

func (s *Service) Create(ctx context.Context, request *model.CreateRequest) (*model.CreateResponse, error) {
	return s.todoRepository.Create(ctx, request)
}
