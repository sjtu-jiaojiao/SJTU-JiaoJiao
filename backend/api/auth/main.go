package main

import (
	"jiaojiao/utils"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
)

type authCode struct {
	Code string `form:"code"`
}

/**
 * @api {get} /auth GetAuth
 * @apiVersion 1.0.0
 * @apiGroup Auth
 * @apiPermission none
 * @apiName GetAuth
 * @apiDescription Redirect to OAuth url.
 *
 * @apiParam {String} [code]  OAuth code callback, DO NOT call it by yourself.
 * @apiSuccess (No param - Redirect 301) {Redirect} url Redirect to OAuth url.
 * @apiSuccess (With param - Success 200) {Number} status 0 for success, 1 for invalid credential
 * @apiSuccess (With param - Success 200) {String} token jwt token, empty when status=1
 */
func getAuth(c *gin.Context) {
	var code authCode

	if !utils.LogContinue(c.BindQuery(&code), utils.Warning) {
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
			c.String(200, "ok")
		}
	}
}

func setupRouter() *gin.Engine {
	router, rg := utils.CreateAPIGroup()
	rg.GET("/auth", getAuth)
	return router
}

func main() {
	utils.RunWebService("auth", setupRouter())
}
