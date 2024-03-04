package main

import (
	"github.com/gin-gonic/gin"
	"product/api/handlers"
	"product/database"
	"product/models"
)

func init() {
	database.InitDB()
}

func main() {
	db := database.GetDB()
	db.AutoMigrate(&models.Product{})
	router := gin.Default()
	router.GET("/products", handlers.GetProducts)
	router.GET("/products/:id", handlers.GetProductByID)
	router.POST("/products", handlers.AddProduct)
	router.PATCH("/products/:id", handlers.UpdateProductName)
	router.DELETE("/products/:id", handlers.DeleteProduct)
	router.Run(":9090")
}