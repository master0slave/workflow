package item

import (
	"fmt"
	"net/http"

	"strconv"
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
			"message": err.Error(),
		})
		return
	}

	// Response
	c.JSON(http.StatusCreated, gin.H{
		"data": item,
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

// Get Item By ID
func (controller Controller) GetItem(c *gin.Context) {
    // Get 'id' from the query string like /items?id=1
    idParam := c.Param("id")

    // Check if id was provided
    if idParam == "" {
        c.JSON(http.StatusBadRequest, gin.H{
            "message": "ID is required",
        })
        return
    }

    // Convert id to integer
    id, err := strconv.Atoi(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "message": "Invalid item ID",
        })
        return
    }

    // Continue with your service logic
	// convert type of id to uint

	// Get item	
    item, err := controller.Service.GetItem(uint(id))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "message": err.Error(),
        })
        return
    }

    // Response
    c.JSON(http.StatusOK, gin.H{
        "data": item,
    })
}

// Update Item by ID
func (controller Controller) UpdateItem(c *gin.Context) {
	// Get the ID from URL parameters
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid item ID",
		})
		return
	}
	//Bind
	// var request models.RequestItemWithID
	var request models.RequestItem

	if err := c.Bind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	fmt.Printf("%#v\n",request)

	// Update item
	item, err := controller.Service.UpdateItem(request, uint(id))
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
	// Get the ID from URL parameters
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid item ID",
		})
		return
	}
	//Bind
	var request models.RequestUpdateItemStatus
	if err := c.Bind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}
	// Update item
	err = controller.Service.UpdateItemStatus(uint(id), request.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Get item
	item, err := controller.Service.GetItem(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
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
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid item ID",
		})
		return
	}

	// Create a new item with the ID to delete
    item := models.Item{
        ID: uint(id), // Assign the ID to the item struct
    }

	// Delete item
	err = controller.Service.DeleteItem(item)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	// Response
	c.JSON(http.StatusOK, gin.H{	
		"message": "Item deleted successfully",
	})
}