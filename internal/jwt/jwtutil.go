package jwtutil

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

// GenerateToken generates a JWT token
func GenerateToken(userId uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userId,
		"exp": time.Now().Add(time.Hour * 24).Unix(), // 24-hour expiry
	})

	return token.SignedString(secretKey)
}

// ParseToken parses the JWT token and validates it
func ParseToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.NewValidationError("invalid signing method", jwt.ValidationErrorSignatureInvalid)
		}
		return secretKey, nil
	})
}

// ValidateToken checks if the token is valid and expired
func ValidateToken(tokenString string) (bool, error) {
	token, err := ParseToken(tokenString)
	if err != nil {
		return false, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		expirationTime := int64(claims["exp"].(float64))
		if time.Now().Unix() > expirationTime {
			return false, nil // Token is expired
		}
		return true, nil // Token is valid
	}

	return false, jwt.NewValidationError("invalid token", jwt.ValidationErrorMalformed)
}
