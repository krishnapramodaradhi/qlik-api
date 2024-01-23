package util

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type (
	HandlerFunc func(w http.ResponseWriter, r *http.Request, db *sql.DB) (int, error)
	ApiError    struct {
		StatusCode int    `json:"statusCode"`
		Message    string `json:"message"`
		Data       any    `json:"data"`
	}
)

func NewApiError(statusCode int, data any) *ApiError {
	return &ApiError{StatusCode: statusCode, Message: getMessageFromStatusCode(statusCode), Data: data}
}

func WithDB(h HandlerFunc, db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		if r.Method == http.MethodOptions {
			return
		}
		if statusCode, err := h(w, r, db); err != nil {
			WriteJSON(w, statusCode, map[string]any{"type": "error", "result": err.Error()})
		}
	}
}

func WriteJSON(w http.ResponseWriter, statusCode int, data map[string]any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	result := data["result"]
	if data["type"] == "error" {
		result = NewApiError(statusCode, data["result"])
	}
	return json.NewEncoder(w).Encode(result)
}

func getMessageFromStatusCode(statusCode int) string {
	switch statusCode {
	case 404:
		return "Not Found"
	case 400:
		return "Bad Request"
	case 401:
		return "Unauthorized"
	case 403:
		return "Forbidden"
	case 500:
		return "Internal Server Error"
	default:
		return "Success"
	}
}
