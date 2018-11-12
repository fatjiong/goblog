package model

import "github.com/jinzhu/gorm"

type Sharelink struct {
	gorm.Model
	Name    string
	Thumb   string
	IsCheck uint
	LinkUrl string
}
