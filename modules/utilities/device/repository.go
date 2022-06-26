package device

import "gorm.io/gorm"

type Repository interface {
	GetAllDevice() ([]Device, error)
	FindByID(ID int) (Device, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAllDevice() ([]Device, error) {
	var device []Device

	err := r.db.Find(&device).Error
	if err != nil {
		return device, err
	}

	return device, nil
}

func (r *repository) FindByID(ID int) (Device, error) {
	var device Device

	err := r.db.Where("no_device = ?", ID).Find(&device).Error
	if err != nil {
		return device, err
	}

	return device, nil
}
