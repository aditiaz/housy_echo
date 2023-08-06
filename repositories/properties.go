package repositories

import (
	"housy/models"

	"gorm.io/gorm"
)

type PropertyRepository interface {
	FindProperties() ([]models.Property, error)
	GetProperty(ID int) (models.Property, error)
	AddProperty(product models.Property) (models.Property, error)
}

func RepositoryProperty(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindProperties() ([]models.Property, error) {
	var Properties []models.Property
	err := r.db.Find(&Properties).Error

	return Properties, err
}

func (r *repository) GetProperty(ID int) (models.Property, error) {
	var Property models.Property
	err := r.db.First(&Property, ID).Error

	return Property, err
}

func (r *repository) AddProperty(Property models.Property) (models.Property, error) {
	err := r.db.Preload("UsersProfileResponse").Create(&Property).Error

	return Property, err
}
