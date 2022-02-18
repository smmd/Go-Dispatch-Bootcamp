package client

import (
    "fmt"
    "time"

    "github.com/dgrijalva/jwt-go"
)

var signingKey = []byte("secretphrase")

type TokenGenerator struct{}

func NewTokenGenerator() TokenGenerator {
    return TokenGenerator{}
}

func (TokenGenerator) GenerateToken() (string, error) {
    token := jwt.New(jwt.SigningMethodHS256)

    claims := token.Claims.(jwt.MapClaims)

    claims["authorized"] = true
    claims["user"] = "Sag Test"
    claims["expiration"] = time.Now().Add(time.Minute * 30).Unix()

    tokenString, err := token.SignedString(signingKey)

    if err != nil {
        fmt.Errorf("ay an error has occured: %s", err.Error())
        return "", err
    }

    return tokenString, nil
}

//Middleware function
