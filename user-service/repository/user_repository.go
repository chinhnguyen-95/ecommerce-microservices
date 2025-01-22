package repository

import (
	"gorm.io/gorm"

	"ecommerce-microservices/user-service/models"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Migrate() error {
	return r.db.AutoMigrate(&models.User{})
}

func (r *UserRepository) CreateUser(user models.User) error {
	return r.db.Create(&user).Error
}

func (r *UserRepository) GetUserByID(id string) (models.User, error) {
	var user models.User
	err := r.db.First(&user, "id = ?", id).Error
	return user, err
}

func (r *UserRepository) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	err := r.db.First(&user, "email = ?", email).Error
	return user, err
}
