package usecase

import (
	"context"

	"github.com/wasuwa/blog-api/app/domain"
)

// PostUseCase ユースケースの投稿インターフェース
type PostUseCase interface {
	FindByID(ctx context.Context, id int64) (*domain.Post, error)
}

type postUseCase struct {
	repository domain.PostRepository
}

// NewPostUseCase postUseCaseを初期化します
func NewPostUseCase(pr domain.PostRepository) PostUseCase {
	return &postUseCase{repository: pr}
}

// FindByID idに紐づく投稿を取得します
func (pu *postUseCase) FindByID(ctx context.Context, id int64) (*domain.Post, error) {
	return pu.repository.FindByID(ctx, id)
}
