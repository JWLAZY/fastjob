package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func addNodeRoutes(rg *gin.RouterGroup) {
	nodes := rg.Group("/node")

	nodes.GET("/list", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "user")
	})

	nodes.POST("/create", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "add ok")
	})
}
