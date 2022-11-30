package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func addTaskRoutes(rg *gin.RouterGroup) {
	tasks := rg.Group("/task")

	tasks.GET("/list", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "user")
	})

	tasks.POST("/create", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "add ok")
	})

	tasks.POST("/update", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "update ok")
	})

}
