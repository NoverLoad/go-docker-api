package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/Noverload/gobank/server"
)

func seedAccount(store server.Storage, fname, lname, pw string) *server.Account {
	acc, err := server.NewAccount(fname, lname, pw)
	if err != nil {
		log.Fatal(err)
	}
	if err := store.CreateAccount(acc); err != nil {
		log.Fatal(err)
	}
	fmt.Println("new account => ", acc.Number)
	return acc
}

func seedAccounts(s server.Storage) {
	seedAccount(s, "lee", "DIR", "json1987")
}

func main() {
	seed := flag.Bool("seed", false, "seed in db")
	flag.Parse()

	store, err := server.NewPostpresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}
	if *seed {
		fmt.Println("seeding the database")
		seedAccounts(store)
	}

	server := server.NewAPIServer(":3000", store)
	server.Run()
}
