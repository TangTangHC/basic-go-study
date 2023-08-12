package repository

import (
	"context"
	"github.com/TangTangHC/basic-go-study/webook/internal/domain"
	"github.com/TangTangHC/basic-go-study/webook/internal/repository/dao"
	"github.com/gin-gonic/gin"
	"time"
)

var (
	ErrUserDuplicateEmail = dao.ErrUserDuplicateEmail
	ErrUserNotFound       = dao.ErrUserNotFound
)

type UserRepository struct {
	dao *dao.UserDao
}

func NewUserRepository(dao *dao.UserDao) *UserRepository {
	return &UserRepository{dao: dao}
}

func (u *UserRepository) Create(ctx context.Context, user domain.User) error {
	return u.dao.Insert(ctx, dao.User{
		Email:    user.Email,
		Password: user.Password,
	})
}

func (u *UserRepository) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	user, err := u.dao.FindByEmail(ctx, email)
	if err != nil {
		return domain.User{}, err
	}
	return domain.User{
		Id:       user.Id,
		Email:    user.Email,
		Password: user.Password,
	}, err
}

func (u *UserRepository) UpdateById(ctx context.Context, edit domain.User) error {
	milli := time.Now().UnixMilli()
	return u.dao.UpdateById(ctx, dao.User{
		Id:        edit.Id,
		Email:     edit.Email,
		NikeName:  edit.NikeName,
		Birthday:  edit.Birthday,
		Signature: edit.Signature,
		Utime:     milli,
	})
}

func (u *UserRepository) FindById(ctx *gin.Context, id int64) (domain.User, error) {
	user, err := u.dao.FindById(ctx, id)
	if err != nil {
		return domain.User{}, err
	}
	return domain.User{
		Email:     user.Email,
		NikeName:  user.NikeName,
		Birthday:  user.Birthday,
		Signature: user.Signature,
	}, nil
}
