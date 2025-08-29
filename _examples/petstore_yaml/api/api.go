package api

import (
	"encoding/json"
	"fmt"
	"github.com/go-andiamo/apui"
	"github.com/go-andiamo/httperr"
	"github.com/go-chi/chi/v5"
	"net/http"
	"petstore_yaml/api/paths"
	"petstore_yaml/models"
	"petstore_yaml/models/params"
	"petstore_yaml/models/requests"
	"petstore_yaml/repository"
	"strings"
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
	// set apiKey to any key to enable authorization
	apiKey string
}

func (a *api) Start() error {
	httperr.DefaultErrorWriterShowCause = true
	httperr.DefaultErrorWriterShowStack = true
	httperr.DefaultPackageName = "petstore_yaml/"
	r := chi.NewRouter()
	a.setupRoutes(r)
	fmt.Printf("Starting server on http://localhost:8080\n")
	return http.ListenAndServe(":8080", r)
}

func (a *api) setupRoutes(r *chi.Mux) {
	r.MethodFunc(http.MethodGet, paths.Root, a.GetRoot)
	r.MethodFunc(http.MethodGet, paths.Root+paths.Pets, a.GetPets)
	r.MethodFunc(http.MethodPost, paths.Root+paths.Pets, a.PostPets)
	r.MethodFunc(http.MethodGet, paths.Root+paths.Pets+paths.UuidPath, a.GetPet)
	r.MethodFunc(http.MethodPut, paths.Root+paths.Pets+paths.UuidPath, a.PutPet)
	r.MethodFunc(http.MethodDelete, paths.Root+paths.Pets+paths.UuidPath, a.DeletePet)
	r.MethodFunc(http.MethodGet, paths.Root+paths.Categories, a.GetCategories)
	r.MethodFunc(http.MethodGet, paths.Root+paths.Categories+paths.UuidPath, a.GetCategory)
}

func (a *api) GetRoot(w http.ResponseWriter, r *http.Request) {
	a.writeResponse(w, r, map[string]any{
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
	}, http.StatusOK)
}

func (a *api) GetPets(w http.ResponseWriter, r *http.Request) {
	var filter *params.PetFilter
	var err error
	if filter, err = params.PetFilterFromRequest(r); err == nil {
		var result []*models.Pet
		if result, err = a.repo.SearchPets(r.Context(), filter); err == nil {
			a.writeResponse(w, r, result, http.StatusOK)
		}
	}
	if err != nil {
		a.writeErrorResponse(w, r, err)
	}
}

func (a *api) PostPets(w http.ResponseWriter, r *http.Request) {
	var request *requests.AddPet
	var err error
	if request, err = requests.AddPetFromRequest(r); err == nil {
		var result *models.Pet
		if result, err = a.repo.AddPet(r.Context(), *request); err == nil {
			a.writeResponse(w, r, result, http.StatusCreated)
		}
	}
	if err != nil {
		a.writeErrorResponse(w, r, err)
	}
}

func (a *api) GetPet(w http.ResponseWriter, r *http.Request) {
	if result, err := a.repo.GetPet(r.Context(), chi.URLParam(r, "id")); err == nil {
		a.writeResponse(w, r, result, http.StatusOK)
	} else {
		a.writeErrorResponse(w, r, err)
	}
}

func (a *api) PutPet(w http.ResponseWriter, r *http.Request) {
	var request *requests.UpdatePet
	var err error
	if request, err = requests.UpdatePetFromRequest(r); err == nil {
		var result *models.Pet
		if result, err = a.repo.UpdatePet(r.Context(), chi.URLParam(r, "id"), *request); err == nil {
			a.writeResponse(w, r, result, http.StatusOK)
		}
	}
	if err != nil {
		a.writeErrorResponse(w, r, err)
	}
}

func (a *api) DeletePet(w http.ResponseWriter, r *http.Request) {
	if err := a.repo.DeletePet(r.Context(), chi.URLParam(r, "id")); err == nil {
		a.writeResponse(w, r, nil, http.StatusNoContent)
	} else {
		a.writeErrorResponse(w, r, err)
	}
}

func (a *api) GetCategories(w http.ResponseWriter, r *http.Request) {
	if result, err := a.repo.ListCategories(r.Context()); err == nil {
		a.writeResponse(w, r, result, http.StatusOK)
	} else {
		a.writeErrorResponse(w, r, err)
	}
}

func (a *api) GetCategory(w http.ResponseWriter, r *http.Request) {
	if result, err := a.repo.GetCategory(r.Context(), chi.URLParam(r, "id")); err == nil {
		a.writeResponse(w, r, result, http.StatusOK)
	} else {
		a.writeErrorResponse(w, r, err)
	}
}

func (a *api) writeResponse(w http.ResponseWriter, r *http.Request, result any, statusCode int) {
	if strings.Contains(r.Header.Get("Accept"), "text/html") {
		a.browser.Write(w, r, result)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if statusCode > 0 {
		w.WriteHeader(statusCode)
	}
	if result != nil {
		_ = json.NewEncoder(w).Encode(result)
	}
}

func (a *api) writeErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	if strings.Contains(r.Header.Get("Accept"), "text/html") {
		a.browser.Write(w, r, err)
		return
	}
	httperr.DefaultErrorWriter.WriteError(err, w)
}
