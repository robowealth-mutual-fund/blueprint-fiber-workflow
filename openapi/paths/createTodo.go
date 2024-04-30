package paths

import "github.com/getkin/kin-openapi/openapi3"

var CreateTodo = &openapi3.PathItem{
	Post: &openapi3.Operation{
		Tags:        []string{"Todo"},
		Description: "Create Todo",
		Parameters: []*openapi3.ParameterRef{
			{
				Value: &openapi3.Parameter{
					Name:        "Authorization",
					In:          "header",
					Required:    true,
					Description: "Authorization: Bearer {token}",
					Schema: &openapi3.SchemaRef{
						Value: &openapi3.Schema{
							Type: "string",
							Example: "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.\n" +
								"eyJjdXN0b21lcl9pZCI6Im5XQllQSitkMWY5VXkwM28rQ\n" +
								"2U4c3c9PSIsImV4cCI6MTY5MjU5OTgxMSwiaWF0IjoxNj\n" +
								"kyNTk2MjExLCJpdiI6Im4vR05DNFc3eHVJYjNyNTRBYW5\n" +
								"sbHc9PSIsInVzZXJfaWQiOiI3a3lPYk9jNEJzT2l5Z1My\n" +
								"MXdxOEpRPT0ifQ.fzZEgKIxqfN9FQKjo48dzOdrmoikMh\n",
						},
					},
				},
			},
		},
		RequestBody: &openapi3.RequestBodyRef{
			Value: &openapi3.RequestBody{
				Description: "Payload to create todo",
				Required:    true,
				Content: openapi3.Content{
					"application/json": &openapi3.MediaType{
						Schema: &openapi3.SchemaRef{
							Value: &openapi3.Schema{
								Type: "object",
								Properties: map[string]*openapi3.SchemaRef{
									"taskName": {
										Value: &openapi3.Schema{
											Type:        "string",
											Example:     "Send Email",
											Description: "Example: ```Send Email```<br>task name",
											MaxLength:   &[]uint64{50}[0],
										},
									},
									"status": {
										Value: &openapi3.Schema{
											Type:        "string",
											Example:     "TODO",
											Description: "Example: ```TODO```<br>status",
											MaxLength:   &[]uint64{20}[0],
										},
									},
								},
							},
						},
					},
				},
			},
		},
		Responses: openapi3.Responses{
			"200": &openapi3.ResponseRef{
				Ref: "#/components/responses/CreateTodoResponse",
			},
			"400": &openapi3.ResponseRef{
				Ref: "#/components/responses/ErrorBadRequestResponse",
			},
			"401": &openapi3.ResponseRef{
				Ref: "#/components/responses/ErrorUnauthorizedResponse",
			},
			"500": &openapi3.ResponseRef{
				Ref: "#/components/responses/ErrorInternalResponse",
			},
		},
	},
}
