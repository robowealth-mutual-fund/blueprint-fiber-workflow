package todo

import (
	"context"

	model "github.com/robowealth-mutual-fund/blueprint-roa-fiber-golang/internals/models/todo"
)

func (s *Service) Update(ctx context.Context, request *model.UpdateRequest) (*model.UpdateResponse, error) {
	return s.todoRepository.Update(ctx, request)
}
