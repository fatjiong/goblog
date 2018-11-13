package model

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model
	Pid  uint
	Name string
}

var DB *gorm.DB
var CategoryList []Category

func GetCategoryList(pid int) []Category {
	DB.Where("pid = ?", pid).Find(&Category{})
	return CategoryList
}
