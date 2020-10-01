package services

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

func CreateToken(userid uint) (string, error) {
	var err error
	//Creating Access Token
	// os.Setenv("ACCESS_SECRET", "TEST_SECRET") //this should be in an env file

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userid
	atClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(viper.GetString("app.secret")))
	if err != nil {
		return "", err
	}
	return token, nil
}
