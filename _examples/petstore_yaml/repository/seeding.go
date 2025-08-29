package repository

import (
	"github.com/google/uuid"
	"petstore_yaml/api/paths"
	"petstore_yaml/models"
	"time"
)

func (r *repository) seed() {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	categoryMap := map[string]models.Category{}
	seedCategories := []models.Category{
		{
			Id:   uuid.New(),
			Name: "Cats",
		},
		{
			Id:   uuid.New(),
			Name: "Dogs",
		},
		{
			Id:   uuid.New(),
			Name: "Hamsters",
		},
	}
	for _, c := range seedCategories {
		c.Uri = paths.CategoryURI(c.Id)
		categoryMap[c.Name] = c
		r.categories = append(r.categories, c)
	}
	seedPets := []models.Pet{
		{
			Id:   uuid.New(),
			Name: "Felix",
			DoB:  models.DoB(time.Now().Add(0 - (time.Hour * 24 * 40))),
			Category: models.Category{
				Name: "Cats",
			},
		},
		{
			Id:   uuid.New(),
			Name: "Rex",
			DoB:  models.DoB(time.Now().Add(0 - (time.Hour * 24 * 50))),
			Category: models.Category{
				Name: "Dogs",
			},
		},
		{
			Id:   uuid.New(),
			Name: "Nibbles",
			DoB:  models.DoB(time.Now().Add(0 - (time.Hour * 24 * 20))),
			Category: models.Category{
				Name: "Hamsters",
			},
		},
	}
	for _, p := range seedPets {
		p.Uri = paths.PetURI(p.Id)
		p.Category = categoryMap[p.Category.Name]
		r.pets = append(r.pets, &p)
	}
}
