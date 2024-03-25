package utils

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func EncodeJSON() {

}

type result struct {
	Message  string
	HttpCode int
}

func DecodedJSON(r io.Reader, v interface{}) result {
	var unmarshalErr *json.UnmarshalTypeError

	decoder := json.NewDecoder(r)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&v)
	if err != nil {
		if errors.As(err, &unmarshalErr) {
			return result{
				Message:  "Bad Request. Wrong Type provided for field " + unmarshalErr.Field,
				HttpCode: http.StatusBadRequest,
			}
		}
		return result{
			Message:  "Bad Request. Wrong Type provided for field " + unmarshalErr.Field,
			HttpCode: http.StatusBadRequest,
		}
	}

	return result{}
}
