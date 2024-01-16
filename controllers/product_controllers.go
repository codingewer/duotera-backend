package controllers

import (
	"duotera/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	admin := models.Admin{}
	idHeaderValue := c.Request.Header.Get("ID")
	aid, err := strconv.ParseUint(idHeaderValue, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	aadmin, err := admin.FindAdminByUserId(uint(aid))
	if aadmin.Role != "ADMIN" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	product := models.Product{}
	c.BindJSON(&product)
	productSavedDB, err := product.SaveToDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, productSavedDB)
}

func UpdateProduct(c *gin.Context) {
	admin := models.Admin{}
	idHeaderValue := c.Request.Header.Get("ID")
	aid, err := strconv.ParseUint(idHeaderValue, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	aadmin, err := admin.FindAdminByUserId(uint(aid))
	if aadmin.Role != "ADMIN" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	product := models.Product{}
	c.BindJSON(&product)
	id := c.Param("id")
	pid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	updatedProduct, err := product.UpdateProduct(pid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedProduct)
}

// remove product from db
func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	// id to uid
	pid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	product := models.Product{}
	err = product.RemoveFromDb(pid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "product deleted"})
}

func GetProduct(c *gin.Context) {
	product := models.Product{}
	id := c.Param("id")
	pid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	productGeted, err := product.GetProductByID(uint(pid))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, productGeted)
}

func GetProducts(c *gin.Context) {
	product := models.Product{}
	productsGeted, err := product.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, productsGeted)
}
