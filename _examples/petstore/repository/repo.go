package repository

import (
	"context"
	"github.com/go-andiamo/httperr"
	"petstore/models"
	"sync"
)

type Repository interface {
	SearchPets(ctx context.Context, category string) ([]*models.Pet, error)
	GetPet(ctx context.Context, id string) (*models.Pet, error)
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

func (r *repository) SearchPets(ctx context.Context, category string) ([]*models.Pet, error) {
	//TODO implement me
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	return r.pets, nil
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
