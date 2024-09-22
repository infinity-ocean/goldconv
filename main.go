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
	}
	repo := repo.NewRepo(pool)
	svc := service.NewService(repo)
	ctrl := controller.NewController(svc)
	
	http.HandleFunc("GET /show-balance/", ctrl.ShowBalance)

	http.ListenAndServe(":8080", nil)

	fmt.Println(http.ListenAndServe(":8080", nil))
	// handlers registration
}
