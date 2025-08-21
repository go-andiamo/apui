package models

import (
	"github.com/google/uuid"
)

type Pet struct {
	Uri      string    `json:"$uri"`
	Id       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	DoB      DoB       `json:"dob"`
	Category Category  `json:"category"`
}
