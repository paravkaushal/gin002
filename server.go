package main

import (
	"gin002/controller"
	"gin002/middlewares"
	"gin002/service"

	"github.com/gin-gonic/gin"
)

var (
	VideoService    service.VideoService       = service.New()
	VideoController controller.VideoController = controller.New(VideoService)
)

func main() {
	server := gin.New()

	server.Use(gin.Recovery(), middlewares.Logger())

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, VideoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, VideoController.Save(ctx))
	})
	server.Run(":8080")
}
