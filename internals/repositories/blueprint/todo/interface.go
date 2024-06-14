package todo

import (
	"context"

	model "github.com/robowealth-mutual-fund/blueprint-roa-fiber-golang/internals/models/todo"
)

//go:generate mockery --name=Interface
type Interface interface {
	Create(ctx context.Context, request *model.CreateRequest) (*model.CreateResponse, error)
	Update(ctx context.Context, request *model.UpdateRequest) (*model.UpdateResponse, error)
	Delete(ctx context.Context, request *model.DeleteRequest) (*model.DeleteResponse, error)
}
