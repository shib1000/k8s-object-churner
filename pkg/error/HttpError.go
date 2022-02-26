package error

import (
	"encoding/json"
	"net/http"
)

type HttpError struct {
	statusCode int
	message    string
	err        error
}

func NewHttpError() *HttpError {
	return NewHttpErrorWithMessage(http.StatusInternalServerError,
		"Could not process request due to Unknown Cause.")
}

func NewHttpErrorWithStatus(statusCode int) *HttpError {
	return NewHttpErrorWithMessage(statusCode,
		"Could not process request due to Unknown Cause.")
}

func NewHttpErrorWithMessage(statusCode int, message string) *HttpError {
	return &HttpError{
		statusCode: statusCode,
		message:    message,
	}
}

func NewHttpErrorWithError(statusCode int, err error) *HttpError {
	return &HttpError{
		statusCode: statusCode,
		err:        err,
	}
}

func (hr *HttpError) Error() (str string) {
	bytes, err := json.Marshal(hr)
	if err != nil {
		str = "failed to serialize Http Error"
	} else {
		str = string(bytes)
	}
	return
}

func (hr *HttpError) Code() (code int) {
	return hr.statusCode
}
