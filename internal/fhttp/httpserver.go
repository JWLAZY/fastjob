package fhttp

import (
	"strconv"

	router "gitlab.xuanke.com/live/fastjob/internal/fhttp/router"

	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	engine *gin.Engine
	port   int
}

func NewHttpServer(port int) *HttpServer {
	return &HttpServer{port: port}
}

func (server *HttpServer) Start() {
	server.engine = gin.Default()
	// server.engine.GET("/start", handlers ...gin.HandlerFunc)
	router.GetRoutes(server.engine)
	server.engine.Run(":" + strconv.Itoa(server.port))
}
