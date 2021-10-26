package lex

import (
	"fmt"

	"github.com/daqiancode/lex/stderrorcode"
)

var AppName string

type CodeError struct {
	App       string `json:",omitempty"`
	ErrorCode stderrorcode.ErrorCode
	Message   string
}

func (s *CodeError) Error() string {
	return fmt.Sprintf("app:%s,code:%d,message:%s", s.App, s.ErrorCode, s.Message)
}

type FieldError struct {
	CodeError
	FieldErrors map[string]string
}

func NewCodeError(code stderrorcode.ErrorCode, message string) *CodeError {
	return &CodeError{ErrorCode: code, Message: message, App: AppName}
}

func NewFieldError(fieldErrors map[string]string) *FieldError {
	return &FieldError{CodeError: CodeError{App: AppName, ErrorCode: stderrorcode.FieldError, Message: "request parameters have errors"}, FieldErrors: fieldErrors}
}
func NewFieldErrorKV(kvs ...string) *FieldError {
	n := len(kvs) / 2
	r := make(map[string]string, n)
	for i := 0; i < n; i++ {
		r[kvs[i]] = kvs[i+1]
	}
	return NewFieldError(r)
}

// func makeFieldErrorMessage(err validator.FieldError) string {
// 	// tag := err.Tag()
// 	// kind := err.Kind().String()

// 	return "Please recheck it."
// }

// func NewFormError(err error) error {
// 	if errs, ok := err.(validator.ValidationErrors); ok {
// 		r := make(map[string]string, len(errs))
// 		for _, e := range errs {
// 			r[e.Field()] = makeFieldErrorMessage(e)
// 		}
// 		return NewFieldError(r)
// 	}
// 	return err
// }
