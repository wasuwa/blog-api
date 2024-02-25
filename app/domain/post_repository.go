package domain

import "context"

// PostRepository 投稿のリポジトリ
type PostRepository interface {
	FindByID(ctx context.Context, id int64) (*Post, error)
}
