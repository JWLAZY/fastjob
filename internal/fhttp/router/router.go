package router

import "github.com/gin-gonic/gin"

func GetRoutes(engine *gin.Engine) {
	v1 := engine.Group("v1")
	addTaskRoutes(v1)
	addJobRoutes(v1)
	addNodeRoutes(v1)
	addQueueRoutes(v1)
}
