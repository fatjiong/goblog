package controllers

import (
	"github.com/fatjiong/goblog/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AboutGet(c *gin.Context) {
	//分类列表
	categoryList, _ := model.GetCategoryList(0)

	c.HTML(http.StatusOK, "about/about.html", gin.H{
		"categoryList": categoryList,
	})
}
