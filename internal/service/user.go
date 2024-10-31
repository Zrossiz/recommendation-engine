package service

import (
	"engine/internal/apperrors"
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
	res := dto.SuccessAuthenticate{}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), u.cfg.Cost)
	if err != nil {
		res.Err = err
		return res
	}
	user.Password = string(hashedPassword)

	_, err = u.db.Create(user)
	if err != nil {
		res.Err = err
		return res
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
		&res,
	)
	if res.Err != nil {
		u.log.Error("create tokens err", zap.Error(err))
	}

	return res
}

func (u *UserService) Login(user dto.User) *dto.SuccessAuthenticate {
	res := dto.SuccessAuthenticate{}
	existUser, err := u.db.GetUserByName(user.Username)
	if err != nil {
		res.Err = err
		return &res
	}

	err = bcrypt.CompareHashAndPassword([]byte(existUser.Password), []byte(user.Password))
	if err != nil {
		res.Err = apperrors.ErrInvalidPassword
		return &res
	}

	payload := dto.JWTPayload{
		ID:       int(existUser.ID),
		Username: existUser.Name,
	}

	generateTokenPair(
		payload,
		u.cfg.JWTAccessSecret,
		u.cfg.JWTRefreshSecret,
		&res,
	)

	return &res
}

func generateTokenPair(
	payload dto.JWTPayload,
	accessSecret string,
	refreshSecret string,
	res *dto.SuccessAuthenticate,
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
		res.Err = err
		return
	}

	claims.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(24 * time.Hour * 30))

	refreshToken, err := token.SignedString(refreshSecret)
	if err != nil {
		res.Err = err
		return
	}

	res.AcessToken = accessToken
	res.RefreshToken = refreshToken
}
