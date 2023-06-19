package main

import (
	"log"

	"github.com/spf13/viper"

	"github.com/KonstantinP85/s-go-service"
	"github.com/KonstantinP85/s-go-service/pkg/handler"
	"github.com/KonstantinP85/s-go-service/pkg/service"
)

func main() {

	if err := initConfig(); err != nil {
		log.Fatalf("error init config: %s", err.Error())
	}

	services := service.NewService()
	handlers := handler.NewHandler(services)

	srv := new(s_go_project.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("server error %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}