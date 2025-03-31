package service

import (
	"context"
	"time"

	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/encryption"
	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/genrators"
	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/models"
	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/repository"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(ur *repository.UserRepository) *UserService {
	return &UserService{
		userRepo: ur,
	}
}

func (ur *UserService) CreateUser(user models.User) (err error) {
	user.CreatedAt = time.Now()
	user.Uuid = genrators.CreateUUID()
	user.Password, err = encryption.HashPassword(user.Password)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = ur.userRepo.NewUser(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (ur *UserService) GetUserByEmail(email string) (user models.User, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	user, err = ur.userRepo.GetAuthUser(ctx, email)
	if err != nil {
		return user, err
	}

	return user, nil
}
