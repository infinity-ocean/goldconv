package main

import (
	"fmt"
	"net/http"
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
	ctrl := controller.NewController(svc)

	http.HandleFunc("/goldconv/show-balance", ctrl.ShowBalance)
	err = http.ListenAndServe("localhost:9090", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
