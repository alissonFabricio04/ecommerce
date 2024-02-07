package utils

import (
	"encoding/json"
	"net/http"
)

func SendResponse(w http.ResponseWriter, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
