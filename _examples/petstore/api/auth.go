package api

import (
	"context"
	"github.com/go-andiamo/httperr"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strings"
)

const (
	ApiKeyHdr      = "X-Api-Key"
	AuthCookieName = "petstore-auth"
)

type Authentication struct {
	Key string
}

func (a *api) getRequestAuth(r *http.Request) (bool, *Authentication) {
	if a.apiKey == "" {
		return true, &Authentication{}
	}
	if key := r.Header.Get(ApiKeyHdr); strings.EqualFold(key, a.apiKey) {
		return true, &Authentication{key}
	}
	for _, cookie := range r.Cookies() {
		if cookie.Name == AuthCookieName {
			if strings.EqualFold(cookie.Value, a.apiKey) {
				return true, &Authentication{a.apiKey}
			} else {
				break
			}
		}
	}
	return false, nil
}

func applyAuthMiddlewares(thisApi any) chi.Middlewares {
	a := thisApi.(*api)
	return chi.Middlewares{
		func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				authed, auth := a.getRequestAuth(r)
				if !authed {
					a.writeErrorResponse(w, r, httperr.NewUnauthorizedError(""))
					return
				}
				r = authRequest(r, auth)
				next.ServeHTTP(w, r)
			})
		},
	}
}

func authRequest(r *http.Request, auth *Authentication) *http.Request {
	ctx := contextWithValue(r.Context(), auth)
	return r.Clone(ctx)
}

type ctxKey[T any] struct{}

// contextWithValue returns a copy of parent that contains the given value which can be
// retrieved by calling From contextWithValue the resulting context.
func contextWithValue[T any](ctx context.Context, v T) context.Context {
	return context.WithValue(ctx, ctxKey[T]{}, v)
}

// contextValueFrom returns the value associated contextWithValue the wanted type.
func contextValueFrom[T any](ctx context.Context) (T, bool) {
	v, ok := ctx.Value(ctxKey[T]{}).(T)
	return v, ok
}
