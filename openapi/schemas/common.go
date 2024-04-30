package schemas

import "github.com/getkin/kin-openapi/openapi3"

var (
	ErrorBadRequest = openapi3.NewSchemaRef("",
		openapi3.NewArraySchema().WithItems(
			openapi3.NewObjectSchema().
				WithProperty("code", openapi3.NewStringSchema().
					WithDefault("PFWF400001")).
				WithProperty("desc", openapi3.NewStringSchema().
					WithDefault("ERROR_BAD_REQUEST: invalid request.")).
				WithProperty("cause", openapi3.NewStringSchema().
					WithDefault("invalid request"))))

	ErrorUnauthorized = openapi3.NewSchemaRef("",
		openapi3.NewArraySchema().WithItems(
			openapi3.NewObjectSchema().
				WithProperty("code", openapi3.NewStringSchema().
					WithDefault("PFWF401001")).
				WithProperty("desc", openapi3.NewStringSchema().
					WithDefault("ERROR_UNAUTHORIZED: Credential is invalid.")).
				WithProperty("cause", openapi3.NewStringSchema().
					WithDefault("authentication"))))

	ErrorForbidden = openapi3.NewSchemaRef("",
		openapi3.NewArraySchema().WithItems(
			openapi3.NewObjectSchema().
				WithProperty("code", openapi3.NewStringSchema().
					WithDefault("PFWF403001")).
				WithProperty("desc", openapi3.NewInt64Schema().
					WithDefault("ERROR_FORBIDDEN: Do not have any permissions.")).
				WithProperty("cause", openapi3.NewStringSchema().
					WithDefault("authorization"))))

	ErrorNotFound = openapi3.NewSchemaRef("",
		openapi3.NewArraySchema().WithItems(
			openapi3.NewObjectSchema().
				WithProperty("code", openapi3.NewStringSchema().
					WithDefault("PFWF404001")).
				WithProperty("desc", openapi3.NewStringSchema().
					WithDefault("ERROR_NOT_FOUND: Entity does not exist.")).
				WithProperty("cause", openapi3.NewStringSchema().
					WithDefault("entity does not exist"))))

	ErrorUnprocessableEntity = openapi3.NewSchemaRef("",
		openapi3.NewArraySchema().WithItems(
			openapi3.NewObjectSchema().
				WithProperty("code", openapi3.NewStringSchema().
					WithDefault("PFWF422001")).
				WithProperty("desc", openapi3.NewStringSchema().
					WithDefault("ERROR_UNPROCESSABLE_ENTITY: Unprocessable entities.")).
				WithProperty("cause", openapi3.NewStringSchema().
					WithDefault("unprocessable entity"))))

	ErrorInternal = openapi3.NewSchemaRef("",
		openapi3.NewArraySchema().WithItems(
			openapi3.NewObjectSchema().
				WithProperty("code", openapi3.NewStringSchema().
					WithDefault("PFWF500999")).
				WithProperty("desc", openapi3.NewStringSchema().
					WithDefault("ERROR_INTERNAL_SERVER: API Internal Server.")).
				WithProperty("cause", openapi3.NewStringSchema().
					WithDefault("internal"))))

	Pagination = openapi3.NewSchemaRef("",
		openapi3.NewObjectSchema().
			WithProperties(map[string]*openapi3.Schema{
				"size": {
					Type:    openapi3.TypeInteger,
					Example: 10,
				},
				"total": {
					Type:    openapi3.TypeInteger,
					Example: 11,
				},
				"totalPages": {
					Type:    openapi3.TypeInteger,
					Example: 2,
				},
			}),
	)
)
