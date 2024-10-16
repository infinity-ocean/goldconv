package main

import (
	"fmt"
	"github.com/infinity-ocean/goldconv/internal/config"
	"github.com/infinity-ocean/goldconv/internal/controller"
	"github.com/infinity-ocean/goldconv/internal/repo"
	"github.com/infinity-ocean/goldconv/internal/service"
)


func main() {
	var config config.Config
	config.Parse()

	pool, err := repo.MakePool(config)
	if err != nil {
		fmt.Println("Pool creation error") 
		return
	}
	//TODO вынести run в main
	//TODO сделать доступ к env исключительно через конфиг
	//TODO перенести логику jwt в service

	repo := repo.NewRepo(pool)
	svc := service.NewService(repo)
	ctrl := controller.NewController(svc, ":8080")

	ctrl.Run(repo, config)
}
