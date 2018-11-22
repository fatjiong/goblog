package controllers

import (
	"github.com/fatjiong/goblog/model"
	"github.com/gin-gonic/gin"
	"net/http"
	//"strconv"
)

// 首页控制器
func IndexGet(c *gin.Context) {
	//page, _ := strconv.Atoi(c.Param("page"))
	//page := c.Param("page")

	//分类列表
	categoryList, _ := model.GetCategoryList(0)

	//推荐文章列表
	recommendList, _ := model.GetArticleRecommend(0)

	//获取点击排行
	hitsList, _ := model.GetArticleHits(10, 0)

	//获取文章列表
	articleList, _ := model.GetArticleList(0)

	sharelinkList, _ := model.GetSharelinkList(3)

	c.HTML(http.StatusOK, "index/index.html", gin.H{
		"categoryList":  categoryList,
		"recommendList": recommendList,
		"hitsList":      hitsList,
		"articleList":   articleList,
		"sharelinkList": sharelinkList,
	})
}
