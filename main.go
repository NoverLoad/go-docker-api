package main

import (
	"log"

	"github.com/Noverload/gobank/server"
)

func main() {
	store, err := server.NewPostpresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	server := server.NewAPIServer(":3000", store)
	server.Run()
}
