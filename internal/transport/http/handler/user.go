package handler

import (
	"encoding/json"
	"engine/internal/apperrors"
	"engine/internal/dto"
	"net/http"
	"time"

	"go.uber.org/zap"
)

type UserHandler struct {
	service UserService
	log     *zap.Logger
}

type UserService interface {
	Registration(user dto.User) *dto.SuccessAuthenticate
	Login(user dto.User) *dto.SuccessAuthenticate
}

func NewUserHandler(serv UserService, log *zap.Logger) *UserHandler {
	return &UserHandler{
		service: serv,
		log:     log,
	}
}

func (u *UserHandler) Registration(rw http.ResponseWriter, r *http.Request) {
	var registrationDTO dto.User

	err := json.NewDecoder(r.Body).Decode(&registrationDTO)
	if err != nil {
		http.Error(rw, "invalid body", http.StatusBadRequest)
		return
	}

	res := u.service.Registration(registrationDTO)
	if res.Err != nil {
		if res.Err == apperrors.ErrAlreadyExist {
			http.Error(rw, "user already exist", http.StatusConflict)
			return
		}
		u.log.Error("user registration error", zap.Error(err))
		http.Error(rw, "internal server error", http.StatusInternalServerError)
		return
	}

	refreshTokenCokie := http.Cookie{
		Name:     "refreshtoken",
		Value:    res.RefreshToken,
		Path:     "/",
		Expires:  time.Now().Add(2 * time.Hour * 24 * 30),
		HttpOnly: true,
		Secure:   false,
	}

	accessTokenCookie := http.Cookie{
		Name:     "accesstoken",
		Value:    res.AcessToken,
		Path:     "/",
		Expires:  time.Now().Add(15 * time.Minute),
		HttpOnly: true,
		Secure:   false,
	}

	http.SetCookie(rw, &refreshTokenCokie)
	http.SetCookie(rw, &accessTokenCookie)
	response := map[string]string{
		"message": "registration successful",
	}

	rw.Header().Set("content-type", "application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(response)
}

func (u *UserHandler) Login(rw http.ResponseWriter, r *http.Request) {
	var loginDTO dto.User

	err := json.NewDecoder(r.Body).Decode(&loginDTO)
	if err != nil {
		http.Error(rw, "invalid body", http.StatusBadRequest)
		return
	}

	res := u.service.Login(loginDTO)
	if res.Err != nil {
		if res.Err == apperrors.ErrNotFound || res.Err == apperrors.ErrInvalidPassword {
			http.Error(rw, "invalid login or password", http.StatusUnauthorized)
			return
		}

		u.log.Error("login user error", zap.Error(err))
		http.Error(rw, "internal server error", http.StatusInternalServerError)
		return
	}

	refreshTokenCokie := http.Cookie{
		Name:     "refreshtoken",
		Value:    res.RefreshToken,
		Path:     "/",
		Expires:  time.Now().Add(2 * time.Hour * 24 * 30),
		HttpOnly: true,
		Secure:   false,
	}

	accessTokenCookie := http.Cookie{
		Name:     "accesstoken",
		Value:    res.AcessToken,
		Path:     "/",
		Expires:  time.Now().Add(15 * time.Minute),
		HttpOnly: true,
		Secure:   false,
	}

	http.SetCookie(rw, &refreshTokenCokie)
	http.SetCookie(rw, &accessTokenCookie)
	response := map[string]string{
		"message": "login successful",
	}

	rw.Header().Set("content-type", "application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(response)
}
