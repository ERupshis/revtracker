package middleware

import (
	"database/sql"
	"fmt"
	"net/http"
	"testing"

	"github.com/erupshis/revtracker/internal/auth/data"
	"github.com/erupshis/revtracker/internal/auth/jwtgenerator"
	"github.com/erupshis/revtracker/internal/logger"
	"github.com/erupshis/revtracker/internal/utils"
	"github.com/erupshis/revtracker/mocks"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testHandler struct {
	Message string
}

func (th testHandler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	_, _ = fmt.Fprintln(w, th.Message)
}

func TestAuthorizeUser(t *testing.T) {
	log, _ := logger.CreateZapLogger("info")
	defer log.Sync()

	jwtGen := jwtgenerator.Create("secret_key", 3, log)
	validToken, _ := jwtGen.BuildJWTString(2)

	user1 := data.User{
		Login:    "u1",
		Password: "p1",
		ID:       1,
		Role:     data.RoleUser,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mocks.NewMockBaseUsersStorage(ctrl)
	gomock.InOrder(
		mockStorage.EXPECT().SelectUserByID(gomock.Any(), gomock.Any()).Return(&user1, nil),
		mockStorage.EXPECT().SelectUserByID(gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("db error")),
		mockStorage.EXPECT().SelectUserByID(gomock.Any(), gomock.Any()).Return(nil, sql.ErrNoRows),
		mockStorage.EXPECT().SelectUserByID(gomock.Any(), gomock.Any()).Return(&user1, nil),
	)

	type args struct {
		authorizationHeader string
		role                int
	}
	type want struct {
		statusCode int
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "valid",
			args: args{
				authorizationHeader: "Bearer " + validToken,
				role:                data.RoleUser,
			},
			want: want{
				statusCode: http.StatusOK,
			},
		},
		{
			name: "without Authorization header",
			args: args{
				authorizationHeader: "",
				role:                data.RoleUser,
			},
			want: want{
				statusCode: http.StatusUnauthorized,
			},
		},
		{
			name: "invalid token",
			args: args{
				authorizationHeader: string("Basic ") + validToken,
				role:                data.RoleUser,
			},
			want: want{
				statusCode: http.StatusUnauthorized,
			},
		},
		{
			name: "invalid token",
			args: args{
				authorizationHeader: validToken,
				role:                data.RoleUser,
			},
			want: want{
				statusCode: http.StatusUnauthorized,
			},
		},
		{
			name: "db error",
			args: args{
				authorizationHeader: string("Bearer ") + validToken,
				role:                data.RoleUser,
			},
			want: want{
				statusCode: http.StatusInternalServerError,
			},
		},
		{
			name: "user is not registered",
			args: args{
				authorizationHeader: string("Bearer ") + validToken,
				role:                data.RoleUser,
			},
			want: want{
				statusCode: http.StatusUnauthorized,
			},
		},
		{
			name: "lack or permission",
			args: args{
				authorizationHeader: string("Bearer ") + validToken,
				role:                data.RoleAdmin,
			},
			want: want{
				statusCode: http.StatusForbidden,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testApp := fiber.New()
			testApp.Use(AuthorizeUser(tt.args.role, mockStorage, jwtGen, log))
			testApp.Post("/handler", adaptor.HTTPHandler(testHandler{}))
			defer utils.ExecuteWithLogError(testApp.Shutdown, log)

			request, errReq := http.NewRequest(http.MethodPost, "/handler", nil)
			require.NoError(t, errReq)

			if tt.args.authorizationHeader != "" {
				request.Header.Add("Authorization", tt.args.authorizationHeader)
			}

			response, errResp := testApp.Test(request)
			require.NoError(t, errResp)
			defer func() {
				_ = response.Body.Close()
			}()

			assert.Equal(t, tt.want.statusCode, response.StatusCode)
		})
	}
}
