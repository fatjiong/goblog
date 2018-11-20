package controllers

import (
	"fmt"
	"github.com/fatjiong/goblog/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

//文章详情
func ArticleDetail(c *gin.Context) {
	id := c.Param("id")

	articleDetail, err := model.GetArticleDetial(id)
	if err != nil {
		fmt.Println(err)
	}

	//推荐文章列表
	recommendList, err := model.GetArticleRecommend(articleDetail.CategoryId)
	if err != nil {
		fmt.Println(err)
	}

	//获取点击排行
	hitsList, err := model.GetArticleHits(10, articleDetail.CategoryId)
	if err != nil {
		fmt.Println(err)
	}

	/**
	获取上一篇文章
	*/
	prevArticle, err := model.GetArticlePrev(id, "prev")
	if err != nil {
		fmt.Println(err)
		prevArticle = nil
	}

	/**
	获取下一篇文章
	*/
	nextArticle, err := model.GetArticlePrev(id, "next")
	if err != nil {
		fmt.Println(err)
		nextArticle = nil
	}

	c.HTML(http.StatusOK, "article/detail.html", gin.H{
		"articleDetail": articleDetail,
		"recommendList": recommendList,
		"hitsList":      hitsList,
		"prevArticle":   prevArticle,
		"nextArticle":   nextArticle,
	})
}
