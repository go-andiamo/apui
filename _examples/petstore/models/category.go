package models

import "github.com/google/uuid"

type Category struct {
	Uri  string    `json:"$uri"`
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
