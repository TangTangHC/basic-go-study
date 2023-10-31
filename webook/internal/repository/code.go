package repository

import (
	"context"

	"github.com/TangTangHC/basic-go-study/webook/internal/repository/cache/redis"
)

type CodeRepository struct {
	cache *redis.CodeCache
}

func NewCodeRepository(cache *redis.CodeCache) *CodeRepository {
	return &CodeRepository{
		cache: cache,
	}
}

func (r *CodeRepository) Store(ctx context.Context, biz, phone, code string) error {
	return r.cache.Set(ctx, biz, phone, code)
}

func (r *CodeRepository) Verify(ctx context.Context, biz, phone, code string) (bool, error) {
	return r.cache.Verify(ctx, biz, phone, code)
}
