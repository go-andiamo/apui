package models

import (
	"github.com/google/uuid"
)

type Pet struct {
	Uri      string    `json:"$uri"     oas:"format:uri,description:'URI of the pet'"`
	Id       uuid.UUID `json:"id"       oas:"format:uuid,description:'Unique identifier of the pet'"`
	Name     string    `json:"name"     oas:"description:'Name of the pet'"`
	DoB      DoB       `json:"dob"      oas:"type:string,format:date,description:'Date of birth'"`
	Category Category  `json:"category" oas:"description:'Category of the pet'"`
}
