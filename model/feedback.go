package model

import "github.com/jinzhu/gorm"

type Feedback struct {
	gorm.Model
	Name   string
	Email  string
	Title  string
	Body   string
	Remark string
}
