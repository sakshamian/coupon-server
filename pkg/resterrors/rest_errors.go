package resterrors

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"runtime"
)

type RestErr interface {
	Message() string
	Status() int
	Error() string
	Causes() []interface{}
	Stack() string
}

type restErr struct {
	ErrMessage string        `json:"message"`
	ErrStatus  int           `json:"status"`
	ErrError   string        `json:"error"`
	ErrCauses  []interface{} `json:"-"`
	ErrStack   string        `json:"-"`
}

func (e restErr) Error() string {
	return fmt.Sprintf("message: %s \tstatus: %d \terror: %s \ncauses: %v \nstack: %s",
		e.ErrMessage, e.ErrStatus, e.ErrError, e.ErrCauses, e.ErrStack)
}

func (e restErr) Message() string {
	return e.ErrMessage
}

func (e restErr) Status() int {
	return e.ErrStatus
}

func (e restErr) Causes() []interface{} {
	return e.ErrCauses
}

func (e restErr) Stack() string {
	return e.ErrStack
}

func NewRestError(message string, status int, err string, causes []interface{}) RestErr {
	_, file, line, _ := runtime.Caller(1)

	result := restErr{
		ErrMessage: message,
		ErrStatus:  status,
		ErrError:   err,
		ErrCauses:  causes,
		ErrStack:   fmt.Sprintf("%s:%d", file, line),
	}

	if os.Getenv("DEBUG") == "true" {
		fmt.Println(result)
	}

	return result
}

func NewRestErrorFromBytes(bytes []byte) (RestErr, error) {
	var apiErr restErr
	if err := json.Unmarshal(bytes, &apiErr); err != nil {
		return nil, errors.New("invalid json")
	}

	return apiErr, nil
}

func NewBadRequestError(message string) RestErr {
	_, file, line, _ := runtime.Caller(1)

	result := restErr{
		ErrMessage: message,
		ErrStatus:  http.StatusBadRequest,
		ErrError:   ERROR_BAD_REQUEST,
		ErrStack:   fmt.Sprintf("%s:%d", file, line),
	}

	if os.Getenv("DEBUG") == "true" {
		fmt.Println(result)
	}

	return result
}

func NewConflictError(message string) RestErr {
	_, file, line, _ := runtime.Caller(1)

	result := restErr{
		ErrMessage: message,
		ErrStatus:  http.StatusConflict,
		ErrError:   ERROR_CONFLICT,
		ErrStack:   fmt.Sprintf("%s:%d", file, line),
	}

	if os.Getenv("DEBUG") == "true" {
		fmt.Println(result)
	}

	return result
}

func NewNotFoundError(message string) RestErr {
	_, file, line, _ := runtime.Caller(1)

	result := restErr{
		ErrMessage: message,
		ErrStatus:  http.StatusNotFound,
		ErrError:   ERROR_NOT_FOUND,
		ErrStack:   fmt.Sprintf("%s:%d", file, line),
	}

	if os.Getenv("DEBUG") == "true" {
		fmt.Println(result)
	}

	return result
}

func NewUnauthorizedError(message string) RestErr {
	_, file, line, _ := runtime.Caller(1)

	result := restErr{
		ErrMessage: message,
		ErrStatus:  http.StatusUnauthorized,
		ErrError:   ERROR_UNAUTHORIZED,
		ErrStack:   fmt.Sprintf("%s:%d", file, line),
	}

	if os.Getenv("DEBUG") == "true" {
		fmt.Println(result)
	}

	return result
}

func NewInternalServerError(message string, err error) RestErr {
	_, file, line, _ := runtime.Caller(1)

	result := restErr{
		ErrMessage: message,
		ErrStatus:  http.StatusInternalServerError,
		ErrError:   ERROR_SERVER,
		ErrStack:   fmt.Sprintf("%s:%d", file, line),
	}
	if err != nil {
		result.ErrCauses = append(result.ErrCauses, err.Error())
	}

	if os.Getenv("DEBUG") == "true" {
		fmt.Println(result)
	}

	return result
}

func NewForbiddenError(message string) RestErr {
	_, file, line, _ := runtime.Caller(1)

	result := restErr{
		ErrMessage: message,
		ErrStatus:  http.StatusForbidden,
		ErrError:   ERROR_FORBIDDEN,
		ErrStack:   fmt.Sprintf("%s:%d", file, line),
	}

	if os.Getenv("DEBUG") == "true" {
		fmt.Println(result)
	}

	return result
}

func NewPaymentRequiredError(message string) RestErr {
	_, file, line, _ := runtime.Caller(1)

	result := restErr{
		ErrMessage: message,
		ErrStatus:  http.StatusPaymentRequired,
		ErrError:   http.StatusText(http.StatusPaymentRequired),
		ErrStack:   fmt.Sprintf("%s:%d", file, line),
	}

	if os.Getenv("DEBUG") == "true" {
		fmt.Println(result)
	}

	return result
}
