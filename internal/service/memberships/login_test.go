package memberships

import (
	"fmt"
	"testing"

	"github.com/joshuatheokurniawansiregar/music-catalog/internal/configs"
	"github.com/joshuatheokurniawansiregar/music-catalog/internal/model/memberships"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
)

func Test_service_Login(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()
	
	mockRepo:= NewMockrepository(ctrlMock)

	type args struct {
		request memberships.LoginRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		mockFn func(args args)
	}{
		// TODO: Add test cases.
		{
			name: "success",
			args: args{
				request: memberships.LoginRequest{
					Email: "test1@gmail.com",
					Password: "password",
				},
			},
			wantErr: false,
			mockFn: func(args args){
				mockRepo.EXPECT().GetUser(args.request.Email, "", uint64(0)).Return(&memberships.User{
					Model: gorm.Model{
						ID: 1,
					},
					Email: "test1@gmail.com",
					Username: "testusdsername",
					Password: "$2a$10$J52rSJGm7NiYgh.glzU7b.YpQ2oDDZ8TzHlmoskdDSvLwDQy2ujj.",
				}, nil)
			},
		},
		{
			name: "failed when GetUser",
			args: args{
				request: memberships.LoginRequest{
					Email: "test1@gmail.com",
					Password: "password",
				},
			},
			wantErr: true,
			mockFn: func(args args){
				mockRepo.EXPECT().GetUser(args.request.Email, "", uint64(0)).Return(nil, assert.AnError)
			},
		},
		{
			name: "failed password is not matched",
			args: args{
				request: memberships.LoginRequest{
					Email: "test1@gmail.com",
					Password: "password",
				},
			},
			wantErr: true,
			mockFn: func(args args){
				mockRepo.EXPECT().GetUser(args.request.Email, "", uint64(0)).Return(&memberships.User{
					Model: gorm.Model{
						ID: 1,
					},
					Email: "test1@gmail.com",
					Username: "testusdsername",
					Password: "wrong password",
				}, nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			var service *service = &service{
				membershipsRepo: mockRepo,
				cfg: &configs.Config{
					Service: configs.Service{
						SecretKey: "abc",
					},
				},
			}
			got, err := service.Login(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr{
				fmt.Printf("test case %+v", tt.name)
				assert.NotEmpty(t, got)
			}else{
				fmt.Printf("\ntest case %s", tt.name)
				assert.Empty(t, got)
			}
		})
	}
}
