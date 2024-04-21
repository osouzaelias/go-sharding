package api

import (
	"encoding/json"
	"net/http"
)

type Success struct {
	statusCode int
	result     interface{}
}

func NewSuccess(result interface{}, status int) Success {
	return Success{
		statusCode: status,
		result:     result,
	}
}

func (r Success) Send(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.statusCode)
	_ = json.NewEncoder(w).Encode(r.result)
}
