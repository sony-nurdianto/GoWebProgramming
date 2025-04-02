package auth

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/database"
	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/encryption"
	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/genrators"
	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/models"
	cr "github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/repository/cache"
	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/service"
	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/service/cache"
)

var (
	ErrUnAuthorizedUser error = errors.New("unauthorized access deniend wrong password")
	ErrUserNotFound     error = errors.New("user not found")
	ErrEncryptToken     error = errors.New("failed to encrypt token")
)

type LoginUpService struct {
	userService *service.UserService
	cache       *database.Cache
}

func NewLoginUpService(service *service.UserService, cache *database.Cache) *LoginUpService {
	return &LoginUpService{
		userService: service,
		cache:       cache,
	}
}

func (ls *LoginUpService) AuthenticateLogin(email string, password string) (string, error) {
	user, err := ls.userService.GetUserByEmail(email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", fmt.Errorf("%w", ErrUserNotFound)
		}
		return "", err
	}

	log.Println("User found: ", user.Email)

	isUser, err := encryption.VerifyPassword(password, user.Password)
	if err != nil {
		return "", err
	}

	if !isUser {
		return "", ErrUnAuthorizedUser
	}

	tokenID := genrators.CreateUUID()

	token, err := encryption.CreateWebToken(tokenID)
	if err != nil {
		return "", fmt.Errorf("%w: %v", ErrEncryptToken, err)
	}

	log.Println("token: ", token)

	session := models.Session{
		Id:        user.Id,
		Uuid:      user.Uuid,
		Email:     user.Email,
		UserId:    user.Id,
		CreatedAt: user.CreatedAt,
	}

	sessionRepo := cr.NewSessionRepo(ls.cache)
	sessionService := cache.NewSessionService(sessionRepo)

	if err := sessionService.SetSession(tokenID, session, 1*time.Hour); err != nil {
		return "", err
	}

	return token, nil
}
