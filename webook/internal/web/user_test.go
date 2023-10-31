package web

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserHandler_SignUp(t *testing.T) {
	testCase := []struct {
		name string
	}{
		{},
	}

	req, err := http.NewRequest(http.MethodPost, "/users/signup", bytes.NewBuffer([]byte(`
{"email":"123@qq.com", "password": "123456@"}
`)))
	require.NoError(t, err)
	resp := httptest.NewRecorder()
	t.Log(resp, req)
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {

		})
	}
}

func TestMock(t *testing.T) {
	//ctl := gomock.NewController(t)
}
