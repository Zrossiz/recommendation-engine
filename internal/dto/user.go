package dto

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type JWTPayload struct {
	ID       int
	Username string
}

type SuccessAuthenticate struct {
	AcessToken   string
	RefreshToken string
	Err          error
}
