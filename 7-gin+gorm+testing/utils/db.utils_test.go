package utils

import (
	"testing"
	"github.com/joho/godotenv"
	"os"
)



func TestSetupDB(t *testing.T) {

	godotenv.Load("../.env")
	
	_, err := SetupDB()
	if err != nil {
		t.Error("Failed to connect to database!")
	}
	t.Log("Connected to database!")
}

func TestSetupDBFailDBContaction(t *testing.T) {

	os.Setenv("DSN", "root:root@tcp(localhost:3306)/go-way")

	_, err := SetupDB()
	
	if err == nil {
		t.Error("Failed to connect to database!")
	}
}
