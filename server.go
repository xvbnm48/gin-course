package main

import (
	"github.com/xvbnm48/gin-course/middlewares"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/xvbnm48/gin-course/controller"
	"github.com/xvbnm48/gin-course/service"
)

var (
	videoService service.VideoService = service.New()
	loginService service.LoginService = service.NewLoginService()
	jwtService   service.JWTService   = service.NewJWTService()

	videoController controller.VideoController = controller.New(videoService)
	loginController controller.LoginController = controller.NewLoginService(loginService, jwtService)
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

	//server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth())
	server.Use(gin.Recovery(), middlewares.Logger())
	//login end point with auth + token
	server.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{"token": token})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})

	// jwt auth only route to "api" only
	apiRoutes := server.Group("/api", middlewares.AuthorizeJWT())
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

	// this route is public non use auth
	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", videoController.ShowAll)
	}

	// server.Run(":8080")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	server.Run(":" + port)

}
