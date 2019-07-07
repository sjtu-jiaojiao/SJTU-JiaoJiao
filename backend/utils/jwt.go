package utils

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

// JWTVerify verify the JWT token and return parsed object
func JWTVerify(token string, secret string) (*jwt.Token, error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if t != nil && t.Valid {
		return t, err
	}
	return nil, err
}

func JWTParse(token *jwt.Token, param string) string {
	if !token.Valid {
		return ""
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && claims[param] != nil {
		return claims[param].(string)
	}
	return ""
}
