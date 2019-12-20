package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Go MySQL Tutorial")

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/users")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	results, err := db.Query("SELECT * FROM users.users")
	if err != nil {
		panic(err.Error())
	}

	defer results.Close()
}
