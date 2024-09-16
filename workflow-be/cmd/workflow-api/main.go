package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"workflow/internal/auth"
	"workflow/internal/item"

	"github.com/fvbock/endless"
	"github.com/gin-contrib/cors"
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

// Logger middleware
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		end := time.Now()
		latency := end.Sub(start)
		log.Printf("Method: %s \nURI: %s \nStatus: %d \nLatency: %s \n\n", c.Request.Method, c.Request.RequestURI, c.Writer.Status(), latency)
	}
}
// main is the entry point of the application. It will connect to the database,
// create a router, and setup the routes. Then it will start the server and
// listen to the port specified by the PORT environment variable.
func main() {

	// Connect to Database
	db_url := os.Getenv("DB_URL")
	if db_url == "" {
		db_url = fmt.Sprintf("postgres://%s:%s@%s:%s/%s", 
							 os.Getenv("DB_USER"), 
							 os.Getenv("DB_PASSWORD"), 
							 os.Getenv("DB_HOST"), 
							 os.Getenv("DB_PORT"), 
							 os.Getenv("DB_NAME"),
							)
	}
	db ,err := gorm.Open(
		postgres.Open(db_url),
	)
	if err != nil {
		log.Panic(err)
	}

	// Controller
	controller := item.NewController(db)
	
	
	r := gin.Default()

	config := cors.DefaultConfig()

	// frontend URL
	config.AllowOrigins = []string{"*"}
	config.AllowCredentials = true
	r.Use(cors.New(config))

	r.Use(Logger())

// TODO: สร้าง Group Rounter
	// Router Registration Group
	items := r.Group("/items")

	items.Use(auth.BasicAuth([]auth.Credential{
		{"admin", "secret"},
	}))
	{
		items.GET("", controller.GetItems) // เพิ่มข้อมูลเบิกงบใหม่ได้
		items.GET("/:id", controller.GetItem) // ดูข้อมูลเบิกงบทั้งหมด
		items.POST("", controller.CreateItem) // ดูข้อมูลเบิกงบที่ต้องการ
		items.PUT("/:id", controller.UpdateItem) // แก้ไขข้อมูล
		items.PATCH("/:id", controller.UpdateItemStatus) // ปรับเปลี่ยนแก้ไขข้อมูลสถานะการเบิกงบ (เป็นสถานะ APPROVED หรือ REJECTED หรือ PENDING) 
		items.DELETE("/:id", controller.DeleteItem) // ลบ
	}
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Start and run the server with greceful shutdown
	if err := endless.ListenAndServe(":" + os.Getenv("PORT"), r); err != nil {
		log.Panic(err)
	}	

}