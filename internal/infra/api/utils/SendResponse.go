package utils

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Data    any
	Message string `json:"message"`
}

func SendResponse(w http.ResponseWriter, response Response) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
