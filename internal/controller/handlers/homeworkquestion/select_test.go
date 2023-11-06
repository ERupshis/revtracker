package homeworkquestion

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/erupshis/revtracker/internal/data"
	"github.com/erupshis/revtracker/internal/logger"
	"github.com/erupshis/revtracker/internal/storage"
	utilsReform "github.com/erupshis/revtracker/internal/storage/manager/reform/utils"
	"github.com/erupshis/revtracker/internal/utils"
	"github.com/erupshis/revtracker/mocks"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSelect(t *testing.T) {
	testLog, _ := logger.CreateMock()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	question := &data.HomeworkQuestion{
		ID:         1,
		HomeworkID: 1,
		QuestionID: 1,
		Order:      1,
	}

	mockStorage := mocks.NewMockBaseStorage(ctrl)
	gomock.InOrder(
		mockStorage.EXPECT().SelectHomeworkQuestionByID(gomock.Any(), gomock.Any()).Return(question, nil),
		mockStorage.EXPECT().SelectHomeworkQuestionByID(gomock.Any(), gomock.Any()).Return(nil, sql.ErrNoRows),
		mockStorage.EXPECT().SelectHomeworkQuestionByID(gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("test err")),
	)

	testApp := fiber.New()
	testApp.Get("/:ID", Select(mockStorage, testLog))
	defer utils.ExecuteWithLogError(testApp.Shutdown, testLog)

	port := 3032
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
				paramURI: "/1",
			},
			want: want{
				statusCode: fiber.StatusOK,
				body:       []byte(`{"Id":1,"Homework_Id":1,"Question_Id":1,"Order":1}`),
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request, errReq := http.NewRequest(http.MethodGet, utilsReform.HostTest+fmt.Sprintf("%d", port)+tt.args.paramURI, nil)
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
