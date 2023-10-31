package local

import (
	"context"
	"github.com/TangTangHC/basic-go-study/webook/internal/domain"
)

type LocalUserCache struct {
}

func (l LocalUserCache) Set1(ctx context.Context, user *domain.User) error {
	//TODO implement me
	panic("implement me")
}

func (l LocalUserCache) get1(ctx context.Context, id int64) (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}
