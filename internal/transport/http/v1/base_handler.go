package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type BaseHandler struct {
}

func NewBaseHandler() *BaseHandler {

	return &BaseHandler{}
}

func (h *BaseHandler) WriteJsonResponse(w http.ResponseWriter, data any, statusCode int) error {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(&data)
	if err != nil {
		return fmt.Errorf("failed to write json response: %v", err)
	}

	return nil
}
