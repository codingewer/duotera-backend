package controllers

import (
	"duotera/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DetailGetByName(c *gin.Context) {
	name := c.Param("name")
	detail := models.Detail{}
	detailGeted, err := detail.GetDetailByName(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, detailGeted)
}

// update by name
func DetailUpdateByName(c *gin.Context) {
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
	name := c.Param("name")
	detail := models.Detail{}
	c.BindJSON(&detail)
	detailUpdated, err := detail.UpdateDetailByName(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, detailUpdated)
}
