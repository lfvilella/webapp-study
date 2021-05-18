package handlers

import "github.com/gin-gonic/gin"

func Run() {
	router := gin.Default()
	rg := router.Group("")
	applyUserRoutes(rg)
	router.Run()
}
