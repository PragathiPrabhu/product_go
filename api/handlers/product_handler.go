package handlers

import (
	"net/http"
	"product/database"
	"product/models"
	"github.com/gin-gonic/gin"
)

func GetProducts(context *gin.Context) {
	db := database.GetDB()
	if db == nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to database"})
		return
	}

	var products []models.Product
	result := db.Find(&products)
	if result.Error != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}

	context.IndentedJSON(http.StatusOK, products)
}

func GetProductByID(context *gin.Context) {
	db := database.GetDB()
	id := context.Param("id")

	// Check if the product exists
	var product models.Product
	result := db.First(&product, id)
	if result.Error != nil {
		context.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	// If product found, return it as JSON
	context.IndentedJSON(http.StatusOK, product)
}

func AddProduct(context *gin.Context) {
    db := database.GetDB()

    // Bind the JSON request body to a struct
    var newProduct models.Product
    if err := context.ShouldBindJSON(&newProduct); err != nil {
        context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Create the product in the database
    if err := db.Create(&newProduct).Error; err != nil {
        context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
        return
    }

    context.IndentedJSON(http.StatusCreated, newProduct)
}

func UpdateProductName(context *gin.Context) {
	db := database.GetDB()

	// Get the ID of the product to update from the URL parameter
	id := context.Param("id")

	// Check if the product exists
	var existingProduct models.Product
	result := db.First(&existingProduct, id)
	if result.Error != nil {
		context.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	// Bind the JSON request body to a struct
	var updateData struct {
		Item string `json:"item" binding:"required"`
	}
	if err := context.ShouldBindJSON(&updateData); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the product name
	existingProduct.Item = updateData.Item
	if err := db.Save(&existingProduct).Error; err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
		return
	}

	context.IndentedJSON(http.StatusOK, existingProduct)
}

func DeleteProduct(context *gin.Context) {
	db := database.GetDB()

	// Get the ID of the product to delete from the URL parameter
	id := context.Param("id")

	// Check if the product exists
	var existingProduct models.Product
	result := db.First(&existingProduct, id)
	if result.Error != nil {
		context.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	// Delete the product
	if err := db.Delete(&existingProduct, id).Error; err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}

	context.IndentedJSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}