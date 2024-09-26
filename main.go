package main

import (
	"fmt"

	"github.com/infinity-ocean/goldconv/internal/controller"
	"github.com/infinity-ocean/goldconv/internal/repo"
	"github.com/infinity-ocean/goldconv/internal/service"
)


func main() {
	pool, err := repo.MakePool()
	if err != nil {
		fmt.Println("Pool creation error") 
		return
	}
	repo := repo.NewRepo(pool)
	svc := service.NewService(repo)
	ctrl := controller.NewController(svc, ":9090")

	ctrl.Run()
}
