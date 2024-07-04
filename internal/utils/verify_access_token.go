package utils

import (
	"errors"
	"os"

	"github.com/golang-jwt/jwt"
)

type CreateAccessTokenUtil struct{}

func NewCreateAccessTokenUtil() *CreateAccessTokenUtil {
	return &CreateAccessTokenUtil{}
}

func (b *CreateAccessTokenUtil) Validate(accessToken string) error {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("there was an error in parsing")
		}
		return os.Getenv("SECRET_JWT"), nil
	})

	if err != nil {
		return err
	}

	if token == nil {
		return errors.New("invalid token was provided")
	}

	return nil
}
