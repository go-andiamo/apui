package api

import (
	"github.com/go-andiamo/apui"
	"github.com/go-chi/chi/v5"
	"net/http"
	"petstore/api/paths"
	"petstore/repository"
)

type Api interface {
	Start() error
}

func New(r repository.Repository) Api {
	result := &api{
		repo: r,
	}
	result.setupBrowser()
	return result
}

type api struct {
	repo    repository.Repository
	browser *apui.Browser
}

func (a *api) Start() error {
	r := chi.NewRouter()
	if err := definition.SetupRoutes(r, a); err != nil {
		return err
	}
	return http.ListenAndServe(":8080", r)
}

func (a *api) GetRoot(w http.ResponseWriter, r *http.Request) {
	a.browser.Write(w, r, map[string]any{
		"root": map[string]any{
			"description": "Root discovery",
			"$uri":        paths.Root,
		},
		"pets": map[string]any{
			"description": "Pets",
			"$uri":        paths.Root + paths.Pets,
		},
		"categories": map[string]any{
			"description": "Categories",
			"$uri":        paths.Root + paths.Categories,
		},
	})
}

func (a *api) GetPets(w http.ResponseWriter, r *http.Request) {
	result, err := a.repo.SearchPets(r.Context(), "")
	_ = err
	a.browser.Write(w, r, result)
}
