package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xvbnm48/gin-course/entity"
	"github.com/xvbnm48/gin-course/service"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) error
	ShowAll(ctx *gin.Context)
}

type controller struct {
	service service.VideoService
}

func New(service service.VideoService) VideoController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) error {
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return err
	}
	c.service.Save(video)
	return nil
}

func (c *controller) ShowAll(ctx *gin.Context) {
	videos := c.service.FindAll()
	data := gin.H{
		"Title":  "Video Page",
		"videos": videos,
	}

	ctx.HTML(http.StatusOK, "index.html", data)
}
