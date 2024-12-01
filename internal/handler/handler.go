package handler

import (
	"context"
	"encoding/json"
	"fmt"
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
	GetCount(ctx context.Context, path string) (ReactionCounts, error)
	Update(ctx context.Context, path, reaction string) (ReactionCounts, error)
}

type Viewer interface {
	GetCount(ctx context.Context, path string) (ViewCount, error)
	Update(ctx context.Context, path string) (ViewCount, error)
}

func naiveAuthVerifier(r *http.Request, validToken string) bool {
	bearer := r.Header.Get("Authorization")
	return bearer == "Bearer "+validToken
}

type Handler struct {
	views     Viewer
	reactions Reactioner
	authToken string
}

var _ api.ServerInterface = Handler{}

func NewHandler(views Viewer, reactions Reactioner, authToken string) Handler {
	return Handler{
		views:     views,
		reactions: reactions,
		authToken: authToken,
	}
}

func (s Handler) GetReactions(
	w http.ResponseWriter,
	r *http.Request,
	params api.GetReactionsParams,
) {
	if !naiveAuthVerifier(r, s.authToken) {
		w.WriteHeader(403)
		return
	}

	counts, err := s.reactions.GetCount(
		r.Context(),
		string(params.Path),
	)
	if err != nil {
		fmt.Println(err)
	}

	resp := api.GetReactions200JSONResponse{
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

func (s Handler) PutReactions(
	w http.ResponseWriter,
	r *http.Request,
	params api.PutReactionsParams,
) {
	if !naiveAuthVerifier(r, s.authToken) {
		w.WriteHeader(403)
		return
	}

	counts, err := s.reactions.Update(
		r.Context(),
		string(params.Path),
		string(params.Reaction),
	)
	if err != nil {
		panic(err)
	}
	resp := api.PutReactions200JSONResponse{
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

func (s Handler) GetViews(
	w http.ResponseWriter,
	r *http.Request,
	params api.GetViewsParams,
) {
	if !naiveAuthVerifier(r, s.authToken) {
		w.WriteHeader(403)
		return
	}

	counts, err := s.views.GetCount(
		r.Context(),
		string(params.Path),
	)
	resp := api.GetViews200JSONResponse{
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

func (s Handler) PutViews(
	w http.ResponseWriter,
	r *http.Request,
	params api.PutViewsParams,
) {
	if !naiveAuthVerifier(r, s.authToken) {
		w.WriteHeader(403)
		return
	}

	counts, err := s.views.Update(
		r.Context(),
		string(params.Path),
	)
	resp := api.PutViews200JSONResponse{
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
