package svc

import (
	"time"

	"github.com/bolg-developers/MikanMusic-API/config"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

func CreateUserToken(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString([]byte(config.Env().TokenSecretKey))
	if err != nil {
		return "", errors.WithStack(err)
	}
	return tokenString, nil
}
