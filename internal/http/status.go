package http

import "net/http"

// HandleStatus maps a status to the corresponding error
func HandleStatus(status int) error {
	switch status {
	case http.StatusUnauthorized:
		return UnauthorizedErr{}
	case http.StatusBadRequest:
		return BadRequestErr{}
	case http.StatusInternalServerError:
		return InternalServerErr{}
	default:
		return nil
	}
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
