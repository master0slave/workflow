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

// // Get All Items
func (service Service) GetItems() ([]models.Item, error) {
	return service.Repository.GetItems()
}

// Get a single item by ID
func (service Service) GetItem(id uint) (models.Item, error) {
    item, err := service.Repository.GetItem(id)
    return item, err
}
// Update Item By Id
func (service Service) UpdateItem(request models.RequestItem, id uint) (*models.Item, error) {
	// Fech item from database using repository
	item, err := service.Repository.GetItem(id)
	if err != nil {
		return nil, err
	}
	// Update item
	item.Title = request.Title
	item.Amount = request.Amount
	item.Quantity = request.Quantity

	// Save item to database
	if err := service.Repository.UpdateItem(&item);err != nil { 
		return nil, err 
	}
	return &item, nil
}

// Update Status
func (service Service) UpdateItemStatus(id uint, status constant.ItemStatus) error {
	return service.Repository.UpdateItemStatus(id, status)
}
// Delete
func (service Service) DeleteItem(item models.Item) error {
	return service.Repository.DeleteItem(&item)
}