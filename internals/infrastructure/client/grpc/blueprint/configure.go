package blueprint

import (
	todoV1 "github.com/robowealth-mutual-fund/blueprint-proto/pkg/v1/todo"
)

func (c *Client) configure() {
	info := c.connectToService(c.Config.Connection.BluePrintService)
	c.TodoServiceClient = todoV1.NewTodoServiceClient(info)
}
