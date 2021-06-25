package main

import (
	"log"

	"github.com/spf13/viper"

	"github.com/qDes/todo-app"
	"github.com/qDes/todo-app/pkg/handler"
	"github.com/qDes/todo-app/pkg/repository"
	"github.com/qDes/todo-app/pkg/service"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)

	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while runnig http server: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
