package handlers

import (
	"net/http"

	"github.com/ebcrowder/goshr/db"
)

func SetUpRoutes(redis *db.Redis) *http.ServeMux {
	Handlers := &Handlers{
		redis: redis,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			Handlers.getFiles(w, r)
		case http.MethodPost:
			Handlers.postFiles(w, r)
		case http.MethodDelete:
			Handlers.deleteFiles(w, r)
		default:
			responseError(w, http.StatusNotFound, "")
		}
	})
	return mux
}
