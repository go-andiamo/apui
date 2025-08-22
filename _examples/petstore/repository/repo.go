package repository

import (
	"context"
	"github.com/go-andiamo/httperr"
	"github.com/google/uuid"
	"petstore/api/paths"
	"petstore/models"
	"petstore/models/params"
	"petstore/models/requests"
	"slices"
	"strings"
	"sync"
	"time"
)

/*
Nothing to see here!

This repository implementation is just to provide data to the example api

Not intended as recommended code!
*/

type Repository interface {
	SearchPets(ctx context.Context, filter *params.PetFilter) ([]*models.Pet, error)
	GetPet(ctx context.Context, id string) (*models.Pet, error)
	AddPet(ctx context.Context, pet requests.AddPet) (*models.Pet, error)
	DeletePet(ctx context.Context, id string) error
	ListCategories(ctx context.Context) ([]models.Category, error)
	GetCategory(ctx context.Context, id string) (models.Category, error)
}

func New() Repository {
	result := &repository{}
	result.seed()
	return result
}

type repository struct {
	mutex      sync.RWMutex
	pets       []*models.Pet
	categories []models.Category
}

func (r *repository) SearchPets(ctx context.Context, filter *params.PetFilter) ([]*models.Pet, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	if filter == nil {
		return r.pets, nil
	}
	result := make([]*models.Pet, 0, len(r.pets))
	for _, pet := range r.pets {
		if filter.Matches(pet) {
			result = append(result, pet)
		}
	}
	return sortPets(result, filter.Order), nil
}

func sortPets(pets []*models.Pet, orders []string) []*models.Pet {
	if len(orders) == 0 {
		return pets
	}
	slices.SortFunc(pets, func(a, b *models.Pet) int {
		cmp := 0
		reverse := false
		for _, order := range orders {
			if strings.HasPrefix(order, "-") {
				reverse = true
				order = strings.TrimPrefix(order, "-")
			}
			switch order {
			case "name":
				cmp = strings.Compare(a.Name, b.Name)
			case "category":
				cmp = strings.Compare(a.Category.Name, b.Category.Name)
			case "dob":
				cmp = time.Time(a.DoB).Compare(time.Time(b.DoB))
			}
			if cmp != 0 {
				break
			}
		}
		if reverse {
			cmp = 0 - cmp
		}
		return cmp
	})
	return pets
}

func (r *repository) GetPet(ctx context.Context, id string) (*models.Pet, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	for _, pet := range r.pets {
		if pet.Id.String() == id {
			return pet, nil
		}
	}
	return nil, httperr.NewNotFoundError("pet not found")
}

func (r *repository) AddPet(ctx context.Context, pet requests.AddPet) (*models.Pet, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	var category *models.Category
	for _, c := range r.categories {
		if pet.Category.Id == c.Id.String() || pet.Category.Name == c.Name {
			category = &c
			break
		}
	}
	if category == nil {
		return nil, httperr.NewUnprocessableEntityErrorf("category not found")
	}
	id := uuid.New()
	result := &models.Pet{
		Uri:      paths.PetURI(id),
		Id:       id,
		Name:     pet.Name,
		DoB:      pet.DoB,
		Category: *category,
	}
	r.pets = append(r.pets, result)
	return result, nil
}

func (r *repository) DeletePet(ctx context.Context, id string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	for i, pet := range r.pets {
		if pet.Id.String() == id {
			r.pets[i] = r.pets[len(r.pets)-1]
			r.pets = r.pets[:len(r.pets)-1]
			return nil
		}
	}
	return httperr.NewNotFoundError("pet not found")
}

func (r *repository) ListCategories(ctx context.Context) ([]models.Category, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	return append([]models.Category{}, r.categories...), nil
}

func (r *repository) GetCategory(ctx context.Context, id string) (models.Category, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	for _, c := range r.categories {
		if c.Id.String() == id {
			return c, nil
		}
	}
	return models.Category{}, httperr.NewNotFoundError("category not found")
}
