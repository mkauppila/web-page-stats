package main

import (
	"fmt"
	"log"

	"net/http"

	"github.com/mkauppila/web-page-stats/internal/api"
	"github.com/mkauppila/web-page-stats/internal/database"
	"github.com/mkauppila/web-page-stats/internal/handler"
)

func main() {
	fmt.Println("http server starting")

	handler := handler.NewHandler()
	mux := http.NewServeMux()
	db, err := database.NewDatabase()
	if err != nil {
		log.Fatal(err)
	}
	db.Buu()

	// get an `http.Handler` that we can use
	h := api.HandlerFromMux(handler, mux)

	s := &http.Server{
		Handler: h,
		Addr:    "0.0.0.0:8080",
	}

	// And we serve HTTP until the world ends.
	log.Fatal(s.ListenAndServe())
}
