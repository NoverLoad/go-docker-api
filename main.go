package main

import "github.com/Noverload/gobank/server"

func main() {

	server := server.NewAPIServer(":9000")
	server.Run()
}
