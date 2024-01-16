package controllers

import (
	"duotera/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func NewDealerShip(c *gin.Context) {
	dealership := models.Dealership{}
	c.BindJSON(&dealership)
	dealerSaved, err := dealership.Create()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, dealerSaved)

}

func GetDealerShips(c *gin.Context) {
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
	dealership := models.Dealership{}
	dealerShips, err := dealership.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, dealerShips)
}

func GetDealerShipsActive(c *gin.Context) {
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
	dealership := models.Dealership{}
	dealerShips, err := dealership.FindByIsActive()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, dealerShips)
}

func GetDealerShipsActiveFalse(c *gin.Context) {
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
	dealership := models.Dealership{}
	dealerShips, err := dealership.FindByIsActiveFalse()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, dealerShips)
}

func IsActiveUpdate(c *gin.Context) {
	admin := models.Admin{}
	idHeaderValue := c.Request.Header.Get("ID")
	aid, err := strconv.ParseUint(idHeaderValue, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unauthorized"})
		return
	}
	aadmin, err := admin.FindAdminByUserId(uint(aid))
	if aadmin.Role != "ADMIN" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	dealership := models.Dealership{}
	i := c.Param("id")
	did, err := strconv.ParseUint(i, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "id okunmadı" + i})
		return
	}
	dealerSaved, err := dealership.IsActiveChange(did)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, dealerSaved)
}

func GetDealerShipById(c *gin.Context) {
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
	dealership := models.Dealership{}
	i := c.Param("id")
	did, err := strconv.ParseUint(i, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	dealerSaved, err := dealership.FindById(did)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, dealerSaved)
}
