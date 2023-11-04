package web

import (
	"bytes"
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
		}}

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

	/*ctl := gomock.NewController(t)
		defer ctl.Finish()
		userSvc := svcmocks.NewMockUserService(ctl)
		userSvc.EXPECT().SignUp(gomock.Any(), gomock.Any()).Return(nil)
		handler := NewUserHandler(userSvc, nil)
		server := gin.Default()
		handler.RegisterHandler(server)

		req, err := http.NewRequest(http.MethodPost, "/users/signup", bytes.NewBufferString(`
	{
		"email": "123@qq.com",
		"password": "hello#world123",
		"confirmPassword": "hello#world123"
	}
	`))
		require.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")
		recorder := httptest.NewRecorder()
		server.ServeHTTP(recorder, req)

		assert.Equal(t, 200, recorder.Code)
		assert.Equal(t, "注册成功:邮箱123@qq.com", recorder.Body.String())*/

}

func TestMock(t *testing.T) {
	//ctl := gomock.NewController(t)
}
