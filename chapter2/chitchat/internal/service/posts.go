package service

import (
	"context"
	"time"

	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/genrators"
	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/models"
	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/repository"
)

type PostService struct {
	repo *repository.PostRepository
}

func NewPostService(repo *repository.PostRepository) *PostService {
	return &PostService{
		repo: repo,
	}
}

func (ps *PostService) PostComment(body string, threadId int, userId int) (models.PostComment, error) {
	var post models.PostComment

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt, err := ps.repo.CreatePost(ctx)
	if err != nil {
		return post, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(genrators.CreateUUID(), body, userId, threadId, time.Now())
	if err := row.Scan(&post.Id, &post.Uuid, &post.Body, &post.UserId, &post.ThreadId, &post.CreatedAt); err != nil {
		return post, err
	}

	return post, nil
}
