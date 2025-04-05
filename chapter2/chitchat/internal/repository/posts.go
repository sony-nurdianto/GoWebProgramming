package repository

import (
	"context"
	"database/sql"

	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/database"
)

type PostRepository struct {
	data *database.Database
}

func NewPostRepository(data *database.Database) *PostRepository {
	return &PostRepository{
		data: data,
	}
}

func (d *PostRepository) CreatePost(ctx context.Context) (*sql.Stmt, error) {
	statement := "insert into posts (uuid, body, user_id, thread_id ,created_at) values ($1, $2, $3, $4, $5) returning id, uuid, body, user_id, thread_id ,created_at"
	stmt, err := d.data.Prepare(ctx, statement)
	if err != nil {
		return nil, err
	}

	return stmt, nil
}
