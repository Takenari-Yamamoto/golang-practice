package errors

import (
	"github.com/vektah/gqlparser/v2/gqlerror"
)

type ErrorCode int

const (
	UNKNOWN ErrorCode = iota
	BAD_REQUEST
	UNAUTHORIZED
	FORBIDDEN
	NOT_FOUND
	UNPROCESSABLE_ENTITY
	INTERNAL_SERVER_ERROR
)

type CustomError struct {
	Code         string        `json:"code"`
	UserMessage  string        `json:"userMessage"`
	ErrorDetails []ErrorDetail `json:"errorDetails"`
}

type ErrorDetail struct {
	Message   string `json:"message"`
	Attribute string `json:"attribute"`
}

type GqlError struct {
	Message     string
	Code        ErrorCode
	UserMessage string
	Details     *[]ErrorDetail
}

func NewCustomError(err GqlError) *gqlerror.Error {
	return &gqlerror.Error{
		Message: err.Message,
		Extensions: map[string]interface{}{
			"code":         err.Code,
			"userMessage":  err.UserMessage,
			"errorDetails": &err.Details,
		},
	}
}
