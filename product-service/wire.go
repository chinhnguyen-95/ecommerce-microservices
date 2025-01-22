//go:build wireinject
// +build wireinject

package main

import (
	"ecommerce-microservices/common/config"
	"ecommerce-microservices/common/middleware"
	sc "ecommerce-microservices/order-service/config"
	"ecommerce-microservices/product-service/handlers"
	"ecommerce-microservices/product-service/repository"

	"github.com/google/wire"
	"github.com/gorilla/mux"
)

func ProvideRouter(handler *handlers.ProductHandler) *mux.Router {
	router := mux.NewRouter()
	router.Use(middleware.LoggingMiddleware)
	router.HandleFunc("/products", handler.AddProduct).Methods("POST")
	router.HandleFunc("/products", handler.GetProducts).Methods("GET")
	router.HandleFunc("/products/{id}", handler.GetProductByID).Methods("GET")
	router.HandleFunc("/products/{id}", handler.UpdateProduct).Methods("PUT")
	router.HandleFunc("/products/{id}", handler.DeleteProduct).Methods("DELETE")
	return router
}

func ProvideAppConfig() (*config.AppConfig, error) {
	return config.LoadConfig()
}

func InitializeRouter() (*mux.Router, error) {
	wire.Build(
		ProvideAppConfig,
		sc.NewDBConnection,
		repository.NewProductRepository,
		handlers.NewProductHandler,
		ProvideRouter,
	)
	return &mux.Router{}, nil
}
