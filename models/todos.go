package models

import "github.com/jinzhu/gorm"

type Todo struct {
	gorm.Model
	Name string `gorm:"size:10; not null"`
}
