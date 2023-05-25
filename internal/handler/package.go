package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vanyovan/dating-apps/internal/entity"
)

type PackageRequest struct {
	PackageId int    `json:"premium_id"`
	Username  string `json:"username"`
}

func (h *Handler) UpdateUserPremiumHandler(w http.ResponseWriter, r *http.Request) {
	request := PackageRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid request payload"))
		return
	}
	param := entity.UpdateUserPremium{
		PackageId: request.PackageId,
		Username:  request.Username,
	}

	//get package information
	packageEntity, err := h.PackageUc.GetPackageByPackageId(r.Context(), param.PackageId)
	if packageEntity.PackageId == 0 {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Invalid package id"))
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to retrieve package data"))
		return
	}

	//get user information
	user, err := h.SignUpUc.GetUserByUsername(r.Context(), param.Username)
	if err != nil || user.ID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Username not exist"))
		return
	}

	//update user premium
	err = h.PackageUc.UpdateUserPremium(r.Context(), param)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to create user"))
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Premium successful")
}
