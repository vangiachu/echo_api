package middleware

import (
    "github.com/golang-jwt/jwt/v5"
    "github.com/labstack/echo/v4"
    "net/http"
    "os"
)

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        tokenString := c.Request().Header.Get("Authorization")
        if tokenString == "" {
            return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Missing token"})
        }

        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            return []byte(os.Getenv("JWT_SECRET")), nil
        })

        if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
            c.Set("userID", uint(claims["userID"].(float64)))
            return next(c)
        } else {
            return c.JSON(http.StatusUnauthorized, echo.Map{"error": err.Error()})
        }
    }
}
