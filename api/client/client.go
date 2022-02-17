package client

import (
    "fmt"
    "time"

    jwt "github.com/dgrijalva/jwt-go"
)

var signingKey = []byte("secretphrase")

func GenerateJWT() (string, error) {
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

func main() {
    fmt.Println("My Simple Client")

    tokenString, err := GenerateJWT()

    if err != nil {
        fmt.Println("Error generating token string")
    }

    fmt.Println(tokenString)
}
