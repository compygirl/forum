package service

import (
	"forum/internal/database"
	"forum/internal/models"
	"net/http"
	"time"
)

type UserServiceInterface interface {
	CreateUser(*models.User) (int, int, error)
	Login(string, string) (*models.Session, error)
	IsUserLoggedIn(*http.Request) bool
	Logout(string) error
	IsTokenExist(string) bool
	GetUserByUserID(int) (*models.User, error)
	GetSession(string) (*models.Session, error)
	ExtendSessionTimeout(string) (time.Time, error)
}

type PostServiceInterface interface {
	GetAllPosts() ([]*models.Post, error)
	GetPostByID(int) (*models.Post, error)
	CreatePost(*models.Post) (int, int, error)
	GetCategories(int) ([]string, error)
	UpdateReaction(int, int, int) error
	Filter(string, int) ([]*models.Post, error)
}

type CommentServiceInterface interface {
	CreateComment(*models.Comment) (int, int, error)
	GetAlCommentsForPost(int) ([]*models.Comment, error)
	UpdateReaction(int, int, int) error
}

type Service struct {
	UserServiceInterface // interface
	PostServiceInterface
	CommentServiceInterface
}

func NewService(repo *database.Repository) *Service {
	serviceObj := Service{
		UserServiceInterface:    CreateNewUserService(repo.UserRepoInterface),
		PostServiceInterface:    CreateNewPostService(repo.PostRepoInterface),
		CommentServiceInterface: CreateNewCommentService(repo.CommentRepoInterface),
	}
	return &serviceObj
}
