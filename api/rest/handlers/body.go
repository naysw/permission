package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func parseBody(r *http.Request, v interface{}) error {
	if r.Body == nil {
		return errors.New("request body is nil")
	}
	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		if err == io.EOF {
			return errors.New("request body is empty")
		}
		if _, ok := err.(*json.SyntaxError); ok {
			return fmt.Errorf("request body has invalid JSON syntax")
		}
		return fmt.Errorf("failed to decode request body: %v", err)
	}

	return nil
}
