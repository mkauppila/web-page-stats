package main

import (
	"fmt"
	"log"
	"os"

	"net/http"

	"github.com/mkauppila/web-page-stats/internal/api"
	"github.com/mkauppila/web-page-stats/internal/counters"
	"github.com/mkauppila/web-page-stats/internal/database"
	"github.com/mkauppila/web-page-stats/internal/handler"
)

func main() {
	fmt.Println("http server starting")

	dbUrl := os.Getenv("DATABASE_URL")
	fmt.Println("database url:", dbUrl)

	db, err := database.CreateConnection(dbUrl)
	if err != nil {
		panic(err)
	}
	vc := counters.CreateViewCounter(db)
	rc := counters.CreateReactionCounter(db)
	handler := handler.NewHandler(vc, rc)

	mux := http.NewServeMux()
	si := api.NewStrictHandler(handler, nil)
	hmux := api.HandlerFromMux(si, mux)

	s := &http.Server{
		Handler: hmux,
		Addr:    "0.0.0.0:8080",
	}

	// And we serve HTTP until the world ends.
	log.Fatal(s.ListenAndServe())
}
