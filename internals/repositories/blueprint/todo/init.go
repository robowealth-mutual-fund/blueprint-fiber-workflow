package todo

import blueprintGPRC "github.com/robowealth-mutual-fund/blueprint-roa-fiber-golang/internals/infrastructure/client/grpc/blueprint"

type Repository struct {
	blueprintSvcGPRC *blueprintGPRC.Client
}

func New(blueprintSvcGPRC *blueprintGPRC.Client) Interface {
	return &Repository{
		blueprintSvcGPRC: blueprintSvcGPRC,
	}
}
