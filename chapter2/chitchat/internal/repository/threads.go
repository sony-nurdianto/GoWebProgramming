package repository

import (
	"context"
	"database/sql"

	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/database"
)

type ThreadRepository struct {
	db *database.Database
}

func NewThreadRepository(data *database.Database) *ThreadRepository {
	return &ThreadRepository{db: data}
}

func (r *ThreadRepository) GetThreads(ctx context.Context) (rows *sql.Rows, err error) {
	rows, err = r.db.Query(ctx, `
        SELECT 
            t.id, t.uuid, t.topic, t.created_at,
            u.id, u.name, u.email,
            COUNT(p.id) AS num_replies
        FROM threads t
        JOIN users u ON t.user_id = u.id
        LEFT JOIN posts p ON p.thread_id = t.id
        GROUP BY t.id, u.id, p.id
        ORDER BY t.created_at DESC
    `)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func (r *ThreadRepository) CreateThread(ctx context.Context) (*sql.Stmt, error) {
	statement := "insert into threads (uuid, topic, user_id, created_at) values ($1, $2, $3, $4) returning id, uuid, topic, user_id, created_at"
	stmt, err := r.db.Prepare(ctx, statement)
	if err != nil {
		return nil, err
	}

	return stmt, nil
}
