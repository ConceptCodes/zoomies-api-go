package main

import (
	"log"
	"net/http"
	"strconv"

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
	healthHandler := handlers.NewHealthHandler()

	router := mux.NewRouter()

	router.Use(middlewares.TraceRequest)
	router.Use(middlewares.LogRequest)
	router.Use(middlewares.LogResponse)

	router.HandleFunc("/api/auth/login", authHandler.LoginHandler).Methods("POST")
	router.HandleFunc("/api/auth/register", authHandler.RegisterHandler).Methods("POST")

	router.HandleFunc("/api/health/alive", healthHandler.ServiceAliveHandler).Methods("GET")

	port := strconv.Itoa(config.AppConfig.Port)

	color.Green("Zoomies Api started on port %s", port)

	log.Fatal(http.ListenAndServe(":"+port, router))

}
