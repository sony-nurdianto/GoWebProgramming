package models

import "time"

type Post struct {
	Id            int
	Uuid          string
	Body          string
	UserId        int
	UserName      string
	ThreadId      int
	ThreadUuid    string
	CreatedAt     time.Time
	CreatedAtDate string
}

type PostComment struct {
	Id        int
	Uuid      string
	Body      string
	UserId    int
	ThreadId  int
	CreatedAt time.Time
}
