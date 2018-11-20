package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"math/rand"
	"strconv"
)

var DB *gorm.DB

type Category struct {
	gorm.Model
	Pid  uint
	Name string
}

type Adver struct {
	gorm.Model
	Title     string
	Thumb     string
	Sort      uint
	LinkUrl   string
	ArticleId uint
}
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

type Comment struct {
	gorm.Model
	ArticleId uint
	CommentId uint
	IsCheck   uint
	Title     string
	Body      string
}

type Feedback struct {
	gorm.Model
	Name   string
	Email  string
	Title  string
	Body   string
	Remark string
}

type Sharelink struct {
	gorm.Model
	Name    string
	Thumb   string
	IsCheck uint
	LinkUrl string
}

/**
初始化数据库
*/
func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root:123456@/gblog?charset=utf8mb4&parseTime=True&loc=Local")
	if err == nil {
		DB = db
		//设置默认表名前缀
		gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
			return "gb_" + defaultTableName
		}
		db.AutoMigrate(&Category{}, &Article{}, &Adver{}, &Comment{}, &Feedback{}, &Sharelink{})
		return db, nil
	}
	return nil, err
}

/**
填充数据
*/
func GenerateData() {
	for i := 1; i < 11; i++ {
		DB.Create(&Category{
			Name: "测试分类" + strconv.Itoa(i),
			Pid:  0,
		})
		for j := 1; j < 11; j++ {
			DB.Create(&Category{
				Name: "测试分类" + strconv.Itoa(i) + "_" + strconv.Itoa(j),
				Pid:  uint(i),
			})
		}
	}

	for i := 1; i < 1001; i++ {
		DB.Create(&Article{
			CategoryId:  uint(rand.Intn(219) + 1),
			Title:       "标题" + strconv.Itoa(rand.Intn(1024)),
			Author:      "作者" + strconv.Itoa(rand.Intn(1024)),
			Thumb:       strconv.Itoa(rand.Int()),
			Description: "详情简介" + strconv.Itoa(rand.Intn(1024)),
			Status:      uint(rand.Intn(2)),
			IsRecommend: uint(rand.Intn(2)),
			Hits:        uint(rand.Intn(9999)),
			Body:        "内容" + strconv.Itoa(rand.Intn(1024)),
		})
	}
	return
}

/**
获取推荐文章
*/
func GetArticleRecommend(categoryId uint) ([]*Article, error) {
	var articleList []*Article
	var err error
	if categoryId == 0 {
		err = DB.Where("is_recommend = ?", 1).Where("status=?", 1).Offset(0).Limit(6).Find(&articleList).Error
	} else {
		err = DB.Where("is_recommend = ?", 1).Where("status=?", 1).Where("category_id = ?", categoryId).Offset(0).Limit(6).Find(&articleList).Error
	}

	return articleList, err
}

/**
获取点击前10的文章
*/
func GetArticleHits(limit int, categoryId uint) ([]*Article, error) {
	var articleList []*Article
	var err error
	if categoryId == 0 {
		err = DB.Where("is_recommend = ?", 1).Where("status=?", 1).Offset(0).Limit(limit).Order("hits desc").Find(&articleList).Error
	} else {
		err = DB.Where("is_recommend = ?", 1).Where("category_id = ?", categoryId).Where("status=?", 1).Offset(0).Limit(limit).Order("hits desc").Find(&articleList).Error
	}

	return articleList, err
}

/**
获取文章列表
*/
func GetArticleList(page int) ([]*Article, error) {
	var articleList []*Article
	var err error
	err = DB.Where("is_recommend = ?", 1).Offset((page - 1) * 10).Limit(10).Order("hits desc").Find(&articleList).Error
	return articleList, err
}

//获取文章详情
func GetArticleDetial(id string) (*Article, error) {
	aid, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
	}
	var article Article
	var err2 error
	err2 = DB.First(&article, aid).Error
	return &article, err2
}

/**
获取其他文章
*/
func GetArticlePrev(id string, source string) (*Article, error) {
	aid, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
	}
	var article Article
	var err2 error
	if source == "next" {
		err2 = DB.Where("id > ? and status = ?", aid, 1).First(&article).Error
	} else {
		err2 = DB.Where("id < ? and status = ?", aid, 1).Last(&article).Error
	}
	return &article, err2
}

/**
根据父类获取那啥
*/
func GetCategoryList(pid uint) ([]*Category, error) {
	var categoryList []*Category
	var err error
	err = DB.Where("pid = ?", pid).Find(&categoryList).Error
	return categoryList, err
}
