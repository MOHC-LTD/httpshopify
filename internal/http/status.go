package http

import (
	"fmt"
	"net/http"
)

// HandleStatus maps a status to the corresponding error
func HandleStatus(status int, body []byte) error {
	if status >= http.StatusBadRequest {
		err := NewErrHTTP(
			status,
			string(body),
		)

		fmt.Println(err)

		return err
	}

	return nil
}

// ErrHTTP thrown when the http error code is >= http.StatusBadRequest
type ErrHTTP struct {
	Code int
	Body string
}

func (err ErrHTTP) Error() string {
	return fmt.Sprintf("%v %v", err.Code, err.Body)
}

// NewErrHTTP builds the error
func NewErrHTTP(code int, body string) ErrHTTP {
	return ErrHTTP{code, body}
}
