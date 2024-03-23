package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

type User struct {
	ID    int
	Name  string
	Email string
	Pass  string
}

func mysqlConnection() (*sql.DB, error) {
	cfg := mysql.Config{
		User:   "go",
		Passwd: "go",
		Net:    "tcp",
		Addr:   "mysql:3306",
		DBName: "go",
	}

	return sql.Open("mysql", cfg.FormatDSN())
}

func index(c *gin.Context) {

	db, err := mysqlConnection()

	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
	}

	rows, err := db.Query("SELECT * FROM users")

	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
	}

	retorno := []User{}

	for rows.Next() {
		var user User

		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Pass)

		if err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
		}
		
		retorno = append(retorno, user)
	}

	c.JSON(200, gin.H{
		"data": retorno,
	})
}

func main() {

	router := gin.Default()

	router.GET("/users", index);

	router.Run(":8080")
}
