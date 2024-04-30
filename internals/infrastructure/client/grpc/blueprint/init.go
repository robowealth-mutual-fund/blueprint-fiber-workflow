package blueprint

import (
	todoV1 "github.com/robowealth-mutual-fund/blueprint-proto/pkg/v1/todo"
	"github.com/robowealth-mutual-fund/blueprint-roa-fiber-golang/internals/config"
)

type Client struct {
	Config            config.Config
	TodoServiceClient todoV1.TodoServiceClient
}

func New(config config.Config) *Client {
	c := Client{
		Config: config,
	}
	c.configure()
	return &c
}
