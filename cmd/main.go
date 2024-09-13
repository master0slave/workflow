package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"workflow/internal/item"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(`Error loading .env file`)
	}
}
func main() {

	// Fetch environment variables
	db_user := os.Getenv("DB_USER")
	db_password := os.Getenv("DB_PASSWORD")
	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")
	db_name := os.Getenv("DB_NAME")

	// Connect to Database
	db_url := os.Getenv("DB_URL")
	if db_url == "" {
		db_url = fmt.Sprintf("postgres://%s:%s@%s:%s/%s", db_user, db_password, db_host, db_port, db_name)
	}
	db ,err := gorm.Open(
		postgres.Open(db_url),
	)
	if err != nil {
		log.Panic(err)
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "2024"
	}
	fmt.Printf("Running on port %s\n", port)

	// Controller
	controller := item.NewController(db)
	
	
	r := gin.Default()


	// Router Registration
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	// เพิ่มข้อมูลเบิกงบใหม่ได้
	r.POST("/items", controller.CreateItem)

	// ดูข้อมูลเบิกงบทั้งหมด
	r.GET("/items", controller.GetItems)

	// ดูข้อมูลเบิกงบที่ต้องการ
	r.GET("/items/:id", controller.GetItem)

	// แก้ไขข้อมูลเบิกงบ
	r.PUT("/items/:id", controller.UpdateItem)

	// ปรับเปลี่ยนแก้ไขข้อมูลสถานะการเบิกงบ (เป็นสถานะ APPROVED หรือ REJECTED หรือ PENDING) 
	r.PATCH("/items/:id", controller.UpdateItemStatus)

	// ลบข้อมูลเบิกงบ
	r.DELETE("/items/:id", controller.DeleteItem)


	// Start and run the server
	if err := r.Run(":" + port); err != nil {
		log.Panic(err)
	}

}