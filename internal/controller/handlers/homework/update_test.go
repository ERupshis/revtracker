package homework

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/erupshis/revtracker/internal/logger"
	"github.com/erupshis/revtracker/internal/storage"
	"github.com/erupshis/revtracker/internal/utils"
	"github.com/erupshis/revtracker/mocks"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUpdate(t *testing.T) {
	testLog, _ := logger.CreateMock()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mocks.NewMockBaseStorage(ctrl)
	gomock.InOrder(
		mockStorage.EXPECT().UpdateHomework(gomock.Any(), gomock.Any()).Return(nil),
		mockStorage.EXPECT().UpdateHomework(gomock.Any(), gomock.Any()).Return(nil),
		mockStorage.EXPECT().UpdateHomework(gomock.Any(), gomock.Any()).Return(nil),
		mockStorage.EXPECT().UpdateHomework(gomock.Any(), gomock.Any()).Return(fmt.Errorf("test err")),
	)

	type args struct {
		storage  storage.BaseStorage
		log      logger.BaseLogger
		paramURI string
		body     []byte
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
				paramURI: "/1",
				body:     []byte(`{"Id":1,"Name":"hw1"}`),
			},
			want: want{
				statusCode: fiber.StatusOK,
				body:       []byte(`{"Id":1,"Name":"hw1"}`),
			},
		},
		{
			name: "missing data in body",
			args: args{
				storage:  nil,
				log:      testLog,
				paramURI: "/1",
				body:     []byte(``),
			},
			want: want{
				statusCode: fiber.StatusBadRequest,
				body:       []byte(""),
			},
		},
		{
			name: "incorrect json body",
			args: args{
				storage:  nil,
				log:      testLog,
				paramURI: "/1",
				body:     []byte(`{"Id":1"Name":"hw1"}`),
			},
			want: want{
				statusCode: fiber.StatusBadRequest,
				body:       []byte(""),
			},
		},
		{
			name: "missing ID in URI and in body",
			args: args{
				storage:  nil,
				log:      testLog,
				paramURI: "/",
				body:     []byte(`{"Name":"hw1"}`),
			},
			want: want{
				statusCode: fiber.StatusBadRequest,
				body:       []byte(""),
			},
		},
		{
			name: "ID in URI only",
			args: args{
				storage:  nil,
				log:      testLog,
				paramURI: "/1",
				body:     []byte(`{"Name":"hw1"}`),
			},
			want: want{
				statusCode: fiber.StatusOK,
				body:       []byte(`{"Id":1,"Name":"hw1"}`),
			},
		},
		{
			name: "ID in body only",
			args: args{
				storage:  nil,
				log:      testLog,
				paramURI: "/",
				body:     []byte(`{"Id":1,"Name":"hw1"}`),
			},
			want: want{
				statusCode: fiber.StatusOK,
				body:       []byte(`{"Id":1,"Name":"hw1"}`),
			},
		},
		{
			name: "incorrect json in body (missing or empty name)",
			args: args{
				storage:  nil,
				log:      testLog,
				paramURI: "/1",
				body:     []byte(`{"Id":1}`),
			},
			want: want{
				statusCode: fiber.StatusBadRequest,
				body:       []byte(""),
			},
		},
		{
			name: "error from db",
			args: args{
				storage:  nil,
				log:      testLog,
				paramURI: "/1",
				body:     []byte(`{"Id":1,"Name":"hw1"}`),
			},
			want: want{
				statusCode: fiber.StatusInternalServerError,
				body:       []byte(""),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testApp := fiber.New()
			testApp.Put("/:ID", Update(mockStorage, testLog))
			testApp.Put("/", Update(mockStorage, testLog))
			defer utils.ExecuteWithLogError(testApp.Shutdown, testLog)

			request, errReq := http.NewRequest(http.MethodPut, tt.args.paramURI, bytes.NewBuffer(tt.args.body))
			require.NoError(t, errReq)

			response, errResp := testApp.Test(request)
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
