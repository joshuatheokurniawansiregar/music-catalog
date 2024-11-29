package memberships

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joshuatheokurniawansiregar/music-catalog/internal/model/memberships"
	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)

func TestHandler_Login(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()
	mockSvc := NewMockservice(ctrlMock)

	tests := []struct {
		name string
		mockFn func()
		expectedStatusCode int64
		expectedBody memberships.LoginResponse
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "success",
			mockFn: func ()  {
				mockSvc.EXPECT().Login(memberships.LoginRequest{
					Email: "test@gmail.com",
					Password: "password",
				}).Return("accessToken",nil)
			},
			expectedStatusCode: 200,
			expectedBody: memberships.LoginResponse{
				AccessToken: "accessToken",
			},
			wantErr: false,
		},
		{
			name: "failed",
			mockFn: func ()  {
				mockSvc.EXPECT().Login(memberships.LoginRequest{
					Email: "test@gmail.com",
					Password: "password",
				}).Return("", assert.AnError)
			},
			expectedStatusCode: 500,
			expectedBody: memberships.LoginResponse{
				AccessToken: "",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn()
			var handler *Handler = &Handler{
				Engine: gin.New(),
				service: mockSvc,
			}
			handler.RegisterRoute()
			w:= httptest.NewRecorder()
			endpoint:= `/memberships/login`
			model:= memberships.LoginRequest{
				Email: "test@gmail.com",
				Password:"password",
			}
			val, err := json.Marshal(model)
			assert.NoError(t, err)

			body:= bytes.NewReader(val)
			req, err := http.NewRequest(http.MethodPost, endpoint, body)
			assert.NoError(t, err)
			handler.ServeHTTP(w, req)
			assert.Equal(t, tt.expectedStatusCode, int64(w.Code))

			if(!tt.wantErr){
				res:= w.Result()
				defer res.Body.Close()
				
				response := memberships.LoginResponse{}
				err = json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)

				assert.Equal(t, tt.expectedBody, response)
			}
		})
	}
}
