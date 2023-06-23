package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	_ "github.com/mattn/go-sqlite3"
	"github.com/vanyovan/dating-apps/internal/handler"
	"github.com/vanyovan/dating-apps/internal/repo"
	"github.com/vanyovan/dating-apps/internal/usecase"
)

func main() {
	db, err := sql.Open("sqlite3", "../database.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	otpRepo := repo.NewOTPRepo(db)
	signupUsecase := usecase.NewSignUpUsecase(db)
	packageUsecase := usecase.NewPackageUsecase(db)
	otpUsecase := usecase.NewOtpService(otpRepo)

	handler := &handler.Handler{
		SignUpUc:   signupUsecase,
		PackageUc:  packageUsecase,
		OtpService: otpUsecase,
	}
	router := chi.NewRouter()

	router.Method(http.MethodPost, "/signup", http.HandlerFunc(handler.SignUpHandler))
	router.Method(http.MethodPost, "/premium", http.HandlerFunc(handler.UpdateUserPremiumHandler))
	router.Method(http.MethodPost, "/otp", http.HandlerFunc(handler.RequestOTP))

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Println("server listening on", server.Addr)
	server.ListenAndServe()
}
