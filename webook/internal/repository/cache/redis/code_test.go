package redis

import (
	"context"
	redismock "github.com/TangTangHC/basic-go-study/webook/internal/repository/cache/redis/mock"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestRedisCodeCache_Set(t *testing.T) {
	testCase := []struct {
		name    string
		mockFun func(*gomock.Controller) redis.Cmdable

		biz   string
		phone string
		code  string

		wantErr error
	}{
		{
			name: "校验通过",
			mockFun: func(ctl *gomock.Controller) redis.Cmdable {
				cmdable := redismock.NewMockCmdable(ctl)
				cmd := redis.NewCmd(context.Background())
				cmd.SetErr(nil)
				cmd.SetVal("0")
				cmdable.EXPECT().Eval(context.Background(), luaSetCode, []string{"phone_code:login:123"}, []any{"123456"}).
					Return(cmd)
				return cmdable
			},
			biz:   "login",
			phone: "123",
			code:  "123456",
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			cache := NewCodeCache(tc.mockFun(ctl))
			err := cache.Set(context.Background(), tc.biz, tc.phone, tc.code)
			assert.Equal(t, tc.wantErr, err)
		})
	}
}
