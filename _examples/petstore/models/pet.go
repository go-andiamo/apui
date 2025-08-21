package models

import (
	"github.com/google/uuid"
	"time"
)

type Pet struct {
	Uri      string    `json:"$uri"`
	Id       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	DoB      time.Time `json:"dob"`
	Category Category  `json:"category"`
}
