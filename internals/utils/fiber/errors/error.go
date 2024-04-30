package errors

import (
	"fmt"
)

type Error interface {
	Error() string
	ErrorWithCause() string
	IsCodeEqual(err error) bool
}

type FieldError struct {
	FieldName string
	Desc      string
	Tag       string
}

type GoError struct {
	Status int
	Code   string
	Desc   string
	Cause  string
	Fields []*FieldError
}

func (e *GoError) Error() string {
	if e.Cause != "" {
		return fmt.Sprintf("%s: %s - %s", e.Code, e.Desc, e.Cause)
	}

	return fmt.Sprintf("%s: %s", e.Code, e.Desc)
}

func (e *GoError) ErrorWithCause() string {
	return fmt.Sprintf("%s - %s", e.Error(), e.Cause)
}

func (e *GoError) IsCodeEqual(err error) bool {
	if ge, ok := err.(*GoError); ok {
		return ge.Code == e.Code
	}

	return false
}

func (e *GoError) AddErrorField(fieldName, desc, tag string) {
	e.Fields = append(e.Fields, &FieldError{
		FieldName: fieldName,
		Desc:      desc,
		Tag:       tag,
	})
}

func Set(status int, code string, desc string, cause string) *GoError {
	err := &GoError{
		Status: status,
		Code:   code,
		Desc:   desc,
		Cause:  cause,
	}
	return err
}

func New(desc string) *GoError {
	err := &GoError{
		Desc: desc,
	}

	return err
}
