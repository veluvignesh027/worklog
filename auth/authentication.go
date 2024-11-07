package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey []byte

func InitAuthModule() {
	secretKey = []byte(os.Getenv("SECRET_KEY_JWT"))
}

func CreateToken(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email": email,
			"exp":   time.Now().Add(time.Hour * 1).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (string, error) {
	// Parse the JWT token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	// Check for parsing errors
	if err != nil {
		return "", err
	}

	// Check if the token is valid
	if !token.Valid {
		return "", fmt.Errorf("invalid token")
	}

	// Extract the email from the token claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("failed to extract claims")
	}
	email, ok := claims["email"].(string)
	if !ok {
		return "", fmt.Errorf("email claim not found or invalid")
	}

	// Return the email and nil error if everything is successful
	return email, nil
}
