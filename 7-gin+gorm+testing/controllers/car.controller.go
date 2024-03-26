package controllers

import (
	"github.com/diegofonseca/go-way/models"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/diegofonseca/go-way/utils"
)


func Index(c *gin.Context) {
	db, err := utils.SetupDB()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	cars := []models.Car{}

	db.Find(&cars)

	c.JSON(http.StatusOK, gin.H{
		"cars": cars,
	})
}

func Create(c *gin.Context) {
	db, err := utils.SetupDB()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}

	car := models.Car{}

	c.BindJSON(&car)

	db.Create(&car)

	c.JSON(http.StatusCreated, gin.H{
		"car": car,
	})
}

func Show(c *gin.Context) {
	db, err := utils.SetupDB()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}

	car := models.Car{}

	db.First(&car, c.Param("id"))

	if car.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Car not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"car": car,
	})
}

func Update(c *gin.Context) {
	db, err := utils.SetupDB()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}

	car := models.Car{}

	db.First(&car, c.Param("id"))

	c.BindJSON(&car)

	db.Save(&car)

	c.JSON(http.StatusOK, gin.H{
		"car": car,
	})
}

func Delete(c *gin.Context) {
	db, err := utils.SetupDB()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}

	car := models.Car{}

	db.First(&car, c.Param("id"))

	if car.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Car not found",
		})
		return
	}

	db.Delete(&car)

	c.JSON(http.StatusNoContent, gin.H{})
}
