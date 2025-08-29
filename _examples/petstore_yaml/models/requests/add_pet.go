package requests

import (
	"encoding/json"
	"github.com/go-andiamo/httperr"
	"net/http"
	"petstore_yaml/models"
)

type AddPet struct {
	Name     string     `json:"name"     oas:"required,type:string,description:'Name of the pet'"`
	DoB      models.DoB `json:"dob"      oas:"required,type:string,format:date,description:'Date of birth'"`
	Category Category   `json:"category" oas:"required,type:object,description:'Category of the pet'"`
}

type Category struct {
	Id   string `json:"id"   oas:"format:uuid,description:'ID of the category'"`
	Name string `json:"name" oas:"type:string,description:'Name of the category'"`
}

func AddPetFromRequest(r *http.Request) (*AddPet, error) {
	if r.Body == nil {
		return nil, httperr.NewBadRequestError("body is required")
	}
	result := new(AddPet)
	if err := json.NewDecoder(r.Body).Decode(result); err != nil {
		if hErr, ok := err.(httperr.HttpError); ok {
			return nil, hErr
		}
		return nil, httperr.NewBadRequestError("invalid json body").WithCause(err)
	}
	return result, nil
}
