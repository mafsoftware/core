package core

import (
	"encoding/json"
	"fmt"
)

type Error struct {
	Code        int               `json:"domainCode"`
	StatusCode  int               `json:"statusCode"`
	Description string            `json:"description"`
	Info        map[string]string `json:"info"`
}

func (e *Error) Error() string {
	return e.JSON()
}

// IsEqual compares itself with other error. This method compares all fields except Info.
func (e *Error) IsEqual(other Error) bool {
	return e.Code == other.Code && e.StatusCode == other.StatusCode && e.Description == other.Description
}

// JSON returns a json representation of the error.
func (e *Error) JSON() string {
	mJson, err := json.Marshal(e)
	if err != nil {
		s := "{\"code\": \"%v\", \"statusCode\": \"%v\", \"description\": \"%v\"}"
		return fmt.Sprintf(s, e.Code, e.StatusCode, e.Description)
	}
	return string(mJson)
}
