package post_service

import (
	"context"

	"github.com/Ekod/otus-highload/domain"
	"github.com/Ekod/otus-highload/internal/dto"
)

type PostRepository interface {
	CreatePost(ctx context.Context, postContent string, userID int) (int, error)
	UpdatePost(ctx context.Context, postContent string, postID int) error
	DeletePost(ctx context.Context, postID int) error
	GetPost(ctx context.Context, postID int) (domain.Post, error)
	FeedPost(ctx context.Context, friendIDs []int) ([]domain.Post, error)
}

type Service struct {
	postRepository PostRepository
}

func New(PostRepository PostRepository) *Service {
	return &Service{
		postRepository: PostRepository,
	}
}

func (s *Service) CreatePost(
	ctx context.Context,
	postContent string,
	userID int,
) (int, error) {
	id, err := s.postRepository.CreatePost(ctx, postContent, userID)
	if err != nil {
		return 0, err
	}

	return id, nil
}
func (s *Service) UpdatePost(
	ctx context.Context,
	postContent string,
	postID int,
) error {
	err := s.postRepository.UpdatePost(ctx, postContent, postID)
	if err != nil {
		return err
	}

	return nil
}
func (s *Service) DeletePost(ctx context.Context, postID int) error {
	err := s.postRepository.DeletePost(ctx, postID)
	if err != nil {
		return err
	}

	return nil
}
func (s *Service) GetPost(ctx context.Context, postID int) (dto.PostGetResponse, error) {
	post, err := s.postRepository.GetPost(ctx, postID)
	if err != nil {
		return dto.PostGetResponse{}, err
	}

	response := dto.PostGetResponse{
		ID:     post.ID,
		Post:   post.Content,
		UserID: post.UserID,
	}

	return response, nil
}
func (s *Service) FeedPost(ctx context.Context, friendIDs []int) (dto.PostFeedResponse, error) {
	feed, err := s.postRepository.FeedPost(ctx, friendIDs)
	if err != nil {
		return dto.PostFeedResponse{}, err
	}

	postFeed := make(dto.PostFeed)

	for _, post := range feed {
		_, ok := postFeed[post.UserID]
		if ok {
			postFeed[post.UserID] = append(postFeed[post.UserID], post.Content)
		} else {
			postFeed[post.UserID] = []string{post.Content}
		}
	}

	response := dto.PostFeedResponse{
		Posts: postFeed,
	}

	return response, nil
}
