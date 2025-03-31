package models

import "time"

type Post struct {
	Id            int
	Uuid          string
	Body          string
	UserId        int
	ThreadId      int
	CreatedAt     time.Time
	CreatedAtDate string
}
