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

func DecodeToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signin method: %v", token.Header["alg"])
		}

		return []byte(SECRET), nil
	})

	if err != nil {
		return jwt.MapClaims{}, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return jwt.MapClaims{}, err
	}
}
