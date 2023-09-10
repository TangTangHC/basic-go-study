package local

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRedisCodeCache_Set(t *testing.T) {
	testCases := []struct {
		name string
		// 输入
		ctx   context.Context
		biz   string
		phone string
		code  string
		count int
		// 输出
		wantErr error
	}{
		{
			name:    "验证码设置成功151",
			ctx:     context.Background(),
			biz:     "login",
			phone:   "151",
			code:    "123456",
			wantErr: nil,
		},
		{
			name:    "验证码设置成功152",
			ctx:     context.Background(),
			biz:     "login",
			phone:   "152",
			code:    "123456",
			wantErr: nil,
		},
		{
			name:    "验证码设置失败152",
			ctx:     context.Background(),
			biz:     "login",
			phone:   "152",
			code:    "123457",
			wantErr: nil,
		},
	}

	c := NewLocalCodeCache(1)
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := c.Set(tc.ctx, tc.biz, tc.phone, tc.code)
			load, ok := c.hash.Load(tc.biz + ":" + tc.phone)
			if !ok {
				assert.False(t, ok, "没放进去")
			}
			v := load.(*value)
			if err != nil {
				assert.Equal(t, tc.wantErr, err)
			} else {
				assert.Equal(t, v.code, tc.code)
			}
			// 测试定时任务删除过期的key，放置key修改为 endTime:    now.Add(600 * time.Millisecond),
			//time.Sleep(1 * time.Second)
		})
	}
}

func TestRedisCodeCache_Verify_Success(t *testing.T) {
	testCases := []struct {
		name string
		// 输入
		ctx      context.Context
		biz      string
		phone    string
		code     string
		wantCode string
		// 输出
		wantOk  bool
		wantErr error
	}{
		{
			name:     "验证码校验成功",
			ctx:      context.Background(),
			biz:      "login",
			phone:    "152",
			code:     "123456",
			wantCode: "123456",
			wantOk:   true,
			wantErr:  nil,
		},
		{
			name:     "验证码校验失败",
			ctx:      context.Background(),
			biz:      "login",
			phone:    "152",
			code:     "123456",
			wantCode: "123457",
			wantOk:   false,
			wantErr:  nil,
		},
	}

	c := NewLocalCodeCache(1)
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := c.Set(context.Background(), tc.biz, tc.phone, tc.code)
			if err != nil {
				assert.Error(t, err)
			}
			verify, err := c.Verify(tc.ctx, tc.biz, tc.phone, tc.wantCode)
			assert.Equal(t, tc.wantOk, verify)
			assert.Equal(t, tc.wantErr, err)

		})
	}
}
