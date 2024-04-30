package schemas

import "github.com/getkin/kin-openapi/openapi3"

var CreateTodoResponse = openapi3.NewSchemaRef("",
	openapi3.NewObjectSchema().
		WithProperties(map[string]*openapi3.Schema{
			"taskName": {
				Type:        openapi3.TypeString,
				Example:     "Send Email",
				Description: "Example: ```Send Email```<br>Task Name",
			},
			"status": {
				Type:        openapi3.TypeString,
				Example:     "TODO",
				Description: "Example: ```TODO```<br>Status",
			},
			"createdAt": {
				Type:        openapi3.TypeInteger,
				Example:     "1714369904",
				Description: "Example: ```1714369904```<br>Create At",
			},
			"updatedAt": {
				Type:        openapi3.TypeInteger,
				Example:     "1714369904",
				Description: "Example: ```1714369904```<br>Update At",
			},
		}),
)
