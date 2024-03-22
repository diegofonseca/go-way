package main

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"time"
)

type User struct {
	ID   int
	Name string
	Email string
	Pass string
}

func main() {

	// Configurar banco de dados
	cfg := mysql.Config{
		User:   "go",
		Passwd: "go",
		Net:    "tcp",
		Addr:   "mysql:3306",
		DBName: "go",
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		panic(err.Error())
	}

	// infinite loop
	for {
		fmt.Println("Hello, World! From Docker!")

		rows, err := db.Query("SELECT * FROM users")

		if err != nil {
			panic(err.Error())
		}

		for rows.Next() {
			var user User

			err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Pass)

			if err != nil {
				panic(err.Error())
			}

			fmt.Println(user)
		}

		// sleep for 1 second
		time.Sleep(3 * time.Second)
	}
}
