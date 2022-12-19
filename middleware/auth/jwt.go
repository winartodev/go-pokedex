package auth

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/winartodev/go-pokedex/enum"
)

var jwtKey = []byte("supersecretkey")

// JWTClaim is struct represent of jwt.Claims
type JWTClaim struct {
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Role     enum.Role `json:"role"`
	jwt.StandardClaims
}

// GenerateJWT will generate token
func GenerateJWT(username string, email string, role enum.Role) (tokenString string, err error) {
	expirationTime := time.Now().Add(1 * time.Hour)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &JWTClaim{
		Username: username,
		Email:    email,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	},
	)

	return token.SignedString(jwtKey)
}

// ValidateToken will validate token
func ValidateToken(signedToken string) (claims *JWTClaim, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return claims, err
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return claims, err
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return claims, err
	}
	return claims, err
}
