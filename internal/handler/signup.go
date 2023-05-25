package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"

	"github.com/vanyovan/dating-apps/internal/entity"
	"github.com/vanyovan/dating-apps/internal/usecase"
)

type SignUpRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Handler struct {
	SignUpUc  usecase.SignUpUsecase
	PackageUc usecase.PackageUsecase
}

func (h *Handler) SignUpHandler(w http.ResponseWriter, r *http.Request) {
	request := SignUpRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid request payload"))
		return
	}

	err = request.validateSignup()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	param := entity.CreateSignUpParam{
		Username: request.Username,
		Password: request.Password,
		Email:    request.Email,
	}

	// get username, if there is same username cannot inserted
	user, err := h.SignUpUc.GetUserByUsername(r.Context(), param.Username)
	if err != nil || user.ID != 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Username already exists"))
		return
	}

	err = h.SignUpUc.CreateSignUp(r.Context(), param)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to create user"))
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Signup successful")
}

func (s SignUpRequest) validateSignup() error {
	emailRegex := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`

	match, _ := regexp.MatchString(emailRegex, s.Email)
	if match {
		return nil
	} else {
		return errors.New("Email is not valid")
	}
}
