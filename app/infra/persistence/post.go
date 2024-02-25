package persistence

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/wasuwa/blog-api/app/domain"
)

type postPersistence struct {
	conn *redis.Client
}

// NewPostPersistence postPersistenceを初期化します
func NewPostPersistence(conn *redis.Client) domain.PostRepository {
	return &postPersistence{conn: conn}
}

// FindByID IDに紐づく投稿を取得します
func (pp *postPersistence) FindByID(ctx context.Context, id int64) (*domain.Post, error) {
	cmd := pp.conn.JSONGet(ctx, "", "")
	err := cmd.Err()
	if err != nil {
		return nil, err
	}
	cmd.Val()
	return nil, nil
}
