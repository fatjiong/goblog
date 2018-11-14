package main

import (
	"fmt"
	"github.com/fatjiong/goblog/model"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	//router := gin.Default()
	////路由定义
	//router.Static("/static/", "./static")
	//router.LoadHTMLGlob("views/*")
	//router.GET("/", controllers.IndexGet)

	//初始化数据库
	db, err := model.InitDB()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	//categoryList, err := model.GetCategoryList(0)
	//
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//for _, category := range categoryList {
	//	fmt.Println(category.Name)
	//}

	//生成点数据
	model.GenerateData()

	articleList, err := model.GetArticleRecommend()
	if err != nil {
		fmt.Println(err)
	}

	for _, article := range articleList {
		fmt.Println(article.Title + "_" + article.Author)
	}

	////发布文章页面
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
	//router.Run(":8080")
}
