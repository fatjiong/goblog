package main

import (
	"fmt"
	"github.com/fatjiong/goblog/controllers"
	"github.com/fatjiong/goblog/model"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"html/template"
	"time"
)

//格式化时间戳
func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	//初始化数据库
	db, err := model.InitDB()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	//生成点数据
	//model.GenerateData()

	router := gin.Default()
	router.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})
	//路由定义
	router.Static("/static/", "./static")
	router.LoadHTMLGlob("views/**/*")
	//项目首页
	router.GET("/", controllers.IndexGet)
	//分类首页
	router.GET("/category/:cid", controllers.CategoryGet)
	//文章详情页
	router.GET("/article/:id", controllers.ArticleDetail)
	// 更改文章的计数
	router.POST("/article/counter", controllers.ArticleCounter)
	//关键字查询
	router.POST("/search", controllers.SearchPost)
	//关于我页面
	router.GET("/about", controllers.AboutGet)
	////留言页面
	//router.GET("/leave/new/", controller.Feedback)
	router.Run(":8080")
}
