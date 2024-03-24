package main

import (
	"github.com/gin-gonic/gin"
	"github.com/diegofonseca/go-way/controllers"
)

func main() {

	router := gin.Default()

	router.GET("/cars", controllers.Index);
	router.POST("/cars", controllers.Create);
	router.GET("/cars/:id", controllers.Show);
	router.PUT("/cars/:id", controllers.Update);
	router.DELETE("/cars/:id", controllers.Delete);

	router.Run(":8080")
}
