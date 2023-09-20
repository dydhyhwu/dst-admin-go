package model

import "gorm.io/gorm"

type App struct {
	gorm.Model
	Id   int
	Name string
}
