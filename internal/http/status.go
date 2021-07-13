package http

import (
	"fmt"
	"net/http"
)

// HandleStatus maps a status to the corresponding error
func HandleStatus(status int, body []byte) error {
	var err error

	switch status {
	case http.StatusUnauthorized:
		err = UnauthorizedErr{}
	case http.StatusNotFound:
		err = ErrNotFound{}
	case http.StatusBadRequest:
		err = BadRequestErr{}
	case http.StatusInternalServerError:
		err = InternalServerErr{}
	default:
	}

	fmt.Printf("Status: %v\nError: %v\n Body:%v\n", status, err, string(body))

	return err
}

// UnauthorizedErr thrown when the user is unauthorized
type UnauthorizedErr struct{}

func (err UnauthorizedErr) Error() string {
	return "unauthorized"
}

// BadRequestErr thrown when the request is bad
type BadRequestErr struct{}

func (err BadRequestErr) Error() string {
	return "bad request"
}

// InternalServerErr thrown when there is an internal server error
type InternalServerErr struct{}

func (err InternalServerErr) Error() string {
	return "internal server error"
}

// ErrNotFound thrown when not found
type ErrNotFound struct{}

func (err ErrNotFound) Error() string {
	return "not found"
}
