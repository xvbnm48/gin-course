package main

import (
	"github.com/gin-gonic/gin"
	"github.com/xvbnm48/gin-course/controller"
	"github.com/xvbnm48/gin-course/service"
)

var (
	videoService service.VideoService = service.New()
	videoController controller.VideoController = controller.New(videoService)
)
func main() {
	server := gin.Default()

	server.GET("/posts", func(c *gin.Context) {
		c.JSON(200, videoController.FindAll())
	})
	server.POST("/posts", func(c *gin.Context) {
		c.JSON(200, videoController.Save(c))
	})
	server.Run(":8080")
}
