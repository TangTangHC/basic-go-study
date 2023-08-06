package service

import (
	"context"
	"errors"
	"github.com/TangTangHC/basic-go-study/webook/internal/domain"
	"github.com/TangTangHC/basic-go-study/webook/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUserDuplicateEmail = repository.ErrUserDuplicateEmail
	ErrEmailNotSignup     = errors.New("邮箱未注册")
	ErrInvalidUserOrEmail = errors.New("邮箱或密码错误")
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

func (u *UserService) Login(ctx context.Context, email string, password string) (domain.User, error) {
	user, err := u.uRepo.FindByEmail(ctx, email)
	if err == repository.ErrUserNotFound {
		return domain.User{}, ErrEmailNotSignup
	}
	if err != nil {
		return domain.User{}, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return domain.User{}, ErrInvalidUserOrEmail
	}
	return user, nil
}
