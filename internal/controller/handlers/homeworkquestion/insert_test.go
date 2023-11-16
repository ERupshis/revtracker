package homeworkquestion

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"testing"

	utilsHandlers "github.com/erupshis/revtracker/internal/controller/handlers/utils"
	"github.com/erupshis/revtracker/internal/logger"
	"github.com/erupshis/revtracker/internal/storage"
	"github.com/erupshis/revtracker/internal/utils"
	"github.com/erupshis/revtracker/mocks"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInsert(t *testing.T) {
	testLog, _ := logger.CreateMock()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mocks.NewMockBaseStorage(ctrl)
	gomock.InOrder(
		mockStorage.EXPECT().InsertHomeworkQuestion(gomock.Any(), gomock.Any()).Return(nil),
		mockStorage.EXPECT().InsertHomeworkQuestion(gomock.Any(), gomock.Any()).Return(fmt.Errorf("test err")),
		mockStorage.EXPECT().InsertHomeworkQuestion(gomock.Any(), gomock.Any()).Return(utilsHandlers.ErrQuestionAlreadyInHomework),
		mockStorage.EXPECT().InsertHomeworkQuestion(gomock.Any(), gomock.Any()).Return(utilsHandlers.ErrQuestionNotFound),
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
				paramURI: "",
				body:     []byte(`{"Id":1,"Homework_Id":1,"Question_Id":1,"Order":1}`),
			},
			want: want{
				statusCode: fiber.StatusOK,
				body:       []byte("Id: 0"),
			},
		},
		{
			name: "incorrect json in body",
			args: args{
				storage:  nil,
				log:      testLog,
				paramURI: "",
				body:     []byte(`{"Id":1"Homework_Id":1,"Question_Id":1,"Order":1}`),
			},
			want: want{
				statusCode: fiber.StatusBadRequest,
				body:       []byte(""),
			},
		},
		{
			name: "missing data in body",
			args: args{
				storage:  nil,
				log:      testLog,
				paramURI: "",
				body:     []byte(``),
			},
			want: want{
				statusCode: fiber.StatusBadRequest,
				body:       []byte(""),
			},
		},
		{
			name: "incorrect json in body (missing or empty Homework_Id)",
			args: args{
				storage:  nil,
				log:      testLog,
				paramURI: "",
				body:     []byte(`{"Id":1,"Question_Id":1,"Order":1}`),
			},
			want: want{
				statusCode: fiber.StatusBadRequest,
				body:       []byte(""),
			},
		},
		{
			name: "incorrect json in body (missing or empty Question_Id)",
			args: args{
				storage:  nil,
				log:      testLog,
				paramURI: "",
				body:     []byte(`{"Id":1,"Homework_Id":1,"Order":1}`),
			},
			want: want{
				statusCode: fiber.StatusBadRequest,
				body:       []byte(""),
			},
		},
		{
			name: "incorrect json in body (missing or empty Order)",
			args: args{
				storage:  nil,
				log:      testLog,
				paramURI: "",
				body:     []byte(`{"Id":1,"Homework_Id":1,"Question_Id":1}`),
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
				paramURI: "",
				body:     []byte(`{"Id":1,"Homework_Id":1,"Question_Id":1,"Order":1}`),
			},
			want: want{
				statusCode: fiber.StatusInternalServerError,
				body:       []byte(""),
			},
		},
		{
			name: "question already in homework",
			args: args{
				storage:  nil,
				log:      testLog,
				paramURI: "",
				body:     []byte(`{"Id":1,"Homework_Id":1,"Question_Id":1,"Order":1}`),
			},
			want: want{
				statusCode: fiber.StatusConflict,
				body:       []byte(""),
			},
		},
		{
			name: "question is not found",
			args: args{
				storage:  nil,
				log:      testLog,
				paramURI: "",
				body:     []byte(`{"Id":1,"Homework_Id":1,"Question_Id":1,"Order":1}`),
			},
			want: want{
				statusCode: fiber.StatusBadRequest,
				body:       []byte(""),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testApp := fiber.New()
			testApp.Post("/", Insert(mockStorage, testLog))
			defer utils.ExecuteWithLogError(testApp.Shutdown, testLog)

			request, errReq := http.NewRequest(http.MethodPost, tt.args.paramURI, bytes.NewBuffer(tt.args.body))
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
