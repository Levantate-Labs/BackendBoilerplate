package repository

import (
	"github.com/Levantate-labs/backend-boilerplate/models"
	"gorm.io/gorm"
)

type ModelRepository struct {
	db *gorm.DB
}

func NewModelRepository(db *gorm.DB) ModelRepository {
	return ModelRepository{db: db}
}

func (r ModelRepository) Create() (models.Model, error) {
	model := models.Model{}

	err := r.db.Create(&model).Error
	if err != nil {
		return model, err
	}
	return model, nil
}

func (r ModelRepository) Get() (models.Model, error) {
	var model models.Model

	err := r.db.First(&model).Error
	if err != nil {
		return model, err
	}
	return model, nil
}
