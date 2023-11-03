package handlers

import (
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/erupshis/revtracker/internal/data"
	"github.com/erupshis/revtracker/internal/logger"
	"github.com/erupshis/revtracker/internal/storage"
	"github.com/erupshis/revtracker/internal/utils"
	"github.com/erupshis/revtracker/mocks"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const host = "http://localhost:"

func TestAddUser(t *testing.T) {
	testLog, _ := logger.CreateTestPLug()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mocks.NewMockBaseStorage(ctrl)
	gomock.InOrder(
		mockStorage.EXPECT().SelectUser(gomock.Any(), gomock.Any()).Return(nil, nil),
		mockStorage.EXPECT().InsertUser(gomock.Any(), gomock.Any()).Return(int64(1), nil),

		mockStorage.EXPECT().SelectUser(gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("err")),

		mockStorage.EXPECT().SelectUser(gomock.Any(), gomock.Any()).Return(&data.User{}, nil),

		mockStorage.EXPECT().SelectUser(gomock.Any(), gomock.Any()).Return(nil, nil),
		mockStorage.EXPECT().InsertUser(gomock.Any(), gomock.Any()).Return(int64(-1), fmt.Errorf("err")),

		mockStorage.EXPECT().SelectUser(gomock.Any(), gomock.Any()).Return(nil, nil),
		mockStorage.EXPECT().InsertUser(gomock.Any(), gomock.Any()).Return(int64(-1), nil),
	)

	testApp := fiber.New()
	testApp.Post("/:name", AddUser(mockStorage, testLog))
	defer utils.ExecuteWithLogError(testApp.Shutdown, testLog)

	port := 3001
	go func() {
		err := testApp.Listen(":" + fmt.Sprintf("%d", port))
		if err != nil {
			panic(err)
		}
	}()

	type args struct {
		storage  storage.BaseStorage
		log      logger.BaseLogger
		paramURI string
	}
	type want struct {
		statusCode int
		body       []byte
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "valid",
			args: args{
				storage:  nil,
				log:      testLog,
				paramURI: "/any_name",
			},
			want: want{
				statusCode: fiber.StatusOK,
				body:       []byte("1"),
			},
		},
		{
			name: "without name",
			args: args{
				storage:  nil,
				log:      testLog,
				paramURI: "/",
			},
			want: want{
				statusCode: fiber.StatusNotFound,
				body:       []byte("Cannot POST /"),
			},
		},
		{
			name: "error from bd on Select",
			args: args{
				storage:  nil,
				log:      testLog,
				paramURI: "/any_name",
			},
			want: want{
				statusCode: fiber.StatusInternalServerError,
				body:       []byte(""),
			},
		},
		{
			name: "user already exists",
			args: args{
				storage:  nil,
				log:      testLog,
				paramURI: "/any_name",
			},
			want: want{
				statusCode: fiber.StatusConflict,
				body:       []byte(""),
			},
		},
		{
			name: "error from bd on Insert",
			args: args{
				storage:  nil,
				log:      testLog,
				paramURI: "/any_name",
			},
			want: want{
				statusCode: fiber.StatusInternalServerError,
				body:       []byte(""),
			},
		},
		{
			name: "negative id on Insert in bd",
			args: args{
				storage:  nil,
				log:      testLog,
				paramURI: "/any_name",
			},
			want: want{
				statusCode: fiber.StatusInternalServerError,
				body:       []byte(""),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request, errReq := http.NewRequest(http.MethodPost, host+fmt.Sprintf("%d", port)+tt.args.paramURI, nil)
			require.NoError(t, errReq)

			client := http.Client{}
			response, errResp := client.Do(request)
			require.NoError(t, errResp)
			defer func() {
				_ = response.Body.Close()
			}()

			assert.Equal(t, tt.want.statusCode, response.StatusCode)

			respBody, err := io.ReadAll(response.Body)
			require.NoError(t, err)
			assert.Equal(t, string(tt.want.body), string(respBody))
		})
	}

}
