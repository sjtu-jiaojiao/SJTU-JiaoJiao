//+build !test

package utils

import (
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func parseHeader(c *gin.Context) *jwt.Token {
	header := c.Request.Header.Get("Authorization")
	if !strings.HasPrefix(header, "Bearer ") {
		return nil
	}
	s := strings.TrimPrefix(header, "Bearer ")
	t, err := JWTVerify(s, os.Getenv("JJ_JWTSECRET"))
	if LogContinue(err, Warning) {
		return nil
	}
	return t
}

func CheckAdmin(c *gin.Context) bool {
	if CheckInTest() {
		if c.Request.Header.Get("Authorization") == "admin" {
			return true
		} else {
			return false
		}
	}
	t := parseHeader(c)
	if t != nil {
		return JWTParse(t, "role").(float64) == 2
	}
	return false
}

func CheckUserId(c *gin.Context, id int32) bool {
	if CheckInTest() {
		if c.Request.Header.Get("Authorization") == "valid_user" {
			return true
		} else {
			return false
		}
	}
	t := parseHeader(c)
	if t != nil {
		return JWTParse(t, "id").(float64) == float64(id)
	}
	return false
}
