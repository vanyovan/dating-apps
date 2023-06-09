package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vanyovan/dating-apps/internal/entity"
	"github.com/vanyovan/dating-apps/internal/usecase"
)

type OTPRequest struct {
	UserId string `json:"user_id"`
}

type Handler struct {
	SignUpUc  usecase.SignUpUsecase
	PackageUc usecase.PackageUsecase
	OtpUC     usecase.OtpUsecase
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

	err = h.OtpUC.SetOTP(r.Context(), param)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Set OTP Failed"))
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "OTP Request Success")
}
