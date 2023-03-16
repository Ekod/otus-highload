package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Ekod/otus-highload/domain"
	"github.com/Ekod/otus-highload/utils/errors"

	"github.com/huandu/go-sqlbuilder"
)

const (
	queryFeed = "SELECT id, user_id, content FROM posts WHERE user_id IN (?)"
)

type PostRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (p *PostRepository) CreatePost(ctx context.Context, postContent string, userID int) (int, error) {
	ib := sqlbuilder.MySQL.NewInsertBuilder()

	ib.InsertInto("post")
	ib.Cols("user_id", "content")
	ib.Values(userID, postContent)

	query, args := ib.Build()

	res, err := p.db.ExecContext(ctx, query, args)
	if err != nil {
		debugMessageError := fmt.Sprintf("[ERROR] PostRepository_CreatePost - ExecContext: %s", err)

		return 0, errors.NewInternalServerError("Server error", debugMessageError)
	}

	id, err := res.LastInsertId()
	if err != nil {
		debugMessageError := fmt.Sprintf("[ERROR] PostRepository_CreatePost - LastInsertId: %s", err)

		return 0, errors.NewInternalServerError("Server error", debugMessageError)
	}

	return int(id), nil
}
func (p *PostRepository) UpdatePost(ctx context.Context, postContent string, postID int) error {
	ub := sqlbuilder.MySQL.NewUpdateBuilder()

	ub.Update("post")
	ub.Add("content", postContent)
	ub.Where(ub.Equal("id", postID))

	query, args := ub.Build()

	_, err := p.db.ExecContext(ctx, query, args)
	if err != nil {
		debugMessageError := fmt.Sprintf("[ERROR] PostRepository_UpdatePost - ExecContext: %s", err)

		return errors.NewInternalServerError("Server error", debugMessageError)
	}

	return nil
}
func (p *PostRepository) DeletePost(ctx context.Context, postID int) error {
	db := sqlbuilder.MySQL.NewDeleteBuilder()

	db.DeleteFrom("posts")
	db.Where(db.Equal("id", postID))

	query, args := db.Build()

	_, err := p.db.ExecContext(ctx, query, args)
	if err != nil {
		debugMessageError := fmt.Sprintf("[ERROR] PostRepository_DeletePost - ExecContext: %s", err)

		return errors.NewInternalServerError("Server error", debugMessageError)
	}

	return nil
}

func (p *PostRepository) GetPost(ctx context.Context, postID int) (domain.Post, error) {
	sb := sqlbuilder.MySQL.NewSelectBuilder()

	sb.Select("id", "user_id", "content")
	sb.From("posts")
	sb.Where(sb.Equal("id", postID))

	query, args := sb.Build()

	var post domain.Post

	err := p.db.QueryRowContext(ctx, query, args).Scan(&post)
	if err != nil {
		debugMessageError := fmt.Sprintf("[ERROR] PostRepository_GetPost - QueryRowContext: %s", err)

		return domain.Post{}, errors.NewInternalServerError("Server error", debugMessageError)
	}

	return post, nil
}

func (p *PostRepository) FeedPost(ctx context.Context, friendIDs []int) ([]domain.Post, error) {
	sb := sqlbuilder.MySQL.NewSelectBuilder()

	sb.Select("id", "user_id", "content")
	sb.From("posts")
	sb.Where(sb.In("user_id", friendIDs))

	query, args := sb.Build()

	rows, err := p.db.QueryContext(ctx, query, args...)
	if err != nil {
		debugMessageError := fmt.Sprintf("[ERROR] PostRepository_FeedPost - QueryContext: %s", err)

		return nil, errors.NewInternalServerError("Server error", debugMessageError)
	}
	defer rows.Close()

	posts := make([]domain.Post, 0)

	for rows.Next() {
		var post domain.Post

		err = rows.Scan(&post.ID, &post.UserID, &post.Content)
		if err != nil {
			continue
		}

		posts = append(posts, post)
	}

	return posts, nil
}
