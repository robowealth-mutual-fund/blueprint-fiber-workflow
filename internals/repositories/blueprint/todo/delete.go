package todo

import (
	"context"

	"github.com/robowealth-mutual-fund/blueprint-proto/pkg/v1/todo/message"
	model "github.com/robowealth-mutual-fund/blueprint-roa-fiber-golang/internals/models/todo"
)

func (r *Repository) Delete(ctx context.Context, request *model.DeleteRequest) (*model.DeleteResponse, error) {
	response, err := r.blueprintSvcGPRC.TodoServiceClient.Delete(ctx, &message.DeleteRequest{Id: request.Id})
	if err != nil {
		return nil, err
	}

	return &model.DeleteResponse{Message: response.Message}, nil
}
