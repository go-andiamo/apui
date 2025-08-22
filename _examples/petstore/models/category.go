package models

import "github.com/google/uuid"

type Category struct {
	Uri  string    `json:"$uri" oas:"format:uri,description:'URI of the category'"`
	Id   uuid.UUID `json:"id"   oas:"format:uuid,description:'Unique identifier of the category'"`
	Name string    `json:"name" oas:"description:'Name of the category'"`
}
