package app

import (
	"net/http"

	"github.com/go-chi/render"
	validation "github.com/go-ozzo/ozzo-validation"
)

// ErrResponse rendered type for handlling all sorts of errors
type ErrResponse struct {
	Err            error `json:"-"`
	HTTPStatusCode int   `json:"-"`

	StatusText       string            `json:"status"`
	AppCode          int64             `json:"code,omitempty"`
	ErrorText        string            `json:"error,omitempty"`
	ValidationErrors validation.Errors `json:"errors,omitempty"`
}

// Render sets the application-specific error code in AppCode
func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

// ErrInvalidRequest returns status 422 Unprocessable Entity including error messages
func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusUnprocessableEntity,
		StatusText:     http.StatusText(http.StatusUnprocessableEntity),
		ErrorText:      err.Error(),
	}
}

// ErrValidation returns status 422 Unprocessable Entity stating validation errors
func ErrValidation(err error, valErr validation.Errors) render.Renderer {
	return &ErrResponse{
		Err:              err,
		HTTPStatusCode:   http.StatusUnprocessableEntity,
		ErrorText:        err.Error(),
		ValidationErrors: valErr,
	}
}

// ErrRender returns status 422 Unprocessable Entity rendering response error
func ErrRender(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusUnprocessableEntity,
		StatusText:     http.StatusText(http.StatusUnprocessableEntity),
		ErrorText:      err.Error(),
	}
}

var (
	// ErrBadRequest returns status 400 Bad Request for malformed request body
	ErrBadRequest = &ErrResponse{HTTPStatusCode: http.StatusBadRequest, StatusText: http.StatusText(http.StatusBadRequest)}

	// ErrUnauthorized returns 401 Unauthorized
	ErrUnauthorized = &ErrResponse{HTTPStatusCode: http.StatusUnauthorized, StatusText: http.StatusText(http.StatusUnauthorized)}

	// ErrForbidden returns status 403 Forbidden for unauthorized request
	ErrForbidden = &ErrResponse{HTTPStatusCode: http.StatusForbidden, StatusText: http.StatusText(http.StatusForbidden)}

	// ErrNotFound returns status 404 Not Found for invalid resource request
	ErrNotFound = &ErrResponse{HTTPStatusCode: http.StatusNotFound, StatusText: http.StatusText(http.StatusNotFound)}

	// ErrInternalServerError returns status 500 Internal Server Error
	ErrInternalServerError = &ErrResponse{HTTPStatusCode: http.StatusInternalServerError, StatusText: http.StatusText(http.StatusInternalServerError)}
)
