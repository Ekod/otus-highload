package services

import (
	"context"

	"github.com/Ekod/otus-highload/internal/dto"
)

type Services struct {
	UserService   UserService
	FriendService FriendService
	PostService   PostService
}

func New(
	userService UserService,
	friendService FriendService,
	postService PostService,
) *Services {
	return &Services{
		UserService:   userService,
		FriendService: friendService,
		PostService:   postService,
	}
}

type UserService interface {
	LoginUser(ctx context.Context, user *dto.SecurityUser) (*dto.UserResponse, error)
	RegisterUser(ctx context.Context, user *dto.UserRequest) (*dto.UserResponse, error)
	GetUsers(ctx context.Context, id int) ([]dto.UserResponse, error)
	GetCurrentUser(ctx context.Context, userId int) (*dto.UserResponse, error)
	GetUsersByFullName(ctx context.Context, firstName, lastName string) ([]dto.UserResponse, error)
}

type FriendService interface {
	RemoveFriend(ctx context.Context, userId int, friendId int) error
	MakeFriends(ctx context.Context, userId int, friendID int) (int, error)
	GetFriends(ctx context.Context, userId int) ([]dto.UserResponse, error)
}

type PostService interface {
	CreatePost(ctx context.Context, postContent string, userID int) (int, error)
	UpdatePost(ctx context.Context, postContent string, postID int) error
	DeletePost(ctx context.Context, postID int) error
	GetPost(ctx context.Context, postID int) (dto.PostGetResponse, error)
	FeedPost(ctx context.Context, friendIDs []int) (dto.PostFeedResponse, error)
}
