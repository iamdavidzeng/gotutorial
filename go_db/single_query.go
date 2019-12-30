package go_db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type NewTag struct {
	ID    int
	Email string
}

func SimpleQuery() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/users")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	var tag NewTag
	err = db.QueryRow("SELECT id, email FROM users.users WHERE id=?", 2).Scan(&tag.ID, &tag.Email)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(tag.ID)
	fmt.Println(tag.Email)
}
