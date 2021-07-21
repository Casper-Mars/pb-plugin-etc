package main

import "github.com/gin-gonic/gin"

type UserApi struct {
}

func (receiver UserApi) Registry(engine *gin.Engine) {
	engine.Handle("GET", "", receiver.list)
}

func (receiver UserApi) list(ctx *gin.Context) {

}

func server() {
	engine := gin.Default()
	engine.Run("localhost:8080")
}
