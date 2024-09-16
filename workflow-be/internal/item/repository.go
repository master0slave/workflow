package item

import (
	"workflow/internal/constant"
	"workflow/internal/models"

	"gorm.io/gorm"
)

type Repository struct { Database *gorm.DB }
func NewRepository(db *gorm.DB) Repository { 
	return Repository{ 
		Database: db, 
		} 
} 
// Create
func (repo Repository) CreateItem(item *models.Item) error { 
	return repo.Database.Create(item).Error 
}

// Get All Items
func (repo Repository) GetItems() ([]models.Item, error) {
	var items []models.Item
	return items, repo.Database.Find(&items).Error

}

// Get Item By ID
func (repo Repository) GetItem(id int) (models.Item, error) {
	var item models.Item
	return item, repo.Database.First(&item, id).Error
}

// Update Item
func (repo Repository) UpdateItem(item *models.Item) error {
	return repo.Database.Save(item).Error
}

// Update Status
func (repo Repository) UpdateItemStatus(id int, status constant.ItemStatus) error {
	return repo.Database.Model(&models.Item{}).Where("id = ?", id).Update("status", status).Error
}

// Delete
func (repo Repository) DeleteItem(item *models.Item) error {
	return repo.Database.Delete(item).Error
}