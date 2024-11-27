package handler

import (
	"net/http"

	"github.com/mkauppila/web-page-stats/internal/api"
)

type Handler struct{}

func NewHandler() Handler {
	return Handler{}
}

// Get reaction counts
// (GET /reactions/{category}/{slug})
func (s Handler) GetReactionsCategorySlug(
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
func (s Handler) PutReactionsReactionCategorySlug(
	w http.ResponseWriter,
	r *http.Request,
	reaction api.PutReactionsReactionCategorySlugParamsReaction,
	category api.PutReactionsReactionCategorySlugParamsCategory,
	slug string,
) {
}

// Get view count
// (GET /views/{category}/{slug})
func (s Handler) GetViewsCategorySlug(
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
func (s Handler) PutViewsCategorySlug(
	w http.ResponseWriter,
	r *http.Request,
	category api.PutViewsCategorySlugParamsCategory,
	slug string,
) {
	w.WriteHeader(201)
}
