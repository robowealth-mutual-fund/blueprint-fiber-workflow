package errors

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type ErrorResponse struct {
	Errors []ErrorDetail `json:"errors"`
}

type ErrorDetail struct {
	Code  string `json:"code"`
	Desc  string `json:"desc"`
	Cause string `json:"cause"`
}

func RespWithError(c *fiber.Ctx, err error) error {
	var e *GoError
	if errors.As(err, &e) {

		return c.Status(e.Status).JSON(ErrorResponse{Errors: []ErrorDetail{{
			Code:  e.Code,
			Desc:  e.Desc,
			Cause: e.Cause,
		}}})
	}

	return c.Status(http.StatusInternalServerError).JSON(ErrorResponse{Errors: []ErrorDetail{{
		Code:  UnhandledError,
		Desc:  err.Error(),
		Cause: "",
	}}})
}

func RespValidateError(c *fiber.Ctx, err *GoError, serviceCode string) error {
	validateErrors := make([]ErrorDetail, 0)

	if serviceCode == "" {
		serviceCode = "FP"
	}

	for _, field := range err.Fields {
		validateErrors = append(validateErrors, ErrorDetail{
			Code:  serviceCode + TagValidators[field.Tag],
			Desc:  fmt.Sprintf("%s: %s", ErrorUnprocessableEntity, field.Desc),
			Cause: field.FieldName,
		})
	}

	return c.Status(http.StatusUnprocessableEntity).JSON(ErrorResponse{Errors: validateErrors})
}
