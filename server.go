package main

import (
	"gin002/controller"
	"gin002/middlewares"
	"gin002/service"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	VideoService    service.VideoService       = service.New()
	VideoController controller.VideoController = controller.New(VideoService)
)

func setUpLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
func main() {

	setUpLogOutput()
	server := gin.New()

	server.Static("/css", "./templates/css")
	server.LoadHTMLGlob("templates/*.html")
	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth(), gindump.Dump())

	apiRoutes := server.Group("/api")
	{
		apiRoutes.GET("/videos", func(ctx *gin.Context) {
			ctx.JSON(200, VideoController.FindAll())
		})

		apiRoutes.POST("/videos", func(ctx *gin.Context) {
			err := VideoController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"message": "Video Input is Valid!",
				})
			}
		})
	}

	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", VideoController.ShowAll)
	}
	server.Run(":8080")
}
