package web

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/TangTangHC/basic-go-study/webook/internal/domain"
	"github.com/TangTangHC/basic-go-study/webook/internal/service"
	svcmocks "github.com/TangTangHC/basic-go-study/webook/internal/service/mock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserHandler_SignUp(t *testing.T) {
	testCase := []struct {
		name string

		mockFunc func(*gomock.Controller) service.UserService
		reqBody  string

		wantCode int
		wantMsg  string
	}{
		{
			name: "成功注册",
			mockFunc: func(ctl *gomock.Controller) service.UserService {
				userService := svcmocks.NewMockUserService(ctl)
				userService.EXPECT().SignUp(gomock.Any(), domain.User{
					Email:    "123@qq.com",
					Password: "hello#world123",
				}).Return(nil)
				return userService
			},
			reqBody: `
{
	"email": "123@qq.com",
	"password": "hello#world123",
	"confirmPassword": "hello#world123"
}
`,
			wantCode: 200,
			wantMsg:  "注册成功:邮箱123@qq.com",
		},
		{
			name: "参数解析异常",
			mockFunc: func(ctl *gomock.Controller) service.UserService {
				srv := svcmocks.NewMockUserService(ctl)
				return srv
			},
			reqBody: `
{
	"Email":"123@qq.com",
}
`,
			wantCode: 400,
			wantMsg:  "",
		},
		{
			name: "邮箱错误",
			mockFunc: func(ctl *gomock.Controller) service.UserService {
				srv := svcmocks.NewMockUserService(ctl)
				return srv
			},
			reqBody: `
{
	"Email":"123.com",
	"password": "hello#world123",
	"confirmPassword": "hello#world123"
}
`,
			wantCode: 200,
			wantMsg:  "邮箱格式不正确",
		},
		{
			name: "两次密码不同，重新输入",
			mockFunc: func(ctl *gomock.Controller) service.UserService {
				srv := svcmocks.NewMockUserService(ctl)
				return srv
			},
			reqBody: `
{
	"Email":"123@qq.com",
	"password": "hello#world1231",
	"confirmPassword": "hello#world1232"
}
`,
			wantCode: 200,
			wantMsg:  "两次密码不同，重新输入",
		},
		{
			name: "密码校验失败",
			mockFunc: func(ctl *gomock.Controller) service.UserService {
				srv := svcmocks.NewMockUserService(ctl)
				return srv
			},
			reqBody: `
{
	"Email":"123@qq.com",
	"password": "helloworld123",
	"confirmPassword": "helloworld123"
}
`,
			wantCode: 200,
			wantMsg:  "密码必须大于8位，包含数字、特殊字符",
		},
		{
			name: "邮箱冲突",
			mockFunc: func(ctl *gomock.Controller) service.UserService {
				svc := svcmocks.NewMockUserService(ctl)
				svc.EXPECT().SignUp(context.Background(), domain.User{
					Email:    "123@qq.com",
					Password: "hello#world123",
				}).Return(service.ErrUserDuplicateEmail)
				return svc
			},
			reqBody: `
{
	"email": "123@qq.com",
	"password": "hello#world123",
	"confirmPassword": "hello#world123"
}
`,
			wantCode: 200,
			wantMsg:  "邮箱冲突: 123@qq.com",
		},
		{
			name: "系统异常",
			mockFunc: func(ctl *gomock.Controller) service.UserService {
				svc := svcmocks.NewMockUserService(ctl)
				svc.EXPECT().SignUp(context.Background(), domain.User{
					Email:    "123@qq.com",
					Password: "hello#world123",
				}).Return(errors.New(""))
				return svc
			},
			reqBody: `
{
	"email": "123@qq.com",
	"password": "hello#world123",
	"confirmPassword": "hello#world123"
}
`,
			wantCode: 200,
			wantMsg:  "系统错误",
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			userService := tc.mockFunc(ctl)
			handler := NewUserHandler(userService, nil)
			server := gin.Default()
			handler.RegisterHandler(server)
			request, err := http.NewRequest(http.MethodPost, "/users/signup", bytes.NewBufferString(tc.reqBody))
			require.NoError(t, err)
			response := httptest.NewRecorder()
			server.ServeHTTP(response, request)

			assert.Equal(t, tc.wantCode, response.Code)
			assert.Equal(t, tc.wantMsg, response.Body.String())
		})
	}
}

func TestMock(t *testing.T) {
	//ctl := gomock.NewController(t)
}

func TestUserHandler_LoginSMS(t *testing.T) {
	testCase := []struct {
		name     string
		mockFunc func(*gomock.Controller) (service.UserService, service.CodeService)

		req string

		wantCode int
		wantVal  Result
	}{
		{
			name: "登录成功",
			mockFunc: func(ctl *gomock.Controller) (service.UserService, service.CodeService) {
				userService := svcmocks.NewMockUserService(ctl)
				codeService := svcmocks.NewMockCodeService(ctl)
				codeService.EXPECT().Verify(gomock.Any(), "login", "1234567890", "1099").
					Return(true, nil)
				userService.EXPECT().FindOrCreate(gomock.Any(), "1234567890").Return(domain.User{
					Id: 1,
				}, nil)
				return userService, codeService
			},
			req: `
{
	"phone": "1234567890",
	"code": "1099"
}
`,
			wantCode: http.StatusOK,
			wantVal: Result{
				Msg: "验证码校验通过",
			},
		},
		{
			name: "验证码验证异常",
			mockFunc: func(ctl *gomock.Controller) (service.UserService, service.CodeService) {
				userService := svcmocks.NewMockUserService(ctl)
				codeService := svcmocks.NewMockCodeService(ctl)
				codeService.EXPECT().Verify(gomock.Any(), "login", "1234567890", "1099").
					Return(false, errors.New("code验证异常"))
				return userService, codeService
			},
			req: `
{
	"phone": "1234567890",
	"code": "1099"
}
`,
			wantCode: http.StatusOK,
			wantVal: Result{
				Code: 5,
				Msg:  "系统异常",
			},
		},
		{
			name: "验证码验证失败",
			mockFunc: func(ctl *gomock.Controller) (service.UserService, service.CodeService) {
				userService := svcmocks.NewMockUserService(ctl)
				codeService := svcmocks.NewMockCodeService(ctl)
				codeService.EXPECT().Verify(gomock.Any(), "login", "1234567890", "1099").
					Return(false, nil)
				return userService, codeService
			},
			req: `
{
	"phone": "1234567890",
	"code": "1099"
}
`,
			wantCode: http.StatusOK,
			wantVal: Result{
				Code: 4,
				Msg:  "验证码校验失败",
			},
		},
		{
			name: "新用户注册失败",
			mockFunc: func(ctl *gomock.Controller) (service.UserService, service.CodeService) {
				userService := svcmocks.NewMockUserService(ctl)
				codeService := svcmocks.NewMockCodeService(ctl)
				codeService.EXPECT().Verify(gomock.Any(), "login", "1234567890", "1099").
					Return(true, nil)
				userService.EXPECT().FindOrCreate(gomock.Any(), "1234567890").Return(domain.User{}, errors.New("新用户注册失败"))
				return userService, codeService
			},
			req: `
{
	"phone": "1234567890",
	"code": "1099"
}
`,
			wantCode: http.StatusOK,
			wantVal: Result{
				Code: 5,
				Msg:  "系统错误",
			},
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			userService, codeService := tc.mockFunc(ctl)
			handler := NewUserHandler(userService, codeService)
			server := gin.Default()
			handler.RegisterHandler(server)
			request, err := http.NewRequest(http.MethodPost, "/users/login_sms", bytes.NewBufferString(tc.req))
			require.NoError(t, err)
			res := httptest.NewRecorder()
			server.ServeHTTP(res, request)
			assert.Equal(t, tc.wantCode, res.Code)
			var r Result
			err = json.Unmarshal(res.Body.Bytes(), &r)
			require.NoError(t, err)
			assert.Equal(t, tc.wantVal, r)
		})
	}
}
