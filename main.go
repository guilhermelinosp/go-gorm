package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

type Product struct {
	gorm.Model
	Name  string  `json:"name" example:"Laptop"`
	Price float64 `json:"price" example:"1299.99"`
}

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	DB, err = gorm.Open(mysql.Open(os.Getenv("MYSQL")), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// Migrate the Product model
	if err := DB.AutoMigrate(&Product{}); err != nil {
		log.Fatal("Failed to migrate the Product model:", err)
	}

	// Initialize Gin router
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		cors.Default().HandlerFunc(c.Writer, c.Request)
		c.Next()
	})

	// Define routes
	r.POST("/products", createProduct)
	r.GET("/products", getProducts)
	r.GET("/products/:id", getProduct)
	r.PUT("/products/:id", updateProduct)
	r.DELETE("/products/:id", deleteProduct)

	// Start the server
	r.Run(":8081")
}

func createProduct(c *gin.Context) {
	var product Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	DB.Create(&product)
	c.JSON(http.StatusOK, product)
}

func getProducts(c *gin.Context) {
	var products []Product
	DB.Find(&products)
	c.JSON(http.StatusOK, products)
}

func getProduct(c *gin.Context) {
	id := c.Param("id")
	var product Product

	if err := DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, product)
}

func updateProduct(c *gin.Context) {
	id := c.Param("id")
	var product Product

	if err := DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	DB.Save(&product)
	c.JSON(http.StatusOK, product)
}

func deleteProduct(c *gin.Context) {
	id := c.Param("id")
	if err := DB.Delete(&Product{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product successfully deleted"})
}
