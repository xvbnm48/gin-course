package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/xvbnm48/gin-course/controller"
	middlewares "github.com/xvbnm48/gin-course/middlewares"
	"github.com/xvbnm48/gin-course/service"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {

	setupLogOutput()
	server := gin.New()

	server.Static("/css", "./templates/css")
	server.LoadHTMLGlob("templates/*.html")

	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth())
	apiRoutes := server.Group("/api")
	{
		apiRoutes.GET("/posts", func(c *gin.Context) {
			c.JSON(200, videoController.FindAll())
		})
		apiRoutes.POST("/posts", func(c *gin.Context) {
			err := videoController.Save(c)
			if err != nil {
				c.JSON(400, gin.H{"error": err.Error()})
			} else {
				c.JSON(200, gin.H{"message": "video input is valid"})
			}

		})
	}

	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", videoController.ShowAll)
	}

	// server.Run(":8080")
	port := os.Getenv("PORT")
	server.Run(":" + port)
}
