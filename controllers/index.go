package controllers

import (
	"fmt"
	"github.com/fatjiong/goblog/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 首页控制器
func IndexGet(c *gin.Context) {
	//分类列表
	categoryList, err := model.GetCategoryList(0)
	if err != nil {
		fmt.Println(err)
	}
	//推荐文章列表
	recommendList, err := model.GetArticleRecommend()
	if err != nil {
		fmt.Println(err)
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"categoryList":  categoryList,
		"recommendList": recommendList,
	})
}
