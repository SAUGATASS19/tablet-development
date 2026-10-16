package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// In a real AWS environment, this secret is never hardcoded. 
// It is securely fetched from AWS Secrets Manager. We use this string just for local testing.
var jwtSecret = []byte("super_secret_mobile_app_key_change_me_later")

// GenerateTokenPair creates both the short-lived access token and long-lived refresh token.
func GenerateTokenPair(userID uint) (string, string, error) {
	// 1. Create the Access Token (Valid for 15 minutes)
	// We embed the user's database ID directly into the token payload (claims).
	accessTokenClaims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(15 * time.Minute).Unix(),
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessTokenString, err := accessToken.SignedString(jwtSecret)
	if err != nil {
		return "", "", err
	}

	// 2. Create the Refresh Token (Valid for 7 days)
	// This token lives much longer and is used to get a new Access Token without logging in again.
	refreshTokenClaims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(7 * 24 * time.Hour).Unix(),
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	refreshTokenString, err := refreshToken.SignedString(jwtSecret)
	if err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil
}