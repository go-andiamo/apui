package paths

import "fmt"

const (
	Root       = "/api"
	Pets       = "/pets"
	Categories = "/categories"

	pathPets       = Root + Pets
	pathCategories = Root + Categories
)

func PetURI(id any) string {
	return fmt.Sprintf(pathPets+"/%s", id)
}

func CategoryURI(id any) string {
	return fmt.Sprintf(pathCategories+"/%s", id)
}
