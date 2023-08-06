package service

import (
	"context"
	"github.com/TangTangHC/basic-go-study/webook/internal/domain"
	"github.com/TangTangHC/basic-go-study/webook/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUserDuplicateEmail = repository.ErrUserDuplicateEmail
)

type UserService struct {
	uRepo *repository.UserRepository
}

func NewUserService(uRepo *repository.UserRepository) *UserService {
	return &UserService{uRepo: uRepo}
}

func (u *UserService) SignUp(ctx context.Context, user domain.User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hash)
	return u.uRepo.Create(ctx, user)
}
