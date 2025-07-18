package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
	"strings"
	"time"
)

type CustomClaims struct {
	UserId string `json:"userId"`
	jwt.RegisteredClaims
}

func GenerateJWT(userId, secret string) (string, error) {
	claims := CustomClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "CombatIntel",
		},
	}

	secretBytes, err := os.ReadFile("private_key.pem")
	if err != nil {
		return "", err
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(secretBytes))
	if err != nil {
		return "", err
	}

	tempToken := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	token, err := tempToken.SignedString(privateKey)
	if err != nil {
		return "", err
	}

	return token, nil
}

func ValidateJWT(tokenString string) (string, error) {
	publicBytes, err := os.ReadFile("public_key.pem")
	if err != nil {
		return "", err
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	if err != nil {
		return "", err
	}

	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (any, error) {
		return publicKey, nil
	})
	if err != nil {
		return "", err
	}
	if !token.Valid {
		return "", errors.New("token is invalid")
	}

	return token.Claims.(*CustomClaims).UserId, nil
}

func GetApiKey(header http.Header) (string, error) {
	authHeader := header.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("Authorization header is missing")
	}
	token := strings.Split(authHeader, " ")[1]
	return token, nil
}
