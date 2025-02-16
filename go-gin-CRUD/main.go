package main

import (
	"github.com/vansh2308/go-gin-CRUD.git/handlers"
	"github.com/vansh2308/go-gin-CRUD.git/services"
)

func main() {
	services := services.NewServices()
	router := handlers.NewRouter(services)

	router.Run(":5000")
}
