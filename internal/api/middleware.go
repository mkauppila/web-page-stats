package api

import (
	"context"
	"errors"
	"fmt"
	"net/http"
)

func CreateRequestLogger() StrictMiddlewareFunc {
	return func(f StrictHandlerFunc, operationID string) StrictHandlerFunc {
		return func(ctx context.Context, w http.ResponseWriter, r *http.Request, input interface{}) (interface{}, error) {
			fmt.Println("Req: ", r.URL)
			return f(ctx, w, r, input)
		}
	}
}

func CreateAuthMiddleWare(validAuthToken string) StrictMiddlewareFunc {
	validTokenWithBearer := "Bearer " + validAuthToken
	return func(f StrictHandlerFunc, operationID string) StrictHandlerFunc {
		return func(ctx context.Context, w http.ResponseWriter, r *http.Request, input interface{}) (interface{}, error) {
			auth := r.Header.Get("Authorization")
			fmt.Println(auth, validTokenWithBearer)
			if auth != validTokenWithBearer {
				return nil, errors.New("Not authorized")
			}

			return f(ctx, w, r, input)
		}
	}
}

