package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"webapp-go/models"
	"webapp-go/services"
)

func articleCreate(c *gin.Context) {
	var data models.ArticleData
	err := c.ShouldBind(data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	articleDetail := services.ArticleCreate(data)
	c.JSON(http.StatusCreated, articleDetail)
}

func articleDetail(c *gin.Context) {
	var slug string = c.Params.ByName("slug")
	articleDetail := services.ArticleDetail(slug)
	c.JSON(http.StatusOK, articleDetail)
}

func articleList(c *gin.Context) {
        articles := services.ArticleList()
        c.JSON(http.StatusOK, articles)
}


func applyArticleRoutes(rg *gin.RouterGroup) {
	rg.POST("/articles", articleCreate)
	rg.GET("/articles/:slug", articleDetail)
        rg.GET("/articles", articleList)
}
