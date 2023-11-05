package service

import (
	"context"
	"errors"
	"github.com/TangTangHC/basic-go-study/webook/internal/domain"
	"github.com/TangTangHC/basic-go-study/webook/internal/repository"
	repmock "github.com/TangTangHC/basic-go-study/webook/internal/repository/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"golang.org/x/crypto/bcrypt"
	"testing"
	"time"
)

func Test_userService_Login(t *testing.T) {
	now := time.Now()
	testCase := []struct {
		name     string
		mockFunc func(*gomock.Controller) repository.UserRepository

		email    string
		password string

		wantErr error
		wantVal domain.User
	}{
		{
			name: "登录成功",
			mockFunc: func(ctl *gomock.Controller) repository.UserRepository {
				userRepository := repmock.NewMockUserRepository(ctl)
				userRepository.EXPECT().FindByEmail(gomock.Any(), "123@qq.com").
					Return(domain.User{
						Email:    "123@qq.com",
						Password: "$2a$10$1JSB4qHuDsgFv0/zPrxjvevGHo.FaPNmJqFKnbge0YLevcva13HyG",
						Phone:    "12345678900",
						Ctime:    now,
					}, nil)
				return userRepository
			},
			email:    "123@qq.com",
			password: "hello#world123",
			wantErr:  nil,
			wantVal: domain.User{
				Email:    "123@qq.com",
				Password: "$2a$10$1JSB4qHuDsgFv0/zPrxjvevGHo.FaPNmJqFKnbge0YLevcva13HyG",
				Phone:    "12345678900",
				Ctime:    now,
			},
		},
		{
			name: "数据库错误",
			mockFunc: func(ctl *gomock.Controller) repository.UserRepository {
				userRepository := repmock.NewMockUserRepository(ctl)
				userRepository.EXPECT().FindByEmail(gomock.Any(), "123@qq.com").
					Return(domain.User{}, errors.New("数据库错误"))
				return userRepository
			},
			email:    "123@qq.com",
			password: "hello#world123",
			wantErr:  errors.New("数据库错误"),
			wantVal:  domain.User{},
		},
		{
			name: "没找到用户",
			mockFunc: func(ctl *gomock.Controller) repository.UserRepository {
				userRepository := repmock.NewMockUserRepository(ctl)
				userRepository.EXPECT().FindByEmail(gomock.Any(), "123@qq.com").
					Return(domain.User{}, repository.ErrUserNotFound)
				return userRepository
			},
			email:    "123@qq.com",
			password: "hello#world123",
			wantErr:  ErrEmailNotSignup,
			wantVal:  domain.User{},
		},
		{
			name: "密码校验失败",
			mockFunc: func(ctl *gomock.Controller) repository.UserRepository {
				userRepository := repmock.NewMockUserRepository(ctl)
				userRepository.EXPECT().FindByEmail(gomock.Any(), "123@qq.com").
					Return(domain.User{
						Email:    "123@qq.com",
						Password: "$2a$10$1JSB4qHuDsgFv0/zPrxjvevGHo.FaPNmJqFKnbge0YLevcva13HyG",
						Phone:    "12345678900",
						Ctime:    now,
					}, nil)
				return userRepository
			},
			email:    "123@qq.com",
			password: "1hello#world123",
			wantErr:  ErrInvalidUserOrEmail,
			wantVal:  domain.User{},
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			service := NewUserService(tc.mockFunc(ctl))
			user, err := service.Login(context.Background(), tc.email, tc.password)
			assert.Equal(t, tc.wantErr, err)
			assert.Equal(t, tc.wantVal, user)
		})
	}
}

func Test_GetPassword(t *testing.T) {
	password, err := bcrypt.GenerateFromPassword([]byte("hello#world123"), bcrypt.DefaultCost)
	if err == nil {
		t.Log(string(password))
	}
}
