//go:build wireinject
// +build wireinject

package main

import (
	"ecommerce-microservices/common/config"
	"ecommerce-microservices/common/middleware"
	sc "ecommerce-microservices/order-service/config"
	"ecommerce-microservices/order-service/handlers"
	"ecommerce-microservices/order-service/repository"

	"github.com/google/wire"
	"github.com/gorilla/mux"
)

func ProvideRouter(handler *handlers.OrderHandler) *mux.Router {
	router := mux.NewRouter()
	router.Use(middleware.LoggingMiddleware)
	router.HandleFunc("/orders", handler.CreateOrder).Methods("POST")
	router.HandleFunc("/orders", handler.GetOrdersByUserID).Methods("GET").Queries("user_id", "{user_id}")
	router.HandleFunc("/order", handler.GetOrderByID).Methods("GET").Queries("id", "{id}")
	return router
}

func InitializeRouter() (*mux.Router, error) {
	wire.Build(
		config.LoadConfig,
		sc.NewDBConnection,
		repository.NewOrderRepository,
		handlers.NewOrderHandler,
		ProvideRouter,
	)
	return &mux.Router{}, nil
}
