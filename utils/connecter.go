package utils

import "github.com/jinzhu/gorm"

func ConnectDB() *gorm.DB {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres password=postgres sslmode=disable")

	if err != nil {
		panic(err)
	}

	return db
}
