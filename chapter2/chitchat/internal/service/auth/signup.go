package auth

import (
	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/service"
)

type SignUpService struct {
	userService *service.UserService
}

func NewSignUpService(service *service.UserService) *SignUpService {
	return &SignUpService{
		userService: service,
	}
}

func (ss *SignUpService) SignUp() {
}
