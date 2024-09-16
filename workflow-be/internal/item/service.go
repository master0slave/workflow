package item

import (
	"workflow/internal/constant"
	"workflow/internal/models"

	"gorm.io/gorm"
)

type Service struct {
	Repository Repository

}

func NewService(db *gorm.DB) Service {
	return Service{
		Repository: NewRepository(db),
	}
}
// Create Item
func (service Service) CreateItem(req models.RequestItem) (models.Item, error){
	item := models.Item{
		Title:    req.Title,
		Amount:   req.Amount,
		Quantity: req.Quantity,
		Status:   constant.ItemPenddingStatus,
	}

	if err := service.Repository.CreateItem(&item);err != nil { 
			return models.Item{}, err 
	} 
	
	return item, nil 
}

// Get All Items
func (service Service) GetItems() ([]models.Item, error) {
	return service.Repository.GetItems()
}

// Get Item By ID
func (service Service) GetItem(id int) (models.Item, error) {
	return service.Repository.GetItem(id)
}

// Update Item
func (service Service) UpdateItem(item models.Item) error {
	return service.Repository.UpdateItem(&item)
}
// Update Status
func (service Service) UpdateItemStatus(id int, status constant.ItemStatus) error {
	return service.Repository.UpdateItemStatus(id, status)
}
// Delete
func (service Service) DeleteItem(item models.Item) error {
	return service.Repository.DeleteItem(&item)
}