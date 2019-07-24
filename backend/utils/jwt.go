package utils

import (
	"fmt"
	user "jiaojiao/srv/user/proto"
	"os"
	"time"

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

// JWTParse parse jwt token content
func JWTParse(token *jwt.Token, param string) interface{} {
	if claims, ok := token.Claims.(jwt.MapClaims); ok && claims[param] != nil {
		return claims[param]
	}
	return ""
}

// JWTSign sign a jwt token
func JWTSign(id int32, role user.UserInfo_Role) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   id,
		"role": role,
		"exp":  time.Now().Unix() + 1800,
	})

	t, err := token.SignedString([]byte(os.Getenv("JJ_JWTSECRET")))
	LogContinue(err, Warning)
	return t
}
