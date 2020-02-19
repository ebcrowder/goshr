package handlers

import (
	"net/http"

	"github.com/ebcrowder/goshr/db"
)

// SetUpRoutes sets up server routes
func SetUpRoutes(redis *db.Redis) *http.ServeMux {
	shrHandlers := &shrHandlers{
		redis: redis,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			shrHandlers.getFiles(w, r)
		case http.MethodPost:
			shrHandlers.postFiles(w, r)
		case http.MethodDelete:
			shrHandlers.deleteFiles(w, r)
		default:
			responseError(w, http.StatusNotFound, "")
		}
	})
	return mux
}
