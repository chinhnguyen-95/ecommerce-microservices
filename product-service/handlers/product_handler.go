package handlers

import (
	"encoding/json"
	"net/http"

	"ecommerce-microservices/product-service/models"
	"ecommerce-microservices/product-service/repository"

	"github.com/gorilla/mux"
)

type ProductHandler struct {
	repo *repository.ProductRepository
}

func NewProductHandler(repo *repository.ProductRepository) *ProductHandler {
	return &ProductHandler{repo: repo}
}

func (h *ProductHandler) AddProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	// Decode the JSON body into the product struct
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Save the product using the repository
	if err := h.repo.CreateProduct(product); err != nil {
		http.Error(w, "Failed to save product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(product)
}

func (h *ProductHandler) GetProducts(w http.ResponseWriter, _ *http.Request) {
	// Retrieve all products using the repository
	products, err := h.repo.GetAllProducts()
	if err != nil {
		http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(products)
}

func (h *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	// Extract the product ID from the request path
	vars := mux.Vars(r)
	id := vars["id"]

	// Retrieve the product by ID using the repository
	product, err := h.repo.GetProductByID(id)
	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(product)
}

func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	// Extract the product ID from the request path
	vars := mux.Vars(r)
	id := vars["id"]

	var product models.Product
	// Decode the JSON body into the product struct
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Update the product using the repository
	if err := h.repo.UpdateProduct(id, product); err != nil {
		http.Error(w, "Failed to update product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(product)
}

func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	// Extract the product ID from the request path
	vars := mux.Vars(r)
	id := vars["id"]

	// Delete the product using the repository
	if err := h.repo.DeleteProduct(id); err != nil {
		http.Error(w, "Failed to delete product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
