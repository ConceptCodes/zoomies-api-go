package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fatih/color"
	"github.com/gorilla/mux"

	"zoomies-api-go/pkg/config"
	"zoomies-api-go/pkg/handlers"
	"zoomies-api-go/pkg/middlewares"
	"zoomies-api-go/pkg/repository"
	"zoomies-api-go/pkg/storage/postgres"
)

func main() {

	config.LoadAppConfig()

	db, err := postgres.GetDBInstance()
	if err != nil {
		log.Fatal(err)
	}

	userRepo := repository.NewGormUserRepository(db)
	authHandler := handlers.NewAuthHandler(userRepo)

	router := mux.NewRouter()

	router.Use(middlewares.LogRequest)
	router.Use(middlewares.LogResponse)

	router.HandleFunc("/auth/login", authHandler.LoginHandler).Methods("POST")
	router.HandleFunc("/auth/register", authHandler.RegisterHandler).Methods("POST")

	router.HandleFunc("/health/alive", handlers.NewHealthHandler().ServiceAliveHandler).Methods("GET")

	color.Green("Zoomies Api started on port %s", config.AppConfig.Port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", config.AppConfig.Port), nil))

}
