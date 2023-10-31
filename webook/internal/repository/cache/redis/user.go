package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/TangTangHC/basic-go-study/webook/internal/domain"
	"github.com/redis/go-redis/v9"
	"time"
)

var ErrKeyNotExist = redis.Nil

type RedisUserCache struct {
	cli redis.Cmdable
	exp time.Duration
}

func NewRedisUserCache(cli redis.Cmdable) *RedisUserCache {
	return &RedisUserCache{
		cli: cli,
		exp: time.Minute * 15,
	}
}

func (cache *RedisUserCache) Set(ctx context.Context, user domain.User) error {
	marshal, err := json.Marshal(user)
	if err != nil {
		return err
	}
	err = cache.cli.Set(ctx, cache.key(user.Id), marshal, cache.exp).Err()
	return err
}

func (cache *RedisUserCache) Get(ctx context.Context, id int64) (domain.User, error) {
	b, err := cache.cli.Get(ctx, cache.key(id)).Bytes()
	if err != nil {
		return domain.User{}, ErrKeyNotExist
	}
	var user domain.User
	err = json.Unmarshal(b, &user)
	return user, err
}

func (cache *RedisUserCache) key(id int64) string {
	return fmt.Sprintf("user:info:%d", id)
}
