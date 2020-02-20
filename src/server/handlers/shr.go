package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ebcrowder/goshr/db"
	"github.com/ebcrowder/goshr/schema"
	"github.com/ebcrowder/goshr/service"
)

type shrHandlers struct {
	redis *db.Redis
}

func (handlers *shrHandlers) postFiles(w http.ResponseWriter, r *http.Request) {
	ctx := db.SetRepository(r.Context(), handlers.redis)

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var file schema.File
	if err := json.Unmarshal(b, &file); err != nil {
		responseError(w, http.StatusBadRequest, err.Error())
		return
	}

	id, err := service.Insert(ctx, &file)
	if err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseOk(w, id)
}

func (handlers *shrHandlers) deleteFiles(w http.ResponseWriter, r *http.Request) {
	ctx := db.SetRepository(r.Context(), handlers.redis)

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var req struct {
		ID string `json:"id"`
	}

	if err := json.Unmarshal(b, &req); err != nil {
		responseError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := service.Delete(ctx, req.ID); err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)

}

func (handlers *shrHandlers) getFiles(w http.ResponseWriter, r *http.Request) {
	ctx := db.SetRepository(r.Context(), handlers.redis)

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var req struct {
		ID string `json:"id"`
	}

	if err := json.Unmarshal(b, &req); err != nil {
		responseError(w, http.StatusBadRequest, err.Error())
		return
	}

	val, err := service.GetFiles(ctx, req.ID)
	if err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseOk(w, val)
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
