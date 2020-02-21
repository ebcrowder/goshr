package main

import (
	"log"

	"net/http"

	"github.com/ebcrowder/goshr/db"
	"github.com/ebcrowder/goshr/handlers"
)

func main() {
	port := "8000"

	var redis *db.Redis
	var err error

	redis, err = db.ConnectRedis()

	if err != nil {
		panic(err)
	} else if redis == nil {
		panic("redis is nil")
	}

	mux := handlers.SetUpRoutes(redis)

	log.Fatal(http.ListenAndServe(":"+port, mux))
}
