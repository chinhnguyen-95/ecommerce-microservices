package repository

import (
	"ecommerce-microservices/order-service/models"

	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) Migrate() error {
	// Auto-migrate the Order model
	return r.db.AutoMigrate(&models.Order{})
}

func (r *OrderRepository) CreateOrder(order models.Order) error {
	// Add a new order to the database
	return r.db.Create(&order).Error
}

func (r *OrderRepository) GetOrderByID(id string) (models.Order, error) {
	var order models.Order
	// Retrieve an order by ID
	err := r.db.First(&order, "id = ?", id).Error
	return order, err
}

func (r *OrderRepository) GetOrdersByUserID(userID string) ([]models.Order, error) {
	var orders []models.Order
	// Retrieve orders by user ID
	err := r.db.Where("user_id = ?", userID).Find(&orders).Error
	return orders, err
}
