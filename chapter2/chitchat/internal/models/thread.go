package models

import "time"

type Thread struct {
	Id            int
	Uuid          string
	Topic         string
	User          User
	Posts         []Post
	NumReplies    int
	CreatedAt     time.Time
	CreatedAtDate string
}

type PostThread struct {
	Id        int
	Uuid      string
	Topic     string
	UserId    int
	CreatedAt time.Time
}

type ThreadDetails struct {
	Thread
	Posts []Post
}
