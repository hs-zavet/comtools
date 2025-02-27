package problems

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/jsonapi"
	"github.com/pkg/errors"
)

type errBadRequest map[string]error

// Error returns a concatenated errors in string format.
func (e errBadRequest) Error() string {
	var msg strings.Builder
	var sep = ", "

	for f, err := range e {
		msg.WriteString(sep)
		msg.WriteString(fmt.Sprintf("%s: %s", f, err.Error()))
	}

	return strings.TrimPrefix(msg.String(), sep)
}

// Filter returns errBadRequest if it contains any errors, nil -
// otherwise. Do return errs.Filter() when populating errors under for loop.
func (e errBadRequest) Filter() error {
	for k, err := range e {
		if err == nil {
			delete(e, k)
		}
	}

	if len(e) == 0 {
		return nil
	}

	return e
}

// BadRequest returns a message to be rendered to client in
// case it's request is invalid.
func (e errBadRequest) BadRequest() map[string]error {
	return e
}

// BadRequester is an error that indicates bad request.
type BadRequester interface {
	BadRequest() map[string]error
}

func BadRequest(err error) []*jsonapi.ErrorObject {
	cause := errors.Cause(err)
	if cause == io.EOF {
		return []*jsonapi.ErrorObject{
			{
				Title:  http.StatusText(http.StatusBadRequest),
				Status: fmt.Sprintf("%d", http.StatusBadRequest),
				Detail: "Request body were expected",
			},
		}
	}

	switch cause := cause.(type) {
	case validation.Errors:
		return toJsonapiErrors(cause)
	case BadRequester:
		return toJsonapiErrors(cause.BadRequest())
	default:
		return []*jsonapi.ErrorObject{
			{
				Title:  http.StatusText(http.StatusBadRequest),
				Status: fmt.Sprintf("%d", http.StatusBadRequest),
				Detail: "Your request was invalid in some way",
			},
		}
	}
}
