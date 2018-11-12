package model

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model
	Pid  uint
	Name string
}
