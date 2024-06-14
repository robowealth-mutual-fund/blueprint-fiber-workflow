package todo

import (
	"context"

	model "github.com/robowealth-mutual-fund/blueprint-roa-fiber-golang/internals/models/todo"
)

func (s *Service) Delete(ctx context.Context, request *model.DeleteRequest) (*model.DeleteResponse, error) {
	return s.todoRepository.Delete(ctx, request)
}
