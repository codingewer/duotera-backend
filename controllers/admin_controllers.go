package controllers

import (
	"duotera/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//save to db gin

func CreateUser(c *gin.Context) {
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
	c.BindJSON(&admin)
	adminCreated, err := admin.SaveToDb()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "veri tabanına kayıt hatası"})
		return
	}
	c.JSON(http.StatusOK, adminCreated)
}

func AdminLogin(c *gin.Context) {
	admin := models.Admin{}
	c.BindJSON(&admin)
	addmin, err := admin.SignIn(admin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": err.Error()})
		return
	}
	var adminResp = models.AdminResponse{}
	if admin.Password != addmin.Password {
		c.JSON(http.StatusInternalServerError, gin.H{"status": err.Error()})
		return
	}
	if addmin.Role != "ADMIN" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "Yetkisi yok"})
	}
	if admin.Password == addmin.Password {
		adminResp.Username = addmin.Username
		adminResp.Token = addmin.ID
	}
	c.JSON(http.StatusOK, adminResp)
}
