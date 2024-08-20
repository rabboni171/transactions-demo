package handler

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) respondWithError(w http.ResponseWriter, statusCode int, message string) {
	response := map[string]string{"error": message}
	jsonResponse, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(jsonResponse)
}

func (h *Handler) respondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	jsonResponse, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(jsonResponse)
}
