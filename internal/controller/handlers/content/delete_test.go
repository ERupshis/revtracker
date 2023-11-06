package content

import (
	"fmt"
	"io"
	"net/http"
	"testing"

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

func TestDelete(t *testing.T) {
	testLog, _ := logger.CreateMock()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mocks.NewMockBaseStorage(ctrl)
	gomock.InOrder(
		mockStorage.EXPECT().DeleteContentByID(gomock.Any(), gomock.Any()).Return(nil),
		mockStorage.EXPECT().DeleteContentByID(gomock.Any(), gomock.Any()).Return(fmt.Errorf("error")),
	)

	testApp := fiber.New()
	testApp.Delete("/:ID", Delete(mockStorage, testLog))
	defer utils.ExecuteWithLogError(testApp.Shutdown, testLog)

	port := 3011
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
				body:       []byte(""),
			},
		},
		{
			name: "incorrect ID",
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
			name: "error from DB",
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
			request, errReq := http.NewRequest(http.MethodDelete, utilsReform.HostTest+fmt.Sprintf("%d", port)+tt.args.paramURI, nil)
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
