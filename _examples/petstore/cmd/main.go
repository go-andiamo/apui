package main

import (
	"log"
	"petstore/api"
	"petstore/repository"
)

func main() {
	r := repository.New()
	a := api.New(r)
	if err := a.Start(); err != nil {
		log.Fatal(err)
	}
}
