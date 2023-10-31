package repository

import (
	"context"
	"database/sql"
	"github.com/TangTangHC/basic-go-study/webook/internal/domain"
	"github.com/TangTangHC/basic-go-study/webook/internal/repository/cache/redis"
	"github.com/TangTangHC/basic-go-study/webook/internal/repository/dao"
	"time"
)

var (
	ErrUserDuplicate = dao.ErrUserDuplicate
	ErrUserNotFound  = dao.ErrUserNotFound
)

type UserRepository interface {
	Create(ctx context.Context, user domain.User) error
	FindByEmail(ctx context.Context, email string) (domain.User, error)
	UpdateById(ctx context.Context, edit domain.User) error
	FindById(ctx context.Context, id int64) (domain.User, error)
	FindByPhone(ctx context.Context, phone string) (domain.User, error)
}

type CacheUserRepository struct {
	dao   dao.UserDao
	cache redis.UserCache
}

func NewUserRepository(dao dao.UserDao, cache redis.UserCache) UserRepository {
	return &CacheUserRepository{dao: dao, cache: cache}
}

func (u *CacheUserRepository) Create(ctx context.Context, user domain.User) error {
	return u.dao.Insert(ctx, u.domainToEntity(user))
}

func (u *CacheUserRepository) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	user, err := u.dao.FindByEmail(ctx, email)
	if err != nil {
		return domain.User{}, err
	}
	return u.entityToDomain(user), err
}

func (u *CacheUserRepository) UpdateById(ctx context.Context, edit domain.User) error {
	return u.dao.UpdateById(ctx, u.domainToEntity(edit))
}

func (u *CacheUserRepository) FindById(ctx context.Context, id int64) (domain.User, error) {
	user, err := u.cache.Get(ctx, id)
	if err == nil {
		return user, nil
	}
	//if err != redis.ErrKeyNotExist {
	//
	//}
	ue, err := u.dao.FindById(ctx, id)
	if err != nil {
		return domain.User{}, err
	}
	user = domain.User{
		Email:     ue.Email.String,
		NikeName:  ue.NikeName,
		Birthday:  ue.Birthday,
		Signature: ue.Signature,
	}
	go func() {
		err = u.cache.Set(ctx, user)
		if err != nil {
			// todo 打日志做监控
		}
	}()
	return user, nil
}

func (u *CacheUserRepository) FindByPhone(ctx context.Context, phone string) (domain.User, error) {
	user, err := u.dao.FindByPhone(ctx, phone)
	if err != nil {
		return domain.User{}, err
	}
	return u.entityToDomain(user), nil
}

func (r *CacheUserRepository) entityToDomain(u dao.User) domain.User {
	return domain.User{
		Id:       u.Id,
		Email:    u.Email.String,
		Password: u.Password,
		Phone:    u.Phone.String,
		Ctime:    time.UnixMilli(u.Ctime),
	}
}

func (r *CacheUserRepository) domainToEntity(u domain.User) dao.User {
	return dao.User{
		Id: u.Id,
		Email: sql.NullString{
			String: u.Email,
			// 我确实有手机号
			Valid: u.Email != "",
		},
		Phone: sql.NullString{
			String: u.Phone,
			Valid:  u.Phone != "",
		},
		NikeName:  u.NikeName,
		Birthday:  u.Birthday,
		Signature: u.Signature,
		Password:  u.Password,
		Ctime:     u.Ctime.UnixMilli(),
	}
}
