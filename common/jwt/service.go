package jwt

import (
	"errors"
	"fmt"
	"press/config"
	"time"

	jwtLib "github.com/dgrijalva/jwt-go"
)

type Interface interface {
	GenerateFromUserID(id string) (string, error)
	GetUserIDFromToken(token string) (string, error)
}

type service struct {
	config *config.Config
}

func New(config *config.Config) Interface {
	return &service{ config: config }
}


func (s *service) GenerateFromUserID(id string) (string, error) {
	claims := Claims{
		id,
		jwtLib.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			Issuer:    "GoPress",
		},
	}
	token := jwtLib.NewWithClaims(jwtLib.SigningMethodHS256, claims)

	mySecret := []byte(s.config.JWT.Secret)
	signedToken, err := token.SignedString(mySecret)
	if err != nil {
		return "", fmt.Errorf("error found while trying to sign the JWT: %v", err)
	}

	return signedToken, nil
}

func (s *service) GetUserIDFromToken(receivedToken string) (string, error) {
	token, err := jwtLib.Parse(receivedToken, func(token *jwtLib.Token) (interface{}, error) {
		method, ok := token.Method.(*jwtLib.SigningMethodHMAC)
		if !ok || method != jwtLib.SigningMethodHS256 {
			return "", fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		mySecret := []byte(s.config.JWT.Secret)
		return mySecret, nil
	})
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwtLib.MapClaims); ok && token.Valid {
		if id, ok := claims["id"].(string); ok {
			return id, nil
		}
	}

	return "", errors.New("invalid claims")
}