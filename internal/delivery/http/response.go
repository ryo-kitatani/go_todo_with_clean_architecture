package http

import (
	"encoding/json"
	"net/http"
	"todo-api/internal/domain"
)

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, domain.ErrorResponse{Message: message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
