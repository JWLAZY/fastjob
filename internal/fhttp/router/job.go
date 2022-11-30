package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func addJobRoutes(rg *gin.RouterGroup) {
	jobs := rg.Group("/job")

	jobs.GET("/list", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "user")
	})

	jobs.POST("/create", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "add ok")
	})

}
