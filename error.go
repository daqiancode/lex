package lex

import (
	"fmt"

	"github.com/daqiancode/lex/ecode"
)

var AppName string

type CodeError struct {
	App       string `json:",omitempty"`
	ErrorCode int
	Message   string
}

func (s *CodeError) Error() string {
	return fmt.Sprintf("app:%s,code:%d,message:%s", s.App, s.ErrorCode, s.Message)
}

type FieldError struct {
	CodeError
	FieldErrors map[string]string
}

func NewCodeError(code int, message string) *CodeError {
	return &CodeError{ErrorCode: code, Message: message, App: AppName}
}

func NewFieldError(fieldErrors map[string]string) *FieldError {
	return &FieldError{CodeError: CodeError{App: AppName, ErrorCode: ecode.ParamInvalid, Message: "request parameters have errors"}, FieldErrors: fieldErrors}
}
func NewFieldErrorKV(kvs ...string) *FieldError {
	n := len(kvs) / 2
	r := make(map[string]string, n)
	for i := 0; i < n; i++ {
		r[kvs[i]] = kvs[i+1]
	}
	return NewFieldError(r)
}
