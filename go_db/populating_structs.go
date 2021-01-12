package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Tag struct {
	ID    int    `json: "id"`
	Email string `json: "email"`
}

func SimpleStructure() {

	fmt.Println("Go MYSQL Tutorial")

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/users")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// Execute the query
	results, err := db.Query("SELECT id, email FROM users.users")
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var tag Tag
		err = results.Scan(&tag.ID, &tag.Email)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println(tag.Email)
	}
}
