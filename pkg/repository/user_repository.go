package repository

import (
	"gorm.io/gorm"

	"zoomies-api-go/pkg/models"
)

type UserRepository interface {
	FindByEmail(email string) (*models.User, error)
	Save(user *models.User) error
}

type GormUserRepository struct {
	db *gorm.DB
}

func (r *GormUserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *GormUserRepository) Save(user *models.User) error {
	return r.db.Save(user).Error
}

func NewGormUserRepository(db *gorm.DB) UserRepository {
	return &GormUserRepository{db}
}
