package util

import (
	"fmt"
)

type codeError struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

type CodeError interface {
	ErrCode() int
}

func EmptyCodeErr() CodeError {
	return new(codeError)
}

func NewCodeError(code int, format string, a ...interface{}) CodeError {
	return &codeError{Code: code, Text: fmt.Sprintf(format, a...)}
}

func (ce *codeError) ErrCode() int {
	return ce.Code
}

func (ce *codeError) String() string {
	return fmt.Sprintf("error code: %d error text: %s", ce.Code, ce.Text)
}

func (ce *codeError) Error() string {
	return ce.Text
}

// Runs the given function and converts any panic encountered while doing so
// into an error. Useful for sending to channels that will close
func PanicToError(fn func()) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v", r)
		}
	}()
	fn()
	return
}
