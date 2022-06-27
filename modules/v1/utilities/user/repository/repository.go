package repository

import (
	"ClearningPatternGO/modules/v1/utilities/user/models"

	"gorm.io/gorm"
)

type Repository interface {
	Save(user models.User) (models.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
