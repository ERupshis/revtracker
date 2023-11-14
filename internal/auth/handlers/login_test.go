package handlers

import (
	"bytes"
	"fmt"
	"net/http"
	"testing"

	"github.com/erupshis/revtracker/internal/auth/data"
	"github.com/erupshis/revtracker/internal/auth/jwtgenerator"
	"github.com/erupshis/revtracker/internal/logger"
	"github.com/erupshis/revtracker/internal/utils"
	"github.com/erupshis/revtracker/mocks"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLogin(t *testing.T) {
	log, _ := logger.CreateMock()
	defer log.Sync()

	jwtGen := jwtgenerator.Create("secret_key", 3, log)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	user1 := data.User{
		Login:    "u1",
		Password: "p1",
		ID:       1,
		Role:     data.RoleUser,
	}

	mockStorage := mocks.NewMockBaseUsersStorage(ctrl)
	gomock.InOrder(
		mockStorage.EXPECT().SelectUserByLogin(gomock.Any(), gomock.Any()).Return(&user1, nil),
		mockStorage.EXPECT().SelectUserByLogin(gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("test error")),
		mockStorage.EXPECT().SelectUserByLogin(gomock.Any(), gomock.Any()).Return(nil, nil),
		mockStorage.EXPECT().SelectUserByLogin(gomock.Any(), gomock.Any()).Return(&user1, nil),
	)

	type args struct {
		body []byte
	}
	type want struct {
		statusCode          int
		authorizationHeader bool
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "valid",
			args: args{
				body: []byte(`{
						"Login":"u1",
						"Password":"p1"
					}`),
			},
			want: want{
				statusCode:          http.StatusOK,
				authorizationHeader: true,
			},
		},
		{
			name: "fail unmarshalling request body",
			args: args{
				body: []byte(`{
						"Login":"u1"
						"Password":"p1"
					}`),
			},
			want: want{
				statusCode:          http.StatusBadRequest,
				authorizationHeader: false,
			},
		},
		{
			name: "fail to validate unmarshalled data",
			args: args{
				body: []byte(`{
						"Login":"u1",
						"password":"p1"
					}`),
			},
			want: want{
				statusCode:          http.StatusBadRequest,
				authorizationHeader: false,
			},
		},
		{
			name: "error from database",
			args: args{
				body: []byte(`{
						"Login":"u1",
						"Password":"p1"
					}`),
			},
			want: want{
				statusCode:          http.StatusInternalServerError,
				authorizationHeader: false,
			},
		},
		{
			name: "missing user in DB",
			args: args{
				body: []byte(`{
						"Login":"u1",
						"Password":"p1"
					}`),
			},
			want: want{
				statusCode:          http.StatusUnauthorized,
				authorizationHeader: false,
			},
		},
		{
			name: "incorrect password",
			args: args{
				body: []byte(`{
						"Login":"u1",
						"Password":"p2"
					}`),
			},
			want: want{
				statusCode:          http.StatusUnauthorized,
				authorizationHeader: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testApp := fiber.New()
			testApp.Post("/", Login(mockStorage, jwtGen, log))
			defer utils.ExecuteWithLogError(testApp.Shutdown, log)

			request, errReq := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(tt.args.body))
			require.NoError(t, errReq)

			response, errResp := testApp.Test(request)
			require.NoError(t, errResp)
			defer func() {
				_ = response.Body.Close()
			}()

			assert.Equal(t, tt.want.statusCode, response.StatusCode)
			assert.Equal(t, tt.want.authorizationHeader, response.Header.Get("Authorization") != "")
		})
	}
}
