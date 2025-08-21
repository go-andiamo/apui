package paths

import "fmt"

const (
	Root       = "/api"
	Pets       = "/pets"
	Categories = "/categories"

	UuidPattern = "^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12})$"
	UuidPath    = "/{id:" + UuidPattern + "}"

	pathPets       = Root + Pets
	pathCategories = Root + Categories
)

func PetURI(id any) string {
	return fmt.Sprintf(pathPets+"/%s", id)
}

func CategoryURI(id any) string {
	return fmt.Sprintf(pathCategories+"/%s", id)
}
