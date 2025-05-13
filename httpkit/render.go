package httpkit

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/google/jsonapi"
	"github.com/hs-zavet/comtools/httpkit/problems"
	"github.com/pkg/errors"
)

// Render encodes a successful response in JSON API format.
func Render(w http.ResponseWriter, res interface{}) {
	w.Header().Set("content-type", jsonapi.MediaType)
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		panic(errors.Wrap(err, "failed to render response"))
	}
}

func RenderErr(w http.ResponseWriter, errs ...*jsonapi.ErrorObject) {
	if len(errs) == 0 {
		RenderErr(w, problems.Error(problems.ErrorInput{
			Status: http.StatusInternalServerError,
		})...)
		return
	}

	status, err := strconv.Atoi(errs[0].Status)
	if err != nil {
		RenderErr(w, problems.Error(problems.ErrorInput{
			Status: http.StatusInternalServerError,
		})...)
		return
	}

	w.Header().Set("Content-Type", jsonapi.MediaType)
	w.WriteHeader(status)

	if err := jsonapi.MarshalErrors(w, errs); err != nil {
		log.Printf("Failed to marshal errors: %v", err)
	}
}
