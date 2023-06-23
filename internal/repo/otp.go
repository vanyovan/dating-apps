package repo

import (
	"context"
	"crypto/rand"
	"database/sql"
	"fmt"
	"math/big"
	"time"

	"github.com/vanyovan/dating-apps/internal/entity"
)

type Repo struct {
	db *sql.DB
}

type OtpRepo interface {
	SetOTP(ctx context.Context, param entity.OTPRequestParam) (result entity.OTPResponse, err error)
}

func NewOTPRepo(db *sql.DB) OtpRepo {
	return &Repo{
		db: db,
	}
}

func (uc *Repo) SetOTP(ctx context.Context, param entity.OTPRequestParam) (result entity.OTPResponse, err error) {
	tx, err := uc.db.Begin()

	valid2minutes := time.Now().Add(2 * time.Minute)
	otp := generateOTP()

	_, err = tx.ExecContext(ctx, "INSERT INTO otp (user_id, otp, valid) VALUES (?, ?, ?)", param.UserId, otp, valid2minutes)
	if err != nil {
		tx.Rollback()
		return result, fmt.Errorf("failed to create OTP: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return result, fmt.Errorf("failed to create user: %w", err)
	}

	result = entity.OTPResponse{
		UserId: param.UserId,
		Otp:    otp,
	}

	return result, nil
}

func generateOTP() string {
	max := big.NewInt(999999)
	otp, _ := rand.Int(rand.Reader, max)
	return fmt.Sprintf("%06d", otp)
}
