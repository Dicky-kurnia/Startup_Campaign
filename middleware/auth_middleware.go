package middleware

import "github.com/dgrijalva/jwt-go"

type AuthService interface {
	GenerateToken(userID int) (string, error)
}

type jwtService struct {
}

func NewAuthService() *jwtService {
	return &jwtService{}
}

var SECRET_KEY = []byte("BWASTARTUP_s3cr3T__k3Y")

func (service *jwtService) GenerateToken(userID int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}
	return signedToken, nil
}
