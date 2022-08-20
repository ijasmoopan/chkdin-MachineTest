package user

import (
	"encoding/json"
	"net/http"
)

func responseOK(w http.ResponseWriter, message interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"status":   true,
		"response": message,
	}
	json.NewEncoder(w).Encode(&response)
}

func responseBadRequest(w http.ResponseWriter, message interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	response := map[string]interface{}{
		"status":   false,
		"response": message,
	}
	json.NewEncoder(w).Encode(&response)
}