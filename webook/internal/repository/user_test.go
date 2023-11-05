package repository

import (
	"context"
	"database/sql"
	"github.com/TangTangHC/basic-go-study/webook/internal/domain"
	"github.com/TangTangHC/basic-go-study/webook/internal/repository/cache/redis"
	redismock "github.com/TangTangHC/basic-go-study/webook/internal/repository/cache/redis/mock"
	"github.com/TangTangHC/basic-go-study/webook/internal/repository/dao"
	daomock "github.com/TangTangHC/basic-go-study/webook/internal/repository/dao/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
	"time"
)

func TestCacheUserRepository_FindById(t *testing.T) {
	now := time.Now()
	testCase := []struct {
		name    string
		mockFuc func(*gomock.Controller) (dao.UserDao, redis.UserCache)

		id int64

		wantErr error
		wantVal domain.User
	}{
		{
			name: "通过数据库查询",
			mockFuc: func(ctl *gomock.Controller) (dao.UserDao, redis.UserCache) {
				userDao := daomock.NewMockUserDao(ctl)
				userCache := redismock.NewMockUserCache(ctl)
				userCache.EXPECT().Get(gomock.Any(), int64(123)).Return(domain.User{}, redis.ErrKeyNotExist)
				userDao.EXPECT().FindById(gomock.Any(), int64(123)).Return(dao.User{
					Id: 0,
					Email: sql.NullString{
						String: "123@qq.com",
						Valid:  true,
					},
					Phone: sql.NullString{
						String: "1234567890",
						Valid:  true,
					},
					Password:  "123pass",
					NikeName:  "nikeName",
					Birthday:  "2023-01-01",
					Signature: "signature",
					Ctime:     now.UnixMilli(),
					Utime:     now.UnixMilli(),
				}, nil)
				return userDao, userCache
			},
			id:      123,
			wantErr: nil,
			wantVal: domain.User{
				Id:        0,
				Email:     "123@qq.com",
				NikeName:  "nikeName",
				Birthday:  "2023-01-01",
				Signature: "signature",
			},
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			userDao, cache := tc.mockFuc(ctl)
			userRepository := NewUserRepository(userDao, cache)
			user, err := userRepository.FindById(context.Background(), tc.id)
			assert.Equal(t, tc.wantErr, err)
			assert.Equal(t, tc.wantVal, user)
		})
	}
}
