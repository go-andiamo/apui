package requests

import (
	"encoding/json"
	"github.com/go-andiamo/httperr"
	"net/http"
	"petstore/models"
)

type AddPet struct {
	Name     string     `json:"name" oas:"required,type:string"`
	DoB      models.DoB `json:"dob" oas:"required,type:string"`
	Category Category   `json:"category"oas:"required,type:object"`
}

func AddPetFromRequest(r *http.Request) (*AddPet, error) {
	if r.Body == nil {
		return nil, httperr.NewBadRequestError("body is required")
	}
	result := new(AddPet)
	if err := json.NewDecoder(r.Body).Decode(result); err != nil {
		return nil, httperr.NewBadRequestError("invalid json body").WithCause(err)
	}
	return result, nil
}

type Category struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
