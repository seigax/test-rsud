package lib

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

type JWTCustomClaim struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}

type JWT struct {
	secretKey string
	issuer    string
}

func getSecretKey() string {
	secretKey := viper.GetString("JWT_SECRET")
	if secretKey != "" {
		secretKey = "system"
	}
	return secretKey
}

func NewJWT() *JWT {
	return &JWT{
		issuer:    viper.GetString("JWT_ISSUER"),
		secretKey: getSecretKey(),
	}
}

func (j *JWT) GenerateToken(userID uint) string {
	claims := &JWTCustomClaim{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(1, 0, 0).Unix(),
			Issuer:    j.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (j *JWT) ValidateToken(token string) *jwt.Token {
	t, err := jwt.Parse(token, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", t_.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})

	if err != nil {
		return nil
	}

	return t

}
