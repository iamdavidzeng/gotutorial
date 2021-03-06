package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type (
	todoModel struct {
		gorm.Model
		Title     string `json:"title"`
		Completed int    `json:"completed"`
	}

	transformTodo struct {
		ID        uint   `json:"id"`
		Title     string `json:title`
		Completed bool   `json:completed`
	}
)

func init() {
	var err error
	db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/demo")
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&todoModel{})
}

func main() {
	print("hello, world!")
}
