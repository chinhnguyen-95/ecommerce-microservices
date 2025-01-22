//go:build wireinject
// +build wireinject

package main

import (
	"ecommerce-microservices/common/config"
	"ecommerce-microservices/common/middleware"
	sc "ecommerce-microservices/order-service/config"
	"ecommerce-microservices/user-service/handlers"
	"ecommerce-microservices/user-service/repository"

	"github.com/google/wire"
	"github.com/gorilla/mux"
)

func ProvideRouter(handler *handlers.UserHandler) *mux.Router {
	router := mux.NewRouter()
	router.Use(middleware.LoggingMiddleware)
	router.HandleFunc("/register", handler.RegisterUser).Methods("POST")
	router.HandleFunc("/user", handler.GetUserByID).Methods("GET")
	return router
}

func InitializeRouter() (*mux.Router, error) {
	wire.Build(
		config.LoadConfig,
		sc.NewDBConnection,
		repository.NewUserRepository,
		handlers.NewUserHandler,
		ProvideRouter,
	)
	return &mux.Router{}, nil
}
