package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/vanyovan/dating-apps/internal/entity"
	mock_usecase "github.com/vanyovan/dating-apps/internal/usecase/mock_otp"
)

func TestUsecase_OTPRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	any := gomock.Any()

	otp := mock_usecase.NewMockOtpServiceProvider(ctrl)

	service := NewOtpService(otp)

	tests := []struct {
		name    string
		param   entity.OTPRequestParam
		want    entity.OTPResponse
		wantErr bool
		mock    func()
	}{
		{"failed_set_otp", entity.OTPRequestParam{UserId: "test"},
			entity.OTPResponse{UserId: "test"},
			true,
			func() {
				otp.EXPECT().SetOTP(context.TODO(), any).Return(entity.OTPResponse{}, errors.New("errors"))
			}},
		{"success_set_otp", entity.OTPRequestParam{UserId: "test"},
			entity.OTPResponse{UserId: "test"},
			false,
			func() {
				otp.EXPECT().SetOTP(context.TODO(), any).Return(entity.OTPResponse{UserId: "test", Otp: "123"}, nil)
			}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			_, err := service.SetOTP(context.TODO(), tt.param)
			if (err != nil) != tt.wantErr {
				t.Errorf("Usecase.SetOTP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("Usecase.HandoverExportData() = %v, want %v", got, tt.want)
			// }
		})
	}
}
