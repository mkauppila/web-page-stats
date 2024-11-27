package main

import (
	"fmt"
	"log"

	"github.com/mkauppila/web-page-stats/internal/api"
	"net/http"
)

type Server struct{}

func NewServer() Server {
	return Server{}
}

// Get reaction counts
// (GET /reactions/{category}/{slug})
func (s Server) GetReactionsCategorySlug(
	w http.ResponseWriter,
	r *http.Request,
	category api.GetReactionsCategorySlugParamsCategory,
	slug string,
) {
	w.WriteHeader(200)
	w.Write([]byte("/reactoins/{category}/{slug}"))
}

// Increment reaction count
// (PUT /reactions/{reaction}/{category}/{slug})
func (s Server) PutReactionsReactionCategorySlug(
	w http.ResponseWriter,
	r *http.Request,
	reaction api.PutReactionsReactionCategorySlugParamsReaction,
	category api.PutReactionsReactionCategorySlugParamsCategory,
	slug string,
) {
}

// Get view count
// (GET /views/{category}/{slug})
func (s Server) GetViewsCategorySlug(
	w http.ResponseWriter,
	r *http.Request,
	category api.GetViewsCategorySlugParamsCategory,
	slug string,
) {
	w.WriteHeader(200)
	w.Write([]byte("/view/category/slug"))
}

// Increment view count
// (PUT /views/{category}/{slug})
func (s Server) PutViewsCategorySlug(
	w http.ResponseWriter,
	r *http.Request,
	category api.PutViewsCategorySlugParamsCategory,
	slug string,
) {
	w.WriteHeader(201)
}

func main() {
	fmt.Println("http server starting")
	// create a type that satisfies the `api.ServerInterface`, which contains an implementation of every operation from the generated code
	server := NewServer()

	r := http.NewServeMux()

	// get an `http.Handler` that we can use
	h := api.HandlerFromMux(server, r)

	s := &http.Server{
		Handler: h,
		Addr:    "0.0.0.0:8080",
	}

	// And we serve HTTP until the world ends.
	log.Fatal(s.ListenAndServe())
}
