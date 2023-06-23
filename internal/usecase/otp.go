package usecase

import (
	"context"

	"fmt"
"github.com/vanyovan/dating-apps/internal/repo"
	"github.com/vanyovan/dating-apps/internal/entity"
)

type OtpService struct {
	OtpRepo repo.OtpRepo
}

type OtpServiceProvider interface {
	SetOTP(ctx context.Context, param entity.OTPRequestParam) (result entity.OTPResponse, err error)
}

func NewOtpService(OtpRepo repo.OtpRepo) *OtpService {
	return &OtpService{
		OtpRepo: OtpRepo,
	}
}

func (uc *OtpService) SetOTP(ctx context.Context, param entity.OTPRequestParam) (result entity.OTPResponse, err error) {

	// test := uc.db.SetOTP(ctx, param)
	fmt.Println("TEST")
	result, err = uc.OtpRepo.SetOTP(ctx, param)

	// tx, err := uc.db.Begin()

	// valid2minutes := time.Now().Add(2 * time.Minute)
	// otp := generateOTP()

	// _, err = tx.ExecContext(ctx, "INSERT INTO otp (user_id, otp, valid) VALUES (?, ?, ?)", param.UserId, otp, valid2minutes)
	// if err != nil {Context(ctx, "INSERT INTO otp (user_id, otp, valid) VALUES (?, ?, ?)", param.UserId, otp, valid2minutes)
	// if err != nil{Context(ctx, "INSERT INTO otp (user_id, otp, valid) VALUES (?, ?, ?)", param.UserId, otp, valid2minutes)
	// if err != nil
	// 	tx.Rollback()
	// 	turn result, fmt.Errorf("failed to create OTP: %w", err)
		// turn result, fmt.Errorf("failed to create OTP: %w", err)
//}

	// err = tx.Commit
	// if err != nil {
		// turn result, fmt.Errorf("failed to create user: %w", err)
// }

	// result = entity.OTPRonse{
	// 	UserId: pa.UserId,
	// 	:    otp,
	// }
	return result, err
}
