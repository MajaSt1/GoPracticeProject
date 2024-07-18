package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "supersecret"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

func verifyToken(token string) error {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC) // check the type of t (HMC = veriosn of SigningMethodHS256)
		if !ok {
			return nil, errors.New("Unexpected signing method.")
		}

		return secretKey, nil
	})
	if err != nil {
		return errors.New("Could not parse token.")
	}

	tokenIsValid := parsedToken.Valid
	if !tokenIsValid {
		return errors.New("Invalid token!")
	}

	_, ok := parsedToken.Claims.(jwt.MapClaims) //access the data in that token == jwt.NewWithClaims
	if !ok {
		return errors.New("Invalid token claims.")
	}

	// email := claims["email"].(string)
	// userId := claims["userId"].(int64)
	return nil
}
