package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ebcrowder/goshr/db"
	"github.com/ebcrowder/goshr/schema"
)

type shrHandlers struct {
	redis *db.Redis
}

func (handlers *shrHandlers) postFiles(w http.ResponseWriter, r *http.Request) {
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

	id, err := handlers.redis.Insert(&file)
	if err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseOk(w, id)
}

func (handlers *shrHandlers) deleteFiles(w http.ResponseWriter, r *http.Request) {
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

	if err := handlers.redis.Delete(req.ID); err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)

}

func (handlers *shrHandlers) getFiles(w http.ResponseWriter, r *http.Request) {
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

	val, err := handlers.redis.GetFiles(req.ID)
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
