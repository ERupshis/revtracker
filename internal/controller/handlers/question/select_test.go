package question

import (
	"database/sql"
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

var (
	testStr = "some_str"
)

func TestSelect(t *testing.T) {
	testLog, _ := logger.CreateMock()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	question := &data.Question{
		ID:   1,
		Name: "q1",
		Content: data.Content{
			Task:     &testStr,
			Answer:   &testStr,
			Solution: &testStr,
		},
	}

	mockStorage := mocks.NewMockBaseStorage(ctrl)
	gomock.InOrder(
		mockStorage.EXPECT().SelectQuestionByID(gomock.Any(), gomock.Any()).Return(question, nil),
		mockStorage.EXPECT().SelectQuestionByID(gomock.Any(), gomock.Any()).Return(nil, sql.ErrNoRows),
		mockStorage.EXPECT().SelectQuestionByID(gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("test err")),
		mockStorage.EXPECT().SelectQuestions(gomock.Any()).Return([]data.Question{*question}, nil),
	)

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
				paramURI: "/1",
			},
			want: want{
				statusCode: fiber.StatusOK,
				body:       []byte(`{"Id":1,"Name":"q1","Content":{"Task":"some_str","Answer":"some_str","Solution":"some_str"}}`),
			},
		},
		{
			name: "wrong id type",
			args: args{
				storage:  nil,
				log:      testLog,
				paramURI: "/asd1",
			},
			want: want{
				statusCode: fiber.StatusBadRequest,
				body:       []byte(""),
			},
		},
		{
			name: "no errors in result from db",
			args: args{
				storage:  nil,
				log:      testLog,
				paramURI: "/1",
			},
			want: want{
				statusCode: fiber.StatusNoContent,
				body:       []byte(""),
			},
		},
		{
			name: "error from db",
			args: args{
				storage:  nil,
				log:      testLog,
				paramURI: "/1",
			},
			want: want{
				statusCode: fiber.StatusInternalServerError,
				body:       []byte(""),
			},
		},
		{
			name: "valid select all",
			args: args{
				storage:  nil,
				log:      testLog,
				paramURI: "/",
			},
			want: want{
				statusCode: fiber.StatusOK,
				body:       []byte(`[{"Id":1,"Name":"q1","Content":{"Task":"some_str","Answer":"some_str","Solution":"some_str"}}]`),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testApp := fiber.New()
			testApp.Get("/:ID", Select(mockStorage, testLog))
			testApp.Get("/", Select(mockStorage, testLog))
			defer utils.ExecuteWithLogError(testApp.Shutdown, testLog)

			request, errReq := http.NewRequest(http.MethodGet, tt.args.paramURI, nil)
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
