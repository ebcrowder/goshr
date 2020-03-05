package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/ebcrowder/goshr/db"
)

type Handlers struct {
	redis *db.Redis
}

func (h *Handlers) postFiles(w http.ResponseWriter, r *http.Request) {
	// b, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	// 	responseError(w, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	// var file schema.File
	// if err := json.Unmarshal(b, &file); err != nil {
	// 	responseError(w, http.StatusBadRequest, err.Error())
	// 	return
	// }

	// id, err := h.redis.Insert(&file)
	// if err != nil {
	// 	responseError(w, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	// handle file
	// r.ParseMultipartForm(32 << 20)

	// fileBytes, handler, err := r.FormFile("myFile")
	// if err != nil {
	// 	responseError(w, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	r.ParseMultipartForm(32 << 20)

	log.Println(r.Form)

	fileBytes, handler, err := r.FormFile("myFile")

	defer fileBytes.Close()

	f, err := os.OpenFile(handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		fmt.Println(err)
	}

	io.Copy(f, fileBytes)

	responseOk(w, 1)
}

func (h *Handlers) deleteFiles(w http.ResponseWriter, r *http.Request) {
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

	if err := h.redis.Delete(req.ID); err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)

}

func (h *Handlers) getFiles(w http.ResponseWriter, r *http.Request) {
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

	val, err := h.redis.GetFiles(req.ID)
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
