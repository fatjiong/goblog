package controllers

import (
	"fmt"
	"github.com/fatjiong/goblog/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// 首页控制器
func SearchPost(c *gin.Context) {
	keywords := c.PostForm("keywords")
	keywords = strings.Trim(keywords, "")

	articleList, err := model.GetArticleListByKeywords(keywords, 1)
	if err != nil {
		fmt.Println(err)
	}

	//推荐文章列表
	recommendList, err := model.GetArticleRecommend(0)
	if err != nil {
		fmt.Println(err)
	}

	//获取点击排行
	hitsList, err := model.GetArticleHits(6, 0)
	if err != nil {
		fmt.Println(err)
	}

	c.HTML(http.StatusOK, "search/result.html", gin.H{
		"articleList":   articleList,
		"recommendList": recommendList,
		"hitsList":      hitsList,
	})
}
