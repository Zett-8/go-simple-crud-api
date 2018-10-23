package models

import "github.com/jinzhu/gorm"

type Todo struct {
	gorm.Model
	Id   int    `gorm:"primary_key"`
	Name string `gorm:"size:10"`
}

func connectDB() *gorm.DB {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres password=postgres sslmode=disable")

	if err != nil {
		panic(err)
	}

	return db
}
