package model

type Todo struct {
	gorm.Model
	Id   int    `gorm:"primary_key"`
	Name string `gorm:"size:10"`
}
