package model

import "github.com/jinzhu/gorm"

type Adver struct {
	gorm.Model
	Title     string
	Thumb     string
	Sort      uint
	LinkUrl   string
	ArticleId uint
}
