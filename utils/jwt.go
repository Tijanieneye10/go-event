package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const secretKey = "eventKey"

func GenerateToken(email string, userId int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	signedString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return signedString, nil
}

func VerifyToken(token string) error {
	parseToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing method error")
		}
		return secretKey, nil
	})
	if err != nil {
		return errors.New(err.Error())
	}

	tokenIsValid := parseToken.Valid

	if !tokenIsValid {
		return errors.New("invalid token")
	}

	claims, ok := parseToken.Claims.(jwt.MapClaims)

	if !ok {
		return errors.New("invalid tokens")
	}

	//email := claims["email"].(string)
	//userId := claims["userId"]

	return nil
}
