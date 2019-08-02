package main

import (
	"context"
	"jiaojiao/utils"
	"net/url"
	"os"

	mockauth "jiaojiao/srv/auth/mock"
	auth "jiaojiao/srv/auth/proto"
	mockuser "jiaojiao/srv/user/mock"
	user "jiaojiao/srv/user/proto"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/client"
)

func setupRouter() *gin.Engine {
	router, rg := utils.CreateAPIGroup()
	rg.GET("/auth", getAuth)
	return router
}

type authCode struct {
	Code string `form:"code"`
}

/**
 * @apiDefine InvalidParam
 * @apiError (Error 400) InvalidParam Invalid param
 */

/**
 * @api {get} /auth GetAuth
 * @apiVersion 1.0.0
 * @apiGroup Auth
 * @apiPermission none
 * @apiName GetAuth
 * @apiDescription Redirect to OAuth url
 *
 * @apiParam {string} [code]  OAuth code callback, DO NOT call it by yourself
 * @apiSuccess (No param - Redirect 301) {Redirect} url Redirect to OAuth url
 * @apiSuccess (With param - Success 200) {--} Response see [Auth service](#api-Service-Auth_Auth)
 * @apiUse InvalidParam
 * @apiError (Error 500) AuthServiceDown Auth service down
 */
func getAuth(c *gin.Context) {
	var code authCode

	if !utils.LogContinue(c.ShouldBindQuery(&code), utils.Warning) {
		if code.Code == "" { // no param
			baseURL, err := url.Parse(utils.GetStringConfig("sys_config", "auth_url"))
			utils.LogPanic(err)

			params := url.Values{}
			params.Add("response_type", "code")
			params.Add("scope", "basic")
			params.Add("client_id", os.Getenv("JJ_CLIENTID"))
			params.Add("redirect_uri", utils.GetDeployHost("url")+"/"+
				utils.GetStringConfig("api_config", "version")+"/auth")

			// Add Query Parameters to the URL
			baseURL.RawQuery = params.Encode() // Escape Query Parameters

			c.Redirect(301, baseURL.String())
		} else { // with param
			srv := utils.CallMicroService("auth", func(name string, c client.Client) interface{} { return auth.NewAuthService(name, c) },
				func() interface{} { return mockauth.NewAuthService() }).(auth.AuthService)
			rsp, err := srv.Auth(context.TODO(), &auth.AuthRequest{
				Code: code.Code,
			})
			if utils.LogContinue(err, utils.Error) {
				c.JSON(500, err)
				return
			}

			if rsp.Status == auth.AuthResponse_SUCCESS {
				srv2 := utils.CallMicroService("user", func(name string, c client.Client) interface{} { return user.NewUserService(name, c) },
					func() interface{} { return mockuser.NewUserService() }).(user.UserService)
				rsp2, err := srv2.Create(context.TODO(), &user.UserCreateRequest{
					StudentID:   rsp.StudentID,
					StudentName: rsp.StudentName,
				})
				if utils.LogContinue(err, utils.Error) {
					c.JSON(500, err)
					return
				}

				// sign token
				if rsp2.User.Status == user.UserInfo_NORMAL {
					if rsp2.User.Role == user.UserInfo_USER {
						c.JSON(200, gin.H{
							"status": auth.AuthResponse_SUCCESS,
							"token":  utils.JWTSign(rsp2.User.UserID, user.UserInfo_USER),
						})
					} else {
						c.JSON(200, gin.H{
							"status": auth.AuthResponse_SUCCESS,
							"token":  utils.JWTSign(rsp2.User.UserID, user.UserInfo_ADMIN),
						})
					}
				} else {
					c.JSON(200, gin.H{
						"status": auth.AuthResponse_FROZEN_USER,
					})
				}
			} else {
				c.JSON(200, gin.H{
					"status": rsp.Status,
				})
			}
		}
	} else {
		c.AbortWithStatus(400)
	}
}

func main() {
	utils.RunWebService("auth", setupRouter())
}
