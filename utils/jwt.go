package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

const secretKey = "superSecretKey"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

func VerifyToken(tokenString string) error {
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.NewValidationError("Unexpected signing method", jwt.ValidationErrorSignatureInvalid)
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return errors.New("Could not parse token: " + err.Error())
	}
	tokenIsValid := parsedToken.Valid
	if !tokenIsValid {
		return errors.New("Invalid token.")
	}

	// claims, ok := parsedToken.Claims.(jwt.MapClaims)
	// if !ok {
	// 	return errors.New("Invalid token claims.")
	// }

	// email := claims["email"].(string)
	// userId := claims["userId"].(int64)
	return nil
}
