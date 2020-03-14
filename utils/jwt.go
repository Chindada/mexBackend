package utils

import (
	"time"

	"github.com/astaxie/beego"

	jwt "github.com/dgrijalva/jwt-go"
)

var key []byte

// GenToken GenToken
func GenToken(username string, interval int64) string {
	key = []byte(username)
	claims := &jwt.StandardClaims{
		NotBefore: int64(time.Now().Unix()),
		ExpiresAt: int64(time.Now().Unix() + interval),
		Issuer:    username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(key)
	if err != nil {
		LogError(err)
		return ""
	}
	return ss
}

// CheckToken CheckToken
func CheckToken(token string) (string, bool) {
	_, err := jwt.Parse(token, func(*jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		beego.Informational("Token Is Not Valid", err)
		return "", false
	}
	return string(key), true
}
