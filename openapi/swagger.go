package openapi

import (
	"github.com/getkin/kin-openapi/openapi3"

	"github.com/robowealth-mutual-fund/blueprint-roa-fiber-golang/internals/config"
	"github.com/robowealth-mutual-fund/blueprint-roa-fiber-golang/openapi/paths"
	"github.com/robowealth-mutual-fund/blueprint-roa-fiber-golang/openapi/responses"
	"github.com/robowealth-mutual-fund/blueprint-roa-fiber-golang/openapi/schemas"
)

const Description = `### API Version History
| Version | Description       | Date      | By       |
|---------|-------------------|-----------|----------|
| v0.0.1  | Initial Document | 2023-11-13|   |


### Error Description

| Err Code | Http Code | Error Desc                         | Remark |
|----------|-----------|------------------------------------|------|
| PFWF400001   | 400       | Unprocessable Entity 			|      |
| PFWF401001   | 401       | Unauthorized Error				|      |
| PFWF404001   | 404       | Not Found   		  		  	|      |
| PFWF422001   | 422       | {Field name} is required       |      |
| PFWF500001   | 500       | Unhandled Error                |      |
| PFWF504001   | 504       | Gateway Timeout                |      |
`

func NewOpenAPI3(config config.Config) openapi3.T {
	swagger := initializeSwagger(config)

	initializeSchemas(&swagger.Components.Schemas)
	initializeResponses(&swagger.Components.Responses)
	initializePaths(&swagger.Paths)

	return swagger
}

func initializeSwagger(config config.Config) openapi3.T {
	return openapi3.T{
		OpenAPI: "3.0.0",
		Info: &openapi3.Info{
			Title:       config.Swagger.Title,
			Version:     config.ServiceVersion,
			Description: Description,
		},
		Servers: openapi3.Servers{
			&openapi3.Server{
				URL: config.APIUrl + config.BasePath,
			},
		},
		Components: &openapi3.Components{},
	}
}

func initializeSchemas(input *openapi3.Schemas) {
	*input = openapi3.Schemas{
		"ErrorBadRequest":          schemas.ErrorBadRequest,
		"ErrorUnauthorized":        schemas.ErrorUnauthorized,
		"ErrorForbidden":           schemas.ErrorForbidden,
		"ErrorNotFound":            schemas.ErrorNotFound,
		"ErrorUnprocessableEntity": schemas.ErrorUnprocessableEntity,
		"ErrorInternal":            schemas.ErrorInternal,
		"Pagination":               schemas.Pagination,
		"CreateTodoResponse":       schemas.CreateTodoResponse,
	}
}

func initializeResponses(input *openapi3.Responses) {
	*input = openapi3.Responses{
		"ErrorBadRequestResponse":          responses.ErrorBadRequest,
		"ErrorUnauthorizedResponse":        responses.ErrorUnauthorized,
		"ErrorForbiddenResponse":           responses.ErrorForbidden,
		"ErrorNotFoundResponse":            responses.ErrorNotFound,
		"ErrorUnprocessableEntityResponse": responses.ErrorUnprocessableEntity,
		"ErrorInternalResponse":            responses.ErrorInternal,
		"CreateTodoResponse":               responses.CreateTodoResponse,
	}
}

func initializePaths(input *openapi3.Paths) {
	*input = openapi3.Paths{
		"/api/todo": paths.CreateTodo,
	}
}
