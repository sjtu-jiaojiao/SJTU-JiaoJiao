package main

import (
	"io/ioutil"
	"jiaojiao/utils"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router, rg := utils.CreateAPIGroup()
	rg.GET("/tag", getTag)
	return router
}

/**
 * @api {get} /tag GetTag
 * @apiVersion 1.0.0
 * @apiGroup Tag
 * @apiPermission user/admin
 * @apiName GetTag
 * @apiDescription Get AI tag
 *
 * @apiParam {string} text example text
 * @apiSuccess {list} tag tag list
 * @apiError (Error 500) TagServiceDown Tag service down
 */
func getTag(c *gin.Context) {
	type param struct {
		Text string `form:"text" binding:"required"`
	}
	var p param
	role := utils.GetRole(c)

	if !utils.LogContinue(c.ShouldBindQuery(&p), utils.Warning) {
		if !role.User && !role.Admin {
			c.AbortWithStatus(403)
			return
		}

		// TODO: based on consul
		u, err := url.Parse("http://localhost:5000")
		if err != nil {
			c.JSON(500, err)
			return
		}
		pm := url.Values{}
		pm.Set("text", p.Text)
		u.RawQuery = pm.Encode()
		resp, err := http.Get(u.String())
		if err != nil {
			c.JSON(500, err)
			return
		}
		if resp.StatusCode != http.StatusOK {
			c.AbortWithStatus(resp.StatusCode)
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			c.JSON(500, err)
			return
		}

		c.String(200, string(body))
	} else {
		c.AbortWithStatus(400)
	}
}

func main() {
	utils.RunWebService("tag", setupRouter())
}
