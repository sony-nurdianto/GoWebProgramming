package auth

import (
	"log"

	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/encryption"
	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/service/cache"
)

type LogoutService struct {
	session *cache.SessionService
}

func NewLogoutService(session *cache.SessionService) *LogoutService {
	return &LogoutService{
		session: session,
	}
}

func (ls *LogoutService) Logout(hashToken string) error {
	token, err := encryption.VerifyWebToken(hashToken)
	if err != nil {
		return err
	}

	log.Println(token)

	if err := ls.session.DeleteSession(token.Subject); err != nil {
		return err
	}

	log.Println("success delete Session")

	return nil
}
