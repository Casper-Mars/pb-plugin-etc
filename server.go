package main

import (
	"github.com/gin-gonic/gin"
)

type UserHttpServer interface {
	list(ctx *gin.Context)
}

func RegistryUserHttpServer(engine *gin.Engine, srv UserHttpServer) {
	engine.Handle("GET", "", srv.list)
}

func server() {
	engine := gin.Default()
	engine.Run("localhost:8080")
}
