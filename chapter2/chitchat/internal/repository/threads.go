package repository

import (
	"context"
	"log"
	"time"

	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/database"
	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/models"
)

type ThreadRepository struct {
	db *database.Database
}

func NewThreadRepository(data *database.Database) *ThreadRepository {
	return &ThreadRepository{db: data}
}

func (r *ThreadRepository) GetThreads() (threads []models.Thread, err error) {
	// Contoh penggunaan query
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := r.db.Query(ctx, "SELECT id, uuid, topic, user_id, created_at FROM threads ORDER BY created_at DESC")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		conv := models.Thread{}
		if err = rows.Scan(&conv.Id, &conv.Uuid, &conv.Topic, &conv.UserId, &conv.CreatedAt); err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
		threads = append(threads, conv)
	}

	if err = rows.Err(); err != nil {
		log.Println("Error iterating rows:", err)
		return nil, err
	}

	return threads, nil
}
