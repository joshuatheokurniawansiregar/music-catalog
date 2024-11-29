package memberships

import (
	"testing"

	"github.com/joshuatheokurniawansiregar/music-catalog/internal/configs"
	"github.com/joshuatheokurniawansiregar/music-catalog/internal/model/memberships"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
)

func Test_service_SignUp(t *testing.T) {
	ctrlMock:= gomock.NewController(t)
	defer ctrlMock.Finish()

	mockRepo:= NewMockrepository(ctrlMock)

	type args struct {
		request memberships.SignupRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		mockFn func(args args)
	}{
		// TODO: Add test cases.
		{
			name: "succeed when Createuser",
			args: args{
				request: memberships.SignupRequest{
					Email: "test@gmail.com",
					Username: "username",
					Password: "password",
				},
			},
			wantErr: false,
			mockFn: func(args args) {
				mockRepo.EXPECT().GetUser(args.request.Email, args.request.Username, uint64(0)).Return(nil, gorm.ErrRecordNotFound)
				mockRepo.EXPECT().CreateUser(gomock.Any()).Return(nil)
			},
		},
		{
			name: "failed when Getuser",
			args: args{
				request: memberships.SignupRequest{
					Email: "test@gmail.com",
					Username: "testusername",
					Password: "password",
				},
			},
			wantErr: true,
			mockFn: func(args args){
				mockRepo.EXPECT().GetUser(args.request.Email, args.request.Username, uint64(0)).Return(nil, assert.AnError)
			},
		},
		{
			name: "failed when Createuser",
			args: args{
				request: memberships.SignupRequest{
					Email: "test@gmail.com",
					Username: "testusername",
					Password: "password",
				},
			},
			wantErr: true,
			mockFn: func(args args){
				mockRepo.EXPECT().GetUser(args.request.Email, args.request.Username, uint64(0))
				mockRepo.EXPECT().CreateUser(gomock.Any()).Return(assert.AnError)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			s:= &service{
				membershipsRepo: mockRepo,
				cfg: &configs.Config{},
			}
			if err := s.SignUp(tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("service.SignUp() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
