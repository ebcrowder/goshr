package main

import (
	"log"

	"github.com/ebcrowder/goshr/handlers"
	"net/http"
)

type apiHandler struct{}

func (apiHandler) ServeHTTP(http.ResponseWriter, *http.Request) {}

func main() {
	port := "8000"

	mux := handlers.SetUpRoutes()

	log.Fatal(http.ListenAndServe(":"+port, mux))
}
