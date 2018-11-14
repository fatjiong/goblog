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
	c.HTML(http.StatusOK, "article/detail.html", gin.H{
		"articleDetail": articleDetail,
	})
}
