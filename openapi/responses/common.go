package responses

import "github.com/getkin/kin-openapi/openapi3"

var (
	ErrorBadRequest = &openapi3.ResponseRef{
		Value: openapi3.NewResponse().
			WithDescription("Bad Request").
			WithContent(
				openapi3.NewContentWithJSONSchema(
					openapi3.NewSchema().
						WithPropertyRef("errors",
							openapi3.NewSchemaRef("#/components/schemas/ErrorBadRequest", openapi3.NewObjectSchema()),
						),
				),
			),
	}

	ErrorUnauthorized = &openapi3.ResponseRef{
		Value: openapi3.NewResponse().
			WithDescription("Unauthorized").
			WithContent(
				openapi3.NewContentWithJSONSchema(
					openapi3.NewSchema().
						WithPropertyRef("errors",
							openapi3.NewSchemaRef("#/components/schemas/ErrorUnauthorized", openapi3.NewObjectSchema()),
						),
				),
			),
	}

	ErrorForbidden = &openapi3.ResponseRef{
		Value: openapi3.NewResponse().
			WithDescription("Forbidden").
			WithContent(
				openapi3.NewContentWithJSONSchema(
					openapi3.NewSchema().
						WithPropertyRef("errors",
							openapi3.NewSchemaRef("#/components/schemas/ErrorForbidden", openapi3.NewObjectSchema()),
						),
				),
			),
	}

	ErrorNotFound = &openapi3.ResponseRef{
		Value: openapi3.NewResponse().
			WithDescription("Not Found").
			WithContent(
				openapi3.NewContentWithJSONSchema(
					openapi3.NewSchema().
						WithPropertyRef("errors",
							openapi3.NewSchemaRef("#/components/schemas/ErrorNotFound", openapi3.NewObjectSchema()),
						),
				),
			),
	}

	ErrorUnprocessableEntity = &openapi3.ResponseRef{
		Value: openapi3.NewResponse().
			WithDescription("Unprocessable Entity").
			WithContent(
				openapi3.NewContentWithJSONSchema(
					openapi3.NewSchema().
						WithPropertyRef("errors",
							openapi3.NewSchemaRef("#/components/schemas/ErrorUnprocessableEntity", openapi3.NewObjectSchema()),
						),
				),
			),
	}

	ErrorInternal = &openapi3.ResponseRef{
		Value: openapi3.NewResponse().
			WithDescription("API Internal Error").
			WithContent(
				openapi3.NewContentWithJSONSchema(
					openapi3.NewSchema().
						WithPropertyRef("errors",
							openapi3.NewSchemaRef("#/components/schemas/ErrorInternal", openapi3.NewObjectSchema()),
						),
				),
			),
	}
)
