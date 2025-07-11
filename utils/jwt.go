// package utils

// import (
//     "github.com/golang-jwt/jwt/v5"
//     "os"
//     "time"
// )

// func GenerateToken(userID uint) (string, error) {
//     claims := jwt.MapClaims{
//         "userID": userID,
//         "exp":    time.Now().Add(time.Hour * 72).Unix(),
//     }

//     token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
//     return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
// }

package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secret = []byte(os.Getenv("JWT_SECRET"))

func GenerateToken(userID uint) (string, string, error) {
	accessClaims := jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(time.Minute * 15).Unix(),
	}
	refreshClaims := jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 24 * 7).Unix(),
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)

	accessStr, err := accessToken.SignedString(secret)
	if err != nil {
		return "", "", err
	}

	refreshStr, err := refreshToken.SignedString(secret)
	return accessStr, refreshStr, err
}

func ParseToken(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}
