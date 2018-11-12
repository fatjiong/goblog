package model

import "github.com/jinzhu/gorm"

type Comment struct {
	gorm.Model
	ArticleId uint
	CommentId uint
	IsCheck   uint
	Title     string
	Body      string
}
