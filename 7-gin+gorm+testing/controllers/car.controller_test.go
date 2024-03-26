package controllers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"os"

	"github.com/diegofonseca/go-way/models"
	"github.com/diegofonseca/go-way/utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {

	prepareTest()

	router := gin.Default()
	router.GET("/cars", Index)

	req, _ := http.NewRequest("GET", "/cars", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	os.Setenv("DSN", "root:root@tcp(localhost:3306)/go-way")

	req, _ = http.NewRequest("GET", "/cars", nil)
	w = httptest.NewRecorder()

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestShow(t *testing.T) {

	prepareTest()

	createItem()

	router := gin.Default()
	router.GET("/cars/:id", Show)

	req, _ := http.NewRequest("GET", "/cars/1", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	req, _ = http.NewRequest("GET", "/cars/2", nil)
	w = httptest.NewRecorder()

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)

	os.Setenv("DSN", "root:root@tcp(localhost:3306)/go-way")

	req, _ = http.NewRequest("GET", "/cars/1", nil)
	w = httptest.NewRecorder()

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestCreate(t *testing.T) {

	prepareTest()

	router := gin.Default()
	router.POST("/cars", Create)

	body := []byte(`{"model": "Fusca", "brand": "Volkswagen", "year": 1969}`)

	req, _ := http.NewRequest("POST", "/cars", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)

	os.Setenv("DSN", "root:root@tcp(localhost:3306)/go-way")

	req, _ = http.NewRequest("POST", "/cars", bytes.NewBuffer(body))
	w = httptest.NewRecorder()

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestUpdate(t *testing.T) {
	
	prepareTest()
	createItem()

	db, _ := utils.SetupDB()

	router := gin.Default()
	router.PUT("/cars/:id", Update)

	body := []byte(`{"model": "Fusca", "brand": "Volkswagen", "year": 1969}`)

	car := models.Car{}
	db.First(&car)
	id := strconv.Itoa(int(car.ID))

	req, _ := http.NewRequest("PUT", "/cars/"+id, bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	os.Setenv("DSN", "root:root@tcp(localhost:3306)/go-way")

	req, _ = http.NewRequest("PUT", "/cars/"+id, bytes.NewBuffer(body))
	w = httptest.NewRecorder()

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestDelete(t *testing.T) {

	prepareTest()
	createItem()

	router := gin.Default()
	router.DELETE("/cars/:id", Delete)

	req, _ := http.NewRequest("DELETE", "/cars/1", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNoContent, w.Code)

	req, _ = http.NewRequest("DELETE", "/cars/0", nil)
	w = httptest.NewRecorder()

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)

	os.Setenv("DSN", "root:root@tcp(localhost:3306)/go-way")

	req, _ = http.NewRequest("DELETE", "/cars/1", nil)
	w = httptest.NewRecorder()

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func prepareTest() {

	os.Setenv("DSN", "go:go@tcp(mysql:3306)/go?charset=utf8mb4&parseTime=True&loc=Local")
	
	db, _ := utils.SetupDB()

	db.Migrator().DropTable(&models.Car{})
	db.AutoMigrate(&models.Car{})
}


func createItem() {
	db, _ := utils.SetupDB()

	car := models.Car{
		Modelo: "Fusca",
		Marca:  "Volkswagen",
		Ano:    1969,
	}

	db.Create(&car)
}
