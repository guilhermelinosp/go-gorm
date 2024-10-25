package main

import (
	"github.com/gin-gonic/gin"
	"github.com/guilhermelinosp/go-gorm/controllers"
	"github.com/guilhermelinosp/go-gorm/database"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	database.ConnectDatabase()

	r := gin.Default()

	r.POST("/products", controllers.CreateProduct)
	r.GET("/products", controllers.GetProducts)
	r.GET("/products/:id", controllers.GetProduct)
	r.PUT("/products/:id", controllers.UpdateProduct)
	r.DELETE("/products/:id", controllers.DeleteProduct)

	r.Run(":8081")
}
