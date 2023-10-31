package cache

import (
	"context"
	"errors"
	"github.com/TangTangHC/basic-go-study/webook/internal/domain"
)

type UserCache interface {
	Set(ctx context.Context, user *domain.User) error
	Get(ctx context.Context, id int64) (*domain.User, error)
}

var (
	ErrCodeSendTooMany        = errors.New("发送验证码太频繁")
	ErrCodeVerifyTooManyTimes = errors.New("验证次数太多")
	ErrUnknownForCode         = errors.New("我也不知发生什么了，反正是跟 code 有关")
)
