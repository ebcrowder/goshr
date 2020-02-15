package handlers

import (
	"encoding/json"
	"net/http"
)

type shrHandlers struct{}

func (handlers *shrHandlers) postFiles(w http.ResponseWriter, r *http.Request) {
	responseOk(w, "hi")
}

func (handlers *shrHandlers) deleteFiles(w http.ResponseWriter, r *http.Request) {

	responseOk(w, "hi")

}

func (handlers *shrHandlers) getFiles(w http.ResponseWriter, r *http.Request) {
	responseOk(w, "hi")
}

func responseOk(w http.ResponseWriter, body interface{}) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(body)
}

func responseError(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")

	body := map[string]string{
		"error": message,
	}
	json.NewEncoder(w).Encode(body)
}
