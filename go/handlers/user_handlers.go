package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"webapp-go/models"
	"webapp-go/services"
)

func userCreate(c *gin.Context) {
	var userData models.UserCreate
	err := c.ShouldBind(userData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	userDetail := services.UserCreate(userData)
	c.JSON(http.StatusCreated, userDetail)
}

func userDetail(c *gin.Context) {
	var userID string = c.Params.ByName("id")
	userDetail := services.UserDetail(userID)
	c.JSON(http.StatusOK, userDetail)
}

func applyUserRoutes(rg *gin.RouterGroup) {
	rg.POST("/users", userCreate)
	rg.GET("/user/:id", userDetail)
	//rg.GET("/users", createUser)
}
