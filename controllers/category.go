package controllers

import (
	"github.com/fatjiong/goblog/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

/**
分类列表
*/
func CategoryGet(c *gin.Context) {
	cid, _ := strconv.Atoi(c.Param("cid"))

	//获取文章列表
	articleList, _ := model.GetArticleListByCategoryId(10, uint(cid))
	//分类列表
	categoryList, _ := model.GetCategoryList(0)
	//推荐文章列表
	recommendList, _ := model.GetArticleRecommend(uint(cid))
	//获取点击排行
	hitsList, _ := model.GetArticleHits(10, uint(cid))

	c.HTML(http.StatusOK, "search/result.html", gin.H{
		"articleList":   articleList,
		"categoryList":  categoryList,
		"recommendList": recommendList,
		"hitsList":      hitsList,
	})
}
