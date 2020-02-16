package main

import (
	"log"

	"net/http"

	"github.com/ebcrowder/goshr/db"
	"github.com/ebcrowder/goshr/handlers"
)

type apiHandler struct{}

func (apiHandler) ServeHTTP(http.ResponseWriter, *http.Request) {}

func main() {
	port := "8000"

	var sqlite *db.Sqlite
	var err error

	sqlite, err = db.ConnectSqlite()

	if err != nil {
		panic(err)
	} else if sqlite == nil {
		panic("sqlite is nil")
	}

	mux := handlers.SetUpRoutes(sqlite)

	log.Fatal(http.ListenAndServe(":"+port, mux))
}
