package main

import (
	"devroomtodo/routes"
	"devroomtodo/storage"
)

func main() {
	server := routes.SetupRouter()
	storage.SetupStorage()

	err := server.Run(":3332")
	if err != nil {
		panic(err)
	}
}
