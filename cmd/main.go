package main

import (
	"log"

	"github.com/qDes/todo-app"
	"github.com/qDes/todo-app/pkg/handler"
	"github.com/qDes/todo-app/pkg/repository"
	"github.com/qDes/todo-app/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)

	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while runnig http server: %s", err.Error())
	}

}
