package repository

import (
	"ecommerce-microservices/product-service/models"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) Migrate() error {
	// Auto-migrate the Product model to ensure the database schema matches
	return r.db.AutoMigrate(&models.Product{})
}

func (r *ProductRepository) CreateProduct(product models.Product) error {
	// Add a new product to the database
	return r.db.Create(&product).Error
}

func (r *ProductRepository) GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	// Retrieve all products from the database
	if err := r.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductRepository) GetProductByID(id string) (models.Product, error) {
	var product models.Product
	// Retrieve a product by ID
	if err := r.db.First(&product, "id = ?", id).Error; err != nil {
		return product, err
	}
	return product, nil
}

func (r *ProductRepository) UpdateProduct(id string, updated models.Product) error {
	// Update a product by ID
	return r.db.Model(&models.Product{}).Where("id = ?", id).Updates(updated).Error
}

func (r *ProductRepository) DeleteProduct(id string) error {
	// Delete a product by ID
	return r.db.Delete(&models.Product{}, "id = ?", id).Error
}
