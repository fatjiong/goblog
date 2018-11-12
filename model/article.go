package model

import "github.com/jinzhu/gorm"

type Article struct {
	gorm.Model
	CategoryId  uint
	Title       string
	Thumb       string
	Author      string
	Description string
	Status      uint
	IsRecommend uint
	Hits        uint
	Body        string
}
