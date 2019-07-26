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

// Role is user role struct
type Role struct {
	Guest bool
	User  bool
	Self  bool
	Admin bool
}

// GetRoleID get role and check id to judge self permission
func GetRoleID(c *gin.Context, id int32) Role {
	ret := Role{Guest: true}
	if CheckInTest() {
		h := c.Request.Header.Get("Authorization")
		switch h {
		case "user":
			ret.User = true
		case "self":
			ret.User = true
			ret.Self = true
		case "admin":
			ret.Admin = true
		}
	} else {
		t := parseHeader(c)
		if t != nil {
			switch user.UserInfo_Role(JWTParse(t, "role").(float64)) {
			case user.UserInfo_USER:
				ret.User = true
			case user.UserInfo_ADMIN:
				ret.Admin = true
			}
			if id != 0 && JWTParse(t, "id").(float64) == float64(id) {
				ret.Self = true
			}
		}
	}
	return ret
}

// GetRole get role without self permission
func GetRole(c *gin.Context) Role {
	return GetRoleID(c, 0)
}
