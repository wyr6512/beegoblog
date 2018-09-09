package util

import (
	jwt "github.com/dgrijalva/jwt-go"

	"fmt"
	"time"
)

var (
	key []byte = []byte("wyr6512@163.com")
)

// 产生json web token
func GenToken() string {
	claims := &jwt.StandardClaims{
		NotBefore: int64(time.Now().Unix()),
		ExpiresAt: int64(time.Now().Unix() + 1000),
		Issuer:    "wyr6512",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	str, err := token.SignedString(key)
	if err != nil {
		fmt.Println("generate token failed.", err)
		return ""
	}
	return str
}

// 校验token是否有效
func CheckToken(token string) bool {
	_, err := jwt.Parse(token, func(*jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		fmt.Println("parase with claims failed.", err)
		return false
	}
	return true
}
