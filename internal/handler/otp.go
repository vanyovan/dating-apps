package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/vanyovan/dating-apps/internal/entity"
	"github.com/vanyovan/dating-apps/internal/usecase"
)

type OTPRequest struct {
	UserId string `json:"user_id"`
}

type Handler struct {
	SignUpUc   usecase.SignUpUsecase
	PackageUc  usecase.PackageUsecase
	OtpService *usecase.OtpService
}

func (h *Handler) RequestOTP(w http.ResponseWriter, r *http.Request) {
	request := OTPRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid request payload"))
		return
	}

	param := entity.OTPRequestParam{
		UserId: request.UserId,
	}

	result, err := h.OtpService.SetOTP(context.TODO(), param)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Set OTP Failed"))
		return
	}

	jsonResponse, err := json.Marshal(result)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(jsonResponse))
}
