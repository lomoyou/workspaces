package util

import (

	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSercet []byte

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func GenerateToken(username, passwoed string) (string,error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3*time.Hour) //设置token有效时间

	claims := Claims{
		EncodeMD5(username),
		EncodeMD5(passwoed),
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer: "lomo-blog",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSercet)

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSercet, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}