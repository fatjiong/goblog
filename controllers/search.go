package controllers

import (
	"github.com/fatjiong/goblog/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

//  查询结果
func SearchPost(c *gin.Context) {
	source := c.Param("source")
	keywords := c.PostForm("keywords")
	keywords = strings.Trim(keywords, "")

	var articleList []*model.Article
	if source == "tag" {
		articleList, _ = model.GetArticleListByTag(keywords, 1)
	} else {
		articleList, _ = model.GetArticleListByKeywords(keywords, 1)
	}

	//推荐文章列表
	recommendList, _ := model.GetArticleRecommend(0)
	//获取点击排行
	hitsList, _ := model.GetArticleHits(6, 0)
	//分类列表
	categoryList, _ := model.GetCategoryList(0)

	c.HTML(http.StatusOK, "search/result.html", gin.H{
		"articleList":   articleList,
		"recommendList": recommendList,
		"hitsList":      hitsList,
		"categoryList":  categoryList,
	})
}
