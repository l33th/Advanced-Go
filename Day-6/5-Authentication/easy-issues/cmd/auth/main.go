package main

import (
	"github.com/Advanced-Go/Day-6/5-Authentication/easy-issues"
	"github.com/Advanced-Go/Day-6/5-Authentication/easy-issues/application"
	"github.com/Advanced-Go/Day-6/5-Authentication/easy-issues/domain"
	"github.com/Advanced-Go/Day-6/5-Authentication/easy-issues/persistence/db"
	"github.com/Advanced-Go/Day-6/5-Authentication/easy-issues/web"
	"github.com/Advanced-Go/Day-6/5-Authentication/easy-issues/web/handler"
	"log"
	"net/http"
	"time"
)

func main() {
	authRepo := db.NewAuthRepository()
	authService := application.AuthService{
		AuthRepository: authRepo,
	}
	authController := web.AuthController{
		AuthService: authService,
		Secret:      application.Secret,
	}

	pwhash, _ := easy_issues.HashPassword("opensessame")

	authRepo.Create(&domain.UserRegistration{
		Email:        "aladdin@mail.com",
		Uuid:         "1234",
		PasswordHash: pwhash,
	})

	mux := http.NewServeMux()
	mux.HandleFunc("/auth/login", authController.Login)
	mux.HandleFunc("/auth/verify", handler.JWTAuthHandler(authController.Verify))

	server := &http.Server{
		Addr:           ":8090",
		Handler:        mux,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    120 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(server.ListenAndServe())
}
