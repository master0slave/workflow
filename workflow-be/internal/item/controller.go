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
	// Log the request body for debugging purposes
	fmt.Printf("%#v\n",request)

	// Call the service to create the item
	item, err := controller.Service.CreateItem(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Return the created item in the response with a status code of 201
	c.JSON(http.StatusCreated, gin.H{
		"id":       item.ID,        // Include the item ID
        "title":    item.Title,     // Include the title
        "amount":   item.Amount,    // Include the amount
        "quantity": item.Quantity,  // Include the quantity
        "status":   item.Status,    // Include the status
		"message": "Successfully created item",
	})
}

// Get All Items
func (controller Controller) GetItems(c *gin.Context) {
	// Get items
	items, err := controller.Service.GetItems()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Check if the items list is empty
    if len(items) == 0 {
        c.JSON(http.StatusOK, gin.H{
            "message": "No items found",
        })
        return
    }

	/// Return the items list directly
	c.JSON(http.StatusOK,items)
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

	// Check if the item was found
	if item.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Item not found",
		})
		return
	}

    // Response
    c.JSON(http.StatusOK, item)
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