package main

import (
	"log"

	tasks "github.com/assac453/lessons-zhashkevich"
	"github.com/assac453/lessons-zhashkevich/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(tasks.Server)
	log.Fatal(srv.Run("8000", handlers.InitRoutes()))
}
