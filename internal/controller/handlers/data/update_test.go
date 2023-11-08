package data

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
		mockStorage.EXPECT().UpdateData(gomock.Any(), gomock.Any()).Return(nil),
		mockStorage.EXPECT().UpdateData(gomock.Any(), gomock.Any()).Return(nil),
		mockStorage.EXPECT().UpdateData(gomock.Any(), gomock.Any()).Return(nil),
		mockStorage.EXPECT().UpdateData(gomock.Any(), gomock.Any()).Return(fmt.Errorf("test err")),
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
				body: []byte(`{
									"Homework": {
										"Id": 1,
										"Name": "hw_name",
										"Questions": [
											{
												"Id": 1,
												"Name": "q1_name",
												"Content": {
													"Task": "task1",
													"Answer": "answer1",
													"Solution": "solution1"
												}
											}
										]
									}
								}`),
			},
			want: want{
				statusCode: fiber.StatusOK,
				body:       []byte(`{"Homework":{"Id":1,"Name":"hw_name","Questions":[{"Id":1,"Name":"q1_name","Content":{"Task":"task1","Answer":"answer1","Solution":"solution1"}}]}}`),
			},
		},
		{
			name: "broken json",
			args: args{
				storage:  nil,
				log:      testLog,
				paramURI: "",
				body: []byte(`{
									Homework": {
										"Id": 1,
										"Name": "hw_name",
										"Questions": [
											{
												"Id": 1,
												"Name": "q1_name",
												"Content": {
													"Task": "task1",
													"Answer": "answer1",
													"Solution": "solution1"
												}
											}
										]
									}
								}`),
			},
			want: want{
				statusCode: fiber.StatusBadRequest,
				body:       []byte(""),
			},
		},
		{
			name: "empty json body",
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
			name: "missing ID in URI and in body",
			args: args{
				storage:  nil,
				log:      testLog,
				paramURI: "/",
				body: []byte(`{
									"Homework": {
										"Name": "hw_name",
										"Questions": [
											{
												"Id": 1,
												"Name": "q1_name",
												"Content": {
													"Task": "task1",
													"Answer": "answer1",
													"Solution": "solution1"
												}
											}
										]
									}
								}`),
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
				body: []byte(`{
									"Homework": {
										"Name": "hw_name",
										"Questions": [
											{
												"Id": 1,
												"Name": "q1_name",
												"Content": {
													"Task": "task1",
													"Answer": "answer1",
													"Solution": "solution1"
												}
											}
										]
									}
								}`),
			},
			want: want{
				statusCode: fiber.StatusOK,
				body:       []byte(`{"Homework":{"Id":1,"Name":"hw_name","Questions":[{"Id":1,"Name":"q1_name","Content":{"Task":"task1","Answer":"answer1","Solution":"solution1"}}]}}`),
			},
		},
		{
			name: "ID in body only",
			args: args{
				storage:  nil,
				log:      testLog,
				paramURI: "/",
				body: []byte(`{
									"Homework": {
										"Id": 1,
										"Name": "hw_name",
										"Questions": [
											{
												"Id": 1,
												"Name": "q1_name",
												"Content": {
													"Task": "task1",
													"Answer": "answer1",
													"Solution": "solution1"
												}
											}
										]
									}
								}`),
			},
			want: want{
				statusCode: fiber.StatusOK,
				body:       []byte(`{"Homework":{"Id":1,"Name":"hw_name","Questions":[{"Id":1,"Name":"q1_name","Content":{"Task":"task1","Answer":"answer1","Solution":"solution1"}}]}}`),
			},
		},
		{
			name: "empty homework name",
			args: args{
				storage:  nil,
				log:      testLog,
				paramURI: "",
				body: []byte(`{
									"Homework": {
										"Id": 1,
										"Name": "",
										"Questions": [
											{
												"Id": 1,
												"Name": "q1_name",
												"Content": {
													"Task": "task1",
													"Answer": "answer1",
													"Solution": "solution1"
												}
											}
										]
									}
								}`),
			},
			want: want{
				statusCode: fiber.StatusBadRequest,
				body:       []byte(""),
			},
		},
		{
			name: "missing task in content",
			args: args{
				storage:  nil,
				log:      testLog,
				paramURI: "",
				body: []byte(`{
									"Homework": {
										"Id": 1,
										"Name": "hw1_name",
										"Questions": [
											{
												"Id": 1,
												"Name": "q1_name",
												"Content": {
													"Answer": "answer1",
													"Solution": "solution1"
												}
											}
										]
									}
								}`),
			},
			want: want{
				statusCode: fiber.StatusBadRequest,
				body:       []byte(""),
			},
		},
		{
			name: "missing question name",
			args: args{
				storage:  nil,
				log:      testLog,
				paramURI: "",
				body: []byte(`{
									"Homework": {
										"Id": 1,
										"Name": "hw1_name",
										"Questions": [
											{
												"Id": 1,
												"Name": "",
												"Content": {
													"Task": "task1",
													"Answer": "answer1",
													"Solution": "solution1"
												}
											}
										]
									}
								}`),
			},
			want: want{
				statusCode: fiber.StatusBadRequest,
				body:       []byte(""),
			},
		},
		{
			name: "err from storage",
			args: args{
				storage:  nil,
				log:      testLog,
				paramURI: "",
				body: []byte(`{
									"Homework": {
										"Id": 1,
										"Name": "hw_name",
										"Questions": [
											{
												"Id": 1,
												"Name": "q1_name",
												"Content": {
													"Task": "task1",
													"Answer": "answer1",
													"Solution": "solution1"
												}
											}
										]
									}
								}`),
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
