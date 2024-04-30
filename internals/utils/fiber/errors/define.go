package errors

import "net/http"

func DefineBadRequest(code, desc, cause string) *GoError {
	return &GoError{
		Status: http.StatusBadRequest,
		Code:   code,
		Desc:   desc,
		Cause:  cause,
	}
}

func DefineUnauthorized(code, desc, cause string) *GoError {
	return &GoError{
		Status: http.StatusUnauthorized,
		Code:   code,
		Desc:   desc,
		Cause:  cause,
	}
}

func DefineForbidden(code, desc, cause string) *GoError {
	return &GoError{
		Status: http.StatusForbidden,
		Code:   code,
		Desc:   desc,
		Cause:  cause,
	}
}

func DefineNotFound(code, desc, cause string) *GoError {
	return &GoError{
		Status: http.StatusNotFound,
		Code:   code,
		Desc:   desc,
		Cause:  cause,
	}
}

func DefineStatusUnprocessableEntity(code, desc, cause string) *GoError {
	return &GoError{
		Status: http.StatusUnprocessableEntity,
		Code:   code,
		Desc:   desc,
		Cause:  cause,
	}
}

func DefineInternalServerError(code, desc, cause string) *GoError {
	return &GoError{
		Status: http.StatusInternalServerError,
		Code:   code,
		Desc:   desc,
		Cause:  cause,
	}
}

func DefineNotImplemented(code, desc, cause string) *GoError {
	return &GoError{
		Status: http.StatusNotImplemented,
		Code:   code,
		Desc:   desc,
		Cause:  cause,
	}
}

func DefineBadGateway(code, desc, cause string) *GoError {
	return &GoError{
		Status: http.StatusBadGateway,
		Code:   code,
		Desc:   desc,
		Cause:  cause,
	}
}

func DefineServiceUnavailable(code, desc, cause string) *GoError {
	return &GoError{
		Status: http.StatusServiceUnavailable,
		Code:   code,
		Desc:   desc,
		Cause:  cause,
	}
}

func DefineGatewayTimeout(code, desc, cause string) *GoError {
	return &GoError{
		Status: http.StatusGatewayTimeout,
		Code:   code,
		Desc:   desc,
		Cause:  cause,
	}
}
