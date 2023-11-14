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

func TestRegister(t *testing.T) {
	log, _ := logger.CreateMock()
	defer log.Sync()

	jwtGen := jwtgenerator.Create("secret_key", 3, log)

	selectedUser := data.User{
		Login:    "u1",
		Password: "p1",
		ID:       1,
		Role:     data.RoleUser,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mocks.NewMockBaseUsersStorage(ctrl)
	gomock.InOrder(
		mockStorage.EXPECT().SelectUserByLoginOrName(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, nil),
		mockStorage.EXPECT().InsertUser(gomock.Any(), gomock.Any()).Return(nil),
		mockStorage.EXPECT().SelectUserByLoginOrName(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, nil),
		mockStorage.EXPECT().SelectUserByLoginOrName(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("failed to find user(db error)")),
		mockStorage.EXPECT().SelectUserByLoginOrName(gomock.Any(), gomock.Any(), gomock.Any()).Return(&selectedUser, nil),
		mockStorage.EXPECT().SelectUserByLoginOrName(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, nil),
		mockStorage.EXPECT().InsertUser(gomock.Any(), gomock.Any()).Return(fmt.Errorf("failed to add user(db error)")),
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
						"Login":"u2",
						"Password":"p1",
						"Name": "user_name"
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
						"Login":"u2"
						"Password":"p1",
						"Name": "user_name"
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
						"Login":"u2",
						"Password":"p1",
						"name": "user_name"
					}`),
			},
			want: want{
				statusCode:          http.StatusBadRequest,
				authorizationHeader: false,
			},
		},
		{
			name: "db returns error",
			args: args{
				body: []byte(`{
						"Login":"u2",
						"Password":"p1",
						"Name": "user_name"
					}`),
			},
			want: want{
				statusCode:          http.StatusInternalServerError,
				authorizationHeader: false,
			},
		},
		{
			name: "user login already exists in db",
			args: args{
				body: []byte(`{
						"Login":"u2",
						"Password":"p1",
						"Name": "user_name"
					}`),
			},
			want: want{
				statusCode:          http.StatusConflict,
				authorizationHeader: false,
			},
		},
		{
			name: "error on user add in db",
			args: args{
				body: []byte(`{
						"Login":"u2",
						"Password":"p1",
						"Name": "user_name"
					}`),
			},
			want: want{
				statusCode:          http.StatusInternalServerError,
				authorizationHeader: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testApp := fiber.New()
			testApp.Delete("/", Register(mockStorage, jwtGen, log))
			defer utils.ExecuteWithLogError(testApp.Shutdown, log)

			request, errReq := http.NewRequest(http.MethodDelete, "/", bytes.NewBuffer(tt.args.body))
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
