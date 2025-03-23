package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "secret"

func GenerateToken(email string, userId int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretKey))

}

func VerifyToken(token string) (int, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return 0, err
	}

	if !parsedToken.Valid {
		return 0, jwt.ErrSignatureInvalid
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return 0, jwt.ErrSignatureInvalid
	}

	email := claims["email"].(string)
	userId := int(claims["userId"].(float64))

	if email == "" || userId == 0 {
		return 0, jwt.ErrSignatureInvalid
	}
	return userId, nil
}
