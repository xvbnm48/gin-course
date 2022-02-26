package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/xvbnm48/gin-course/dto"
	"github.com/xvbnm48/gin-course/service"
)

type LoginController interface {
	Login(ctx *gin.Context) string
}
type loginController struct {
	loginService service.LoginService
	jwtService   service.JWTService
}

func NewLoginService(loginService service.LoginService, jwtService service.JWTService) LoginController {
	return &loginController{
		loginService: loginService,
		jwtService:   jwtService,
	}
}

func (controller *loginController) Login(ctx *gin.Context) string {
	var credentials dto.Credentials
	err := ctx.ShouldBind(&credentials)
	if err != nil {
		return " "
	}

	isAuthenticated := controller.loginService.Login(credentials.Username, credentials.Password)
	if isAuthenticated {
		return controller.jwtService.GenerateToken(credentials.Username, true)
	}
	return ""
}
