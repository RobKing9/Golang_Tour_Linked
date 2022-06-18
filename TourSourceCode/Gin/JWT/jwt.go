package main

import (
	"github.com/golang-jwt/jwt"
	"log"
	"time"
)

type Claims struct {
	UserName string
	jwt.StandardClaims
}

func main () {
	key := []byte("mykey")
	c := Claims{
		UserName: "RobKing",
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			ExpiresAt: time.Now().Add(7*24*time.Hour).Unix(),
			Issuer: "zxw",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	log.Println("token=", token)
	tokenstring, err := token.SignedString(key)
	if err != nil {
		log.Println("err=", err.Error())
	}
	log.Println("tokenstring=", tokenstring)

	tokenparse, err := jwt.ParseWithClaims(tokenstring, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil{
		log.Println("err=", err.Error())
	}
	log.Println("parsetoken=", tokenparse)
	log.Println("claim=", tokenparse.Claims.(*Claims))
}
