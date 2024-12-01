package handler

import (
	"encoding/json"
	"net/http"

	"github.com/mkauppila/web-page-stats/internal/api"
)

type ReactionCounts struct {
	Like      int `json:"like"`
	Love      int `json:"love"`
	Mindblown int `json:"mindblown"`
	Puzzling  int `json:"puzzling"`
}

type ViewCount struct {
	Count int `json:"count"`
}

type Reactioner interface {
	GetCount(category, slug string) (ReactionCounts, error)
	Update(category, slug, reaction string) (ReactionCounts, error)
}

type Viewer interface {
	GetCount(category, slug string) (ViewCount, error)
	Update(category, slug string) (ViewCount, error)
}

type Handler struct {
	views     Viewer
	reactions Reactioner
}

func NewHandler(views Viewer, reactions Reactioner) Handler {
	return Handler{
		views:     views,
		reactions: reactions,
	}
}

// Get reaction counts
// (GET /reactions/{category}/{slug})
func (s Handler) GetReactionsCategorySlug(
	w http.ResponseWriter,
	r *http.Request,
	category api.GetReactionsCategorySlugParamsCategory,
	slug string,
) {
	counts, err := s.reactions.GetCount(string(category), slug)
	resp := api.GetReactionsCategorySlug200JSONResponse{
		Like:      &counts.Like,
		Love:      &counts.Love,
		Mindblown: &counts.Mindblown,
		Puzzling:  &counts.Puzzling,
	}

	switch err {
	case nil:
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(resp)
	default:
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(resp)
	}
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
	counts, err := s.reactions.Update(string(category), slug, string(reaction))
	resp := api.PutReactionsReactionCategorySlug200JSONResponse{
		Like:      &counts.Like,
		Love:      &counts.Love,
		Mindblown: &counts.Mindblown,
		Puzzling:  &counts.Puzzling,
	}

	switch err {
	case nil:
		w.WriteHeader(201)
		json.NewEncoder(w).Encode(resp)
	default:
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(resp)
	}
}

// Get view count
// (GET /views/{category}/{slug})
func (s Handler) GetViewsCategorySlug(
	w http.ResponseWriter,
	r *http.Request,
	category api.GetViewsCategorySlugParamsCategory,
	slug string,
) {
	counts, err := s.views.GetCount(string(category), slug)
	resp := api.GetViewsCategorySlug200JSONResponse{
		Views: &counts.Count,
	}

	switch err {
	case nil:
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(resp)
	default:
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(resp)
	}
}

// Increment view count
// (PUT /views/{category}/{slug})
func (s Handler) PutViewsCategorySlug(
	w http.ResponseWriter,
	r *http.Request,
	category api.PutViewsCategorySlugParamsCategory,
	slug string,
) {
	counts, err := s.views.Update(string(category), slug)
	resp := api.PutViewsCategorySlug200JSONResponse{
		Views: &counts.Count,
	}

	switch err {
	case nil:
		w.WriteHeader(201)
		json.NewEncoder(w).Encode(resp)
	default:
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(resp)
	}
}
