package main

import (
	"log"

	"github.com/qDes/todo-app"
	"github.com/qDes/todo-app/pkg/handler"
)

func main() {

	handlers := new(handler.Handler)
	srv := new(todo.Server)

	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while runnig http server: %s", err.Error())
	}

}
