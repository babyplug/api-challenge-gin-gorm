package services

import (
	"errors"

	"github.com/babyplug/api-challenge-gin-gorm/dto"
	"golang.org/x/crypto/bcrypt"
)

func Login(credentials *dto.Credentials) (string, error) {
	user, err := findUserByUsername(credentials.Username)
	if err != nil {
		return "", errors.New("Username or password is not correct")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))
	if err != nil {
		return "", errors.New("Username or password is not correct")
	}

	token, err := CreateToken(user.ID)
	if err != nil {
		return "", errors.New("Server error please contact admin")
	}

	return token, nil
}
