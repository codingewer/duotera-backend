package main

import (
	"duotera/controllers"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	//gin
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Hello")
	})
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization", "ID", "id"}
	//config.AllowOrigins = []string{"http://localhost:3001"}
	router.Use(cors.New(config))

	admin := router.Group("/admin")
	admin.POST("/register", controllers.CreateUser)
	admin.POST("/login", controllers.AdminLogin)
	admin.PUT("/update-password", controllers.AdminUpdatePassword)

	product := router.Group("/product")
	product.GET("/byid/:id", controllers.GetProduct)
	product.PUT("/update/:id", controllers.UpdateProduct)
	product.GET("/all", controllers.GetProducts)
	product.POST("/new", controllers.CreateProduct)
	product.DELETE("/delete/:id", controllers.DeleteProduct)

	dealership := router.Group("/dealership")
	dealership.GET("/byid/:id", controllers.GetDealerShipById)
	dealership.POST("/new", controllers.NewDealerShip)
	dealership.GET("/all", controllers.GetDealerShips)
	dealership.GET("/active", controllers.GetDealerShipsActive)
	dealership.GET("/activeFalse", controllers.GetDealerShipsActiveFalse)
	dealership.PUT("/update/:id", controllers.IsActiveUpdate)

	details := router.Group("/details")
	details.GET("/byname/:name", controllers.DetailGetByName)
	details.PUT("/update/:name", controllers.DetailUpdateByName)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := router.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}

}
