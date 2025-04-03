package utils

import (
	"encoding/json"
	"net/http"
)

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error string `json:"error"`
}

// LoginResponse
type LoginSuccessResponse struct {
	Message string `json:"message"`
}

// RegisterResponse represents a success response
type RegisterSuccessResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

// WriteJSON writes a JSON response to the http.ResponseWriter
func WriteJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if v != nil {
		json.NewEncoder(w).Encode(v)
	}
}

func GetTokenFromRequest(r *http.Request) string {
	tokenAuth := r.Header.Get("Authorization")
	tokenQuery := r.URL.Query().Get("token")

	if tokenAuth != "" {
		return tokenAuth
	}

	if tokenQuery != "" {
		return tokenQuery
	}

	return ""
}
