package models

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Payload interface{} `json:"payload,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func ResponseJSON(w http.ResponseWriter, response interface{}) {
	json.NewEncoder(w).Encode(response)
}

func MethodNotAllowed(w http.ResponseWriter) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Header().Add("Content-Type", "application/json")

	response := Response{
		Error: "MethodNotAllowed",
	}
	ResponseJSON(w, response)
}
