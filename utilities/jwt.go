package utilities

import (
	"errors"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

var SECRET string = os.Getenv("JWT_SECRET")

func GenerateToken(claims jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(SECRET))

	if err != nil {
		fmt.Println(err.Error())
		return "", errors.New("failed to sign token")
	}

	return tokenString, nil
}
