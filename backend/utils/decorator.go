//+build !test

package utils

import (
	user "jiaojiao/srv/user/proto"
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

func GetRoleID(c *gin.Context, id int32) user.UserInfo_Role {
	ret := user.UserInfo_GUEST
	if CheckInTest() {
		h := c.Request.Header.Get("Authorization")
		switch h {
		case "user":
			ret = user.UserInfo_USER
		case "self":
			ret = user.UserInfo_SELF
		case "admin":
			ret = user.UserInfo_ADMIN
		}
	} else {
		t := parseHeader(c)
		if t != nil {
			ret = user.UserInfo_Role(JWTParse(t, "role").(float64))
			if id != 0 && JWTParse(t, "id").(float64) == float64(id) {
				ret = user.UserInfo_SELF
			}
		}
	}
	return ret
}

func GetRole(c *gin.Context) user.UserInfo_Role {
	return GetRoleID(c, 0)
}
