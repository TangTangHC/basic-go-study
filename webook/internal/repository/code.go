package repository

import (
	"context"

	"github.com/TangTangHC/basic-go-study/webook/internal/repository/cache/redis"
)

type CodeRepository interface {
	Store(ctx context.Context, biz, phone, code string) error
	Verify(ctx context.Context, biz, phone, code string) (bool, error)
}

type RedisCodeRepository struct {
	cache redis.CodeCache
}

func NewCodeRepository(cache redis.CodeCache) CodeRepository {
	return &RedisCodeRepository{
		cache: cache,
	}
}

func (r *RedisCodeRepository) Store(ctx context.Context, biz, phone, code string) error {
	return r.cache.Set(ctx, biz, phone, code)
}

func (r *RedisCodeRepository) Verify(ctx context.Context, biz, phone, code string) (bool, error) {
	return r.cache.Verify(ctx, biz, phone, code)
}
