package responses

import "github.com/getkin/kin-openapi/openapi3"

var CreateTodoResponse = &openapi3.ResponseRef{
	Value: openapi3.NewResponse().
		WithContent(
			openapi3.NewContentWithJSONSchemaRef(
				openapi3.NewSchemaRef("#/components/schemas/CreateTodoResponse", openapi3.NewObjectSchema()),
			),
		),
}
