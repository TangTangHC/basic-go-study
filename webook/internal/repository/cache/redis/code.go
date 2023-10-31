package redis

import (
	"context"
	_ "embed"
	"errors"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var (
	ErrCodeSendTooMany        = errors.New("发送验证码太频繁")
	ErrCodeVerifyTooManyTimes = errors.New("验证次数太多")
	ErrUnknownForCode         = errors.New("我也不知发生什么了，反正是跟 code 有关")
)

//go:embed lua/set_code.lua
var luaSetCode string

//go:embed lua/verify_code.lua
var luaVerifyCode string

type CodeCache interface {
	Set(ctx context.Context, biz, phone, code string) error
	Verify(ctx context.Context, biz, phone, code string) (bool, error)
}

type RedisCodeCache struct {
	cli redis.Cmdable
}

func NewCodeCache(cli redis.Cmdable) CodeCache {
	return &RedisCodeCache{
		cli: cli,
	}
}

func (c *RedisCodeCache) Set(ctx context.Context, biz, phone, code string) error {
	i, err := c.cli.Eval(ctx, luaSetCode, []string{c.key(biz, phone)}, code).Int()
	if err != nil {
		return err
	}
	switch i {
	case 0:
		return nil
	case -1:
		return ErrCodeSendTooMany
	default:
		return errors.New("系统错误")
	}
}

func (c *RedisCodeCache) Verify(ctx context.Context, biz, phone, code string) (bool, error) {
	i, err := c.cli.Eval(ctx, luaVerifyCode, []string{c.key(biz, phone)}, code).Int()
	if err != nil {
		return false, err
	}
	switch i {
	case 0:
		return true, nil
	case -1:
		return false, ErrCodeVerifyTooManyTimes
	case -2:
		return false, nil
	}
	return false, ErrUnknownForCode
}

func (c *RedisCodeCache) key(biz, phone string) string {
	return fmt.Sprintf("phone_code:%s:%s", biz, phone)
}
