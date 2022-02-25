package service

type LoginService interface {
	Login(username, password string) bool
}

type loginService struct {
	username string
	password string
}

func NewLoginService() LoginService {
	return &loginService{
		username: "sakura_endo",
		password: "sakura_endo",
	}
}

func (service *loginService) Login(username, password string) bool {
	return service.username == username && service.password == password
}
