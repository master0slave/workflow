package item

import (
	"fmt"
	"net/http"
	"workflow/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Controller struct {
	Service Service
}

func NewController(db *gorm.DB) Controller {	
	return Controller{
		Service: NewService(db),
	}
}


// Create Item
func (controller Controller) CreateItem(c *gin.Context) {
	//Bind
	var request models.RequestItem
	
	if err := c.Bind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}
	fmt.Printf("%#v\n",request)

	// Create item
	item, err := controller.Service.CreateItem(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	// Response
	c.JSON(http.StatusCreated, gin.H{
		"data": item,
	})
	c.JSON(http.StatusCreated, gin.H{
		"message": "Successfully created item",
	})
}

// Get All Items
func (controller Controller) GetItems(c *gin.Context) {
	// Get items
	items, err := controller.Service.GetItems()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	// Response
	c.JSON(http.StatusOK, gin.H{
		"data": items,
	})
}

//Get Item By ID
func (controller Controller) GetItem(c *gin.Context) {
	//Bind
	var request models.RequestItemWithID
	if err := c.Bind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}
	// Get item
	item, err := controller.Service.GetItem(request.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	// Response
	c.JSON(http.StatusOK, gin.H{
		"data": item,
	})
}	

// Update Item
func (controller Controller) UpdateItem(c *gin.Context) {
	//Bind
	var request models.RequestItemWithID
	if err := c.Bind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}
	// Update item
	item, err := controller.Service.UpdateItem(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	// Response
	c.JSON(http.StatusOK, gin.H{	
		"data": item,
	})
}

// Update Status
func (controller Controller) UpdateItemStatus(c *gin.Context) {
	//Bind
	var request models.RequestItem
	if err := c.Bind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}
	// Update item
	item, err := controller.Service.UpdateItemStatus(request.ID, request.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	// Response
	c.JSON(http.StatusOK, gin.H{	
		"data": item,
	})
}	

// Delete
func (controller Controller) DeleteItem(c *gin.Context) {
	//Bind
	var request models.RequestItem
	if err := c.Bind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}
	// Delete item
	item, err := controller.Service.DeleteItem(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	// Response
	c.JSON(http.StatusOK, gin.H{	
		"data": item,
	})
}