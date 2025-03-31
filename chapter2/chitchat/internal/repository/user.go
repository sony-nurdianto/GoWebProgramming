package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/database"
	"github.com/sony-nurdianto/GoWebProgramming/chapter2/chitchat/internal/models"
)

type UserRepository struct {
	data *database.Database
}

func NewUserRepository(data *database.Database) *UserRepository {
	return &UserRepository{
		data: data,
	}
}

func (d *UserRepository) NewUser(ctx context.Context, user models.User) error {
	statement := "insert into users (uuid, name, email, password, created_at) values ($1, $2, $3, $4, $5) returning id, uuid, created_at"

	stmt, err := d.data.Prepare(ctx, statement)
	if err != nil {
		return err
	}

	defer stmt.Close()

	stmt.QueryRow(user.Uuid, user.Name, user.Email, user.Password, user.CreatedAt)

	return nil
}

func (d *UserRepository) GetAuthUser(ctx context.Context, email string) (user models.User, err error) {
	row := d.data.QueryRow(ctx, "SELECT id, uuid,name ,email,password ,created_at FROM users WHERE email = $1", email)
	if err := row.Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return user, fmt.Errorf("user with email %s is not found: %w", email, sql.ErrNoRows)
		default:
			return user, err
		}
	}

	return user, nil
}
