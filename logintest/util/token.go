package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

//------------------------------------------------------------------------------
//以下为公共部分，只能改密钥和过期时间
//------------------------------------------------------------------------------
const (
	SECRETKEY  = "@sd=-!~#$D#G%B&*I^$E@(*_+@!~~sdf@fg#"
	EXPIRETIME = 3600 //单位s
)

func CreateToken(cla jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, cla)
	return token.SignedString([]byte(SECRETKEY))
}

func VerifyToken(token string, cla jwt.Claims) error {
	tk, err := jwt.ParseWithClaims(token, cla, func(t *jwt.Token) (interface{}, error) {
		return []byte(SECRETKEY), nil
	})
	if err != nil {
		return err
	}
	if !tk.Valid {
		return fmt.Errorf("invalid")
	}
	return nil
}

type CheckClaims struct {
	jwt.StandardClaims
}

func GetToken() string {
	cla := CheckClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + EXPIRETIME,
		},
	}
	tk, err := CreateToken(cla)
	if err != nil {
		return ""
	}
	return tk
}

