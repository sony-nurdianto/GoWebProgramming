package service

import (
	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/models"
	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/repository"
)

type ThreadService struct {
	threadRepo *repository.ThreadRepository
}

func NewThreadService(tr *repository.ThreadRepository) *ThreadService {
	return &ThreadService{
		threadRepo: tr,
	}
}

func (t *ThreadService) Threads() (threads []models.Thread, err error) {
	return t.threadRepo.GetThreads()
}
