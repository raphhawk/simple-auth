package main

import (
	"net/http"
	jwtutil "simple-auth/internal/jwt"
	"simple-auth/internal/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

// SignUp handles user registration
func SignUp(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var user models.User
		if err := c.Bind(&user); err != nil {
			return c.JSON(http.StatusBadRequest, "Invalid input")
		}

		// Validate email and password
		validate := validator.New()
		if err := validate.Struct(user); err != nil {
			return c.JSON(http.StatusBadRequest, "Invalid data")
		}

		// Hash password
		if err := user.HashPassword(); err != nil {
			return c.JSON(http.StatusInternalServerError, "Error hashing password")
		}

		// Save user in DB
		if err := db.Create(&user).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, "Error saving user")
		}

		return c.JSON(http.StatusCreated, "User created successfully")
	}
}

// SignIn handles user login
func SignIn(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var credentials struct {
			Email    string `json:"email" validate:"required,email"`
			Password string `json:"password" validate:"required"`
		}

		if err := c.Bind(&credentials); err != nil {
			return c.JSON(http.StatusBadRequest, "Invalid input")
		}

		// Check if user exists
		var user models.User
		if err := db.Where("email = ?", credentials.Email).First(&user).Error; err != nil {
			return c.JSON(http.StatusUnauthorized, "Invalid credentials")
		}

		// Check password
		if !user.CheckPassword(credentials.Password) {
			return c.JSON(http.StatusUnauthorized, "Invalid credentials")
		}

		// Generate token
		token, err := jwtutil.GenerateToken(user.ID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "Error generating token")
		}

		return c.JSON(http.StatusOK, map[string]string{
			"token": token,
		})
	}
}

// RefreshToken handles token refreshing
func RefreshToken(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		if token == "" {
			return c.JSON(http.StatusUnauthorized, "Token missing")
		}

		// Validate token
		valid, err := jwtutil.ValidateToken(token)
		if err != nil || !valid {
			return c.JSON(http.StatusUnauthorized, "Invalid or expired token")
		}

		// Parse the token and refresh it
		claims, _ := jwtutil.ParseToken(token)
		userId := uint(claims.Claims.(jwt.MapClaims)["sub"].(float64))

		newToken, err := jwtutil.GenerateToken(userId)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "Error generating new token")
		}

		return c.JSON(http.StatusOK, map[string]string{
			"token": newToken,
		})
	}
}

// RevokeToken handles token revocation
func RevokeToken(c echo.Context) error {
	return c.JSON(http.StatusOK, "Token revoked (Handled by client-side)")
}
