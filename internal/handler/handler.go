package handler

import (
	"context"

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

type Handler struct {
	views     Viewer
	reactions Reactioner
	authToken string
}

func NewHandler(views Viewer, reactions Reactioner) Handler {
	return Handler{
		views:     views,
		reactions: reactions,
		authToken: "",
	}
}

var _ api.StrictServerInterface = Handler{}

func (s Handler) GetReactions(
	ctx context.Context,
	request api.GetReactionsRequestObject,
) (api.GetReactionsResponseObject, error) {
	counts, err := s.reactions.GetCount(ctx, request.Params.Path)
	if err != nil {
		return nil, err
	}

	return api.GetReactions200JSONResponse{
		Like:      &counts.Like,
		Love:      &counts.Love,
		Mindblown: &counts.Mindblown,
		Puzzling:  &counts.Puzzling,
	}, nil
}

func (s Handler) PutReactions(ctx context.Context, request api.PutReactionsRequestObject) (api.PutReactionsResponseObject, error) {
	counts, err := s.reactions.Update(ctx, request.Params.Path, string(request.Params.Reaction))
	if err != nil {
		return nil, err
	}

	return api.PutReactions200JSONResponse{
		Like:      &counts.Like,
		Love:      &counts.Love,
		Mindblown: &counts.Mindblown,
		Puzzling:  &counts.Puzzling,
	}, nil
}

func (s Handler) GetViews(ctx context.Context, request api.GetViewsRequestObject) (api.GetViewsResponseObject, error) {
	counts, err := s.views.GetCount(ctx, request.Params.Path)
	if err != nil {
		return nil, err
	}

	return api.GetViews200JSONResponse{
		Views: &counts.Count,
	}, err
}

func (s Handler) PutViews(ctx context.Context, request api.PutViewsRequestObject) (api.PutViewsResponseObject, error) {
	counts, err := s.views.Update(ctx, request.Params.Path)
	if err != nil {
		return nil, err
	}

	return api.PutViews200JSONResponse{
		Views: &counts.Count,
	}, err
}
