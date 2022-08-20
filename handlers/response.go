package handlers

import (
	"encoding/json"
	"net/http"
)

func Response(w http.ResponseWriter, status string, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	response := map[string]interface{}{
		"status": status,
		"response": data,
	}
	json.NewEncoder(w).Encode(&response)
}