package main

import (
	"fmt"
	"github.com/fatjiong/goblog/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	//router := gin.Default()
	////路由定义
	//router.Static("/static/", "./static")
	//router.LoadHTMLGlob("views/*")
	//router.GET("/", controllers.IndexGet)

	db, err := gorm.Open("mysql", "root:123456@/gblog?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}

	//设置默认表名前缀
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "gb_" + defaultTableName
	}

	//关闭数据库
	defer db.Close()

	//初始化数据库
	db.AutoMigrate(&model.Category{}, &model.Article{}, &model.Adver{}, &model.Comment{}, &model.Feedback{}, &model.Sharelink{})
	//
	//for i := 1; i < 11; i++ {
	//	db.Create(&model.Category{
	//		Name: "测试分类" + strconv.Itoa(i),
	//		Pid:  0,
	//	})
	//	for j := 1; j < 11; j++ {
	//		db.Create(&model.Category{
	//			Name: "测试分类" + strconv.Itoa(i) + "_" + strconv.Itoa(j),
	//			Pid:  uint(i),
	//		})
	//	}
	//}

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
