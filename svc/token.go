package svc

import (
	"time"

	"github.com/bolg-developers/MikanMusic-API/config"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

const (
	ClaimsKeyUserID = "userID"
	ClaimsKeyEXP    = "exp"
)

func CreateUserToken(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		ClaimsKeyUserID: userID,
		ClaimsKeyEXP:    time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString([]byte(config.Env().TokenSecretKey))
	if err != nil {
		return "", errors.WithStack(err)
	}
	return tokenString, nil
}

func ParseClaimsFromUserToken(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.Errorf("signing methodが不正です: %v", token.Header["alg"])
		}
		return config.Env().TokenSecretKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("トークンが不正です")
}
