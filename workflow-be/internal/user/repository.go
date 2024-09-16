package user

import (
	"workflow/internal/models"

	"gorm.io/gorm"
)

type Repository struct {
	Database *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return Repository{
		Database: db,
	}
}

func (repo Repository) FindOneByUsername(username string) (models.User, error){
	var result models.User

	db := repo.Database
	if err := db.Where("username = ?", username).First(&result).Error; err != nil {
		return result, err
	}
	return result, nil
}