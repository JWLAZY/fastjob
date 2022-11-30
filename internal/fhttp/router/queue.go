package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func addQueueRoutes(rg *gin.RouterGroup) {
	queues := rg.Group("/queue")

	queues.GET("/count", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "user")
	})
}
