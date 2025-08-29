package main

import (
	"log"
	"petstore_yaml/api"
	"petstore_yaml/repository"
)

func main() {
	r := repository.New()
	a := api.New(r)
	if err := a.Start(); err != nil {
		log.Fatal(err)
	}
}
