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
	router.LoadHTMLGlob("views/*")
	router.GET("/", controllers.IndexGet)

	//发布文章页面
	//router.GET("/article/new/", controller.NewArticle)
	////文章提交接口
	//router.POST("/article/submit/", controller.ArticleSubmit)
	////文章详情页
	//router.GET("/article/detail/", controller.ArticleDetail)
	//
	////文件上传接口
	//router.POST("/upload/file/", controller.UploadFile)
	//
	////留言页面
	//router.GET("/leave/new/", controller.LeaveNew)
	////关于我页面
	//router.GET("/about/me/", controller.AboutMe)
	router.Run(":8080")
}
