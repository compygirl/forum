package models

import (
	"database/sql"
	"time"
)

type User struct {
	UserID     int
	FirstName  string
	SecondName string
	Username   string
	Email      string
	Password   string
}

type Session struct {
	UserID  int
	Token   string
	ExpTime time.Time
}

type Post struct {
	PostID            int
	UserID            int
	Username          string
	Title             string
	Content           string
	CreatedTime       time.Time
	CreatedTimeString string
	LikesCounter      int
	DislikeCounter    int
	Categories        []string
	Comments          []*Comment
}

type Comment struct {
	CommentID         int
	PostID            int
	UserID            int
	Username          string
	Content           string
	CreatedTime       time.Time
	CreatedTimeString string
	LikesCounter      int
	DislikeCounter    int
}

type Database struct {
	DB *sql.DB
}
