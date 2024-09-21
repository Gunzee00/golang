package productcontroller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang/models"
	"gorm.io/gorm"
)

// Helper function to convert ID and handle errors
func getIdParam(c *gin.Context) (int, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return 0, err
	}
	return id, nil
}

// Get all products
func Index(c *gin.Context) {
	var products []models.Product

	// Query to get all products
	if err := models.DB.Find(&products).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	// Return products as JSON
	c.JSON(http.StatusOK, gin.H{"products": products})
}

// Get single product by ID
func Show(c *gin.Context) {
	id, err := getIdParam(c)
	if err != nil {
		return
	}

	var product models.Product

	// Query to find product by ID
	if err := models.DB.First(&product, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Product not found"})
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
		return
	}

	// Return product as JSON
	c.JSON(http.StatusOK, gin.H{"product": product})
}

// Create a new product
func Create(c *gin.Context) {
	var product models.Product

	// Bind incoming JSON to the product model
	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Insert product into database
	if err := models.DB.Create(&product).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	// Return created product
	c.JSON(http.StatusCreated, gin.H{"product": product})
}

// Update a product by ID
func Update(c *gin.Context) {
	id, err := getIdParam(c)
	if err != nil {
		return
	}

	var product models.Product

	// Find existing product by ID
	if err := models.DB.First(&product, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Product not found"})
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
		return
	}

	var updatedProduct models.Product

	// Bind incoming JSON to the updatedProduct model
	if err := c.ShouldBindJSON(&updatedProduct); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Update the product in the database (only update provided fields)
	if err := models.DB.Model(&product).Updates(updatedProduct).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	// Return updated product
	c.JSON(http.StatusOK, gin.H{"product": product})
}

// Delete a product by ID
func Delete(c *gin.Context) {
	id, err := getIdParam(c)
	if err != nil {
		return
	}

	var product models.Product

	// Find existing product by ID
	if err := models.DB.First(&product, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Product not found"})
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
		return
	}

	// Delete the product from the database
	if err := models.DB.Delete(&product).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	// Return success message
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}
