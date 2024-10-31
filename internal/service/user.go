package service

import (
	"engine/internal/config"
	"engine/internal/dto"
	"engine/internal/model"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type CustomClaims struct {
	UserID   int64  `json:"userId"`
	Username string `json:"userName"`
	jwt.RegisteredClaims
}

type UserStore interface {
	Create(userDTO dto.User) (bool, error)
	Delete(id int64) (bool, error)
	GetUserById(id int64) (*model.User, error)
	GetUserByName(name string) (*model.User, error)
	Update(id int64, userDTO dto.User) (bool, error)
}

type UserService struct {
	db  UserStore
	log *zap.Logger
	cfg *config.Config
}

func NewUserService(db UserStore, log *zap.Logger, cfg *config.Config) *UserService {
	return &UserService{
		db:  db,
		log: log,
		cfg: cfg,
	}
}

func (u *UserService) Registration(user dto.User) dto.SuccessAuthenticate {
	success := dto.SuccessAuthenticate{}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), u.cfg.Cost)
	if err != nil {
		success.Err = err
		return success
	}
	user.Password = string(hashedPassword)

	_, err = u.db.Create(user)
	if err != nil {
		success.Err = err
		return success
	}

	createdUser, err := u.db.GetUserByName(user.Username)

	payload := dto.JWTPayload{
		ID:       int(createdUser.ID),
		Username: createdUser.Name,
	}

	generateTokenPair(
		payload,
		u.cfg.JWTAccessSecret,
		u.cfg.JWTRefreshSecret,
		&success,
	)
	if success.Err != nil {
		u.log.Error("create tokens err", zap.Error(err))
	}

	return success
}

func generateTokenPair(
	payload dto.JWTPayload,
	accessSecret string,
	refreshSecret string,
	success *dto.SuccessAuthenticate,
) {
	claims := &CustomClaims{
		UserID:   int64(payload.ID),
		Username: payload.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			Issuer:    "exampleIssuer",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessToken, err := token.SignedString(accessSecret)
	if err != nil {
		success.Err = err
		return
	}

	claims.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(24 * time.Hour * 30))

	refreshToken, err := token.SignedString(refreshSecret)
	if err != nil {
		success.Err = err
		return
	}

	success.AcessToken = accessToken
	success.RefreshToken = refreshToken
}
