package problems

import (
	"fmt"
	"time"

	"github.com/google/jsonapi"
)

type ErrorInput struct {
	// Unique identifier for the error
	ErrorID string

	// Unique identifier for the request (correlation ID)
	RequestID string

	// HTTP status code applicable to this problem, as a string ("400", "500" и т.д.)
	Status int

	// Application-specific error code
	Code string

	// Short, human-readable summary (title) of the problem
	Title string

	// Detailed explanation of this occurrence of the problem
	Detail string

	// JSON Pointer to the value in the request document that caused the error
	Pointer string
}

func Error(input ErrorInput) []*jsonapi.ErrorObject {
	now := time.Now().UTC().Format(time.RFC3339)

	statusStr := fmt.Sprintf("%d", input.Status)

	meta := &map[string]interface{}{
		"timestamp": now,
	}

	if input.RequestID != "" {
		(*meta)["request_id"] = input.RequestID
	}

	if input.Pointer != "" {
		(*meta)["pointer"] = input.Pointer
	}

	eo := &jsonapi.ErrorObject{
		ID:     input.ErrorID,
		Status: statusStr,
		Code:   input.Code,
		Title:  input.Title,
		Detail: input.Detail,
		Meta:   meta,
	}

	return []*jsonapi.ErrorObject{eo}
}
