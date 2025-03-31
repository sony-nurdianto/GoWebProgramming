package service

import (
	"context"
	"log"
	"time"

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
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := t.threadRepo.GetThreads(ctx)
	if err != nil {
		return threads, err
	}

	defer rows.Close()

	for rows.Next() {
		var thread models.Thread
		var post models.Post

		if err = rows.Scan(&thread.Id, &thread.Uuid, &thread.Topic, &thread.CreatedAt,
			&thread.User.Id, &thread.User.Name, &thread.User.Email,
			&thread.NumReplies,
			&post.Id, &post.Uuid, &post.Body, &post.UserId, &post.ThreadId, &post.CreatedAt); err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}

		thread.CreatedAtDate = thread.CreatedAt.Format("Jan 2, 2006 at 3:04pm")

		if post.Id != 0 {
			post.CreatedAtDate = post.CreatedAt.Format("Jan 2, 2006 at 3:04pm")
			thread.Posts = append(thread.Posts, post)
		}

		threads = append(threads, thread)
	}

	if err = rows.Err(); err != nil {
		log.Println("Error iterating rows:", err)
		return nil, err
	}

	return threads, nil
}
