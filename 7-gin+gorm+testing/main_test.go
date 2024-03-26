package main

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {

	go main()

	// Index
	response, _ := http.Get("http://localhost:8080/cars")
	assert.Equal(t, http.StatusOK, response.StatusCode)

	// Create
	car := []byte(`{"model": "Fusca", "brand": "Volkswagen", "year": 1969}`)
	response, _ = http.Post("http://localhost:8080/cars", "application/json", bytes.NewBuffer(car))
	assert.Equal(t, http.StatusCreated, response.StatusCode)

	// Show
	response, _ = http.Get("http://localhost:8080/cars/1")
	assert.Equal(t, http.StatusOK, response.StatusCode)

}