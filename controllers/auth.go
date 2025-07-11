// package controllers

// import (
//     "echo_api/models"
//     "echo_api/utils"
//     "echo_api/database"
//     "net/http"

//     "github.com/labstack/echo/v4"
//     "golang.org/x/crypto/bcrypt"
// )

// func Register(c echo.Context) error {
//     var user models.User
//     if err := c.Bind(&user); err != nil {
//         return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
//     }
//     hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
//     user.Password = string(hashedPassword)

//     result := database.DB.Create(&user)
//     if result.Error != nil {
//         return c.JSON(http.StatusInternalServerError, echo.Map{"error": result.Error.Error()})
//     }

//     return c.JSON(http.StatusCreated, user)
// }

// func Login(c echo.Context) error {
//     var input models.User
//     var user models.User
//     if err := c.Bind(&input); err != nil {
//         return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
//     }

//     database.DB.Where("username = ?", input.Username).First(&user)
//     if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
//         return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid credentials"})
//     }

//     token, err := utils.GenerateToken(user.ID)
//     if err != nil {
//         return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Could not create token"})
//     }

//     return c.JSON(http.StatusOK, echo.Map{"token": token})
// }

package controllers

import (
	"echo_api/database"
	"echo_api/models"
	"echo_api/utils"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func Register(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	result := database.DB.Create(&user)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": result.Error.Error()})
	}

	return c.JSON(http.StatusCreated, user)
}

func Login(c echo.Context) error {
	var input models.User
	var user models.User
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	database.DB.Where("username = ?", input.Username).First(&user)
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid credentials"})
	}

	accessToken, refreshToken, err := utils.GenerateToken(user.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Could not create token"})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

func RefreshToken(c echo.Context) error {
	var input struct {
		RefreshToken string `json:"refresh_token"`
	}
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	claims, err := utils.ParseToken(input.RefreshToken)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid refresh token"})
	}

	userID := uint(claims["userID"].(float64))
	accessToken, refreshToken, err := utils.GenerateToken(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to generate token"})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}
