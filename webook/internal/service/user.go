package service

import (
	"context"
	"errors"
	"github.com/TangTangHC/basic-go-study/webook/internal/domain"
	"github.com/TangTangHC/basic-go-study/webook/internal/repository"
	"github.com/TangTangHC/basic-go-study/webook/internal/repository/dao"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUserDuplicateEmail = repository.ErrUserDuplicate
	ErrEmailNotSignup     = errors.New("邮箱未注册")
	ErrInvalidUserOrEmail = errors.New("邮箱或密码错误")
)

type UserService interface {
	Login(ctx context.Context, email, password string) (domain.User, error)
	SignUp(ctx context.Context, u domain.User) error
	Profile(ctx context.Context, id int64) (domain.User, error)
	Edit(ctx context.Context, edit domain.User) error
	FindOrCreate(ctx context.Context, phone string) (domain.User, error)
	//FindOrCreate(ctx context.Context, phone string) (domain.User, error)
	//FindOrCreateByWechat(ctx context.Context, wechatInfo domain.WechatInfo) (domain.User, error)
}

type userService struct {
	uRepo repository.UserRepository
}

func NewUserService(uRepo repository.UserRepository) UserService {
	return &userService{uRepo: uRepo}
}

func (u *userService) SignUp(ctx context.Context, user domain.User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hash)
	return u.uRepo.Create(ctx, user)
}

func (u *userService) Login(ctx context.Context, email string, password string) (domain.User, error) {
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

func (u *userService) Edit(ctx context.Context, edit domain.User) error {
	return u.uRepo.UpdateById(ctx, edit)
}

func (u *userService) Profile(ctx context.Context, id int64) (user domain.User, err error) {
	return u.uRepo.FindById(ctx, id)
}

func (u *userService) FindOrCreate(ctx context.Context, phone string) (domain.User, error) {
	user, err := u.uRepo.FindByPhone(ctx, phone)
	if err != dao.ErrUserNotFound {
		return user, err
	}

	user = domain.User{
		Phone: phone,
	}
	err = u.uRepo.Create(ctx, user)
	if err != nil && err != repository.ErrUserDuplicate {
		return user, err
	}
	return u.uRepo.FindByPhone(ctx, phone)
}
