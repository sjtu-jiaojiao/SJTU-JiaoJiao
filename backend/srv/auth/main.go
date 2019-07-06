package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	auth "jiaojiao/srv/auth/proto"
	"jiaojiao/utils"
	"net/http"
	"net/url"
	"os"

	"github.com/micro/go-micro/server"
)

type srv struct{}

type idToken struct {
	IDToken string `json:"id_token"`
	Error   string `json:"error"`
}

/**
 * @api {rpc} /rpc auth.Auth.Auth
 * @apiVersion 1.0.0
 * @apiGroup Service
 * @apiName auth.Auth.Auth
 * @apiDescription Check OAuth code.
 *
 * @apiParam {String} code OAuth code.
 * @apiSuccess {Number} status 1 for success, 2 for empty code, 3 for invalid code
 * @apiSuccess {String} token verified token when status=1
 * @apiError (Error 500) OAuthServerDown can't connect to OAuth server
 */
func (a *srv) Auth(ctx context.Context, req *auth.AuthRequest, rsp *auth.AuthResponse) error {
	if req.Code != "" {
		// post to oauth server
		params := url.Values{}
		params.Add("grant_type", "authorization_code")
		params.Add("code", req.Code)
		params.Add("redirect_uri", utils.GetDeployHost("url")+"/"+
			utils.GetStringConfig("api_config", "version")+"/auth")
		params.Add("client_id", os.Getenv("JJ_CLIENTID"))
		params.Add("client_secret", os.Getenv("JJ_CLIENTSECRET"))
		resp, err := http.PostForm(utils.GetStringConfig("sys_config", "token_url"), params)
		if utils.LogContinue(err, utils.Warning) {
			return err
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if utils.LogContinue(err, utils.Warning) {
			return err
		}

		id := idToken{}
		err = json.Unmarshal(body, &id)
		if utils.LogContinue(err, utils.Warning) {
			return err
		}

		if id.Error != "" { // invalid code
			rsp.Status = 3
		} else {
			_, err = utils.JWTVerify(id.IDToken, os.Getenv("JJ_CLIENTSECRET"))
			if err != nil {
				rsp.Status = 3
			} else {
				rsp.Status = 1
				rsp.Token = id.IDToken
			}
		}
	} else {
		rsp.Status = 2
	}
	return nil
}

func main() {
	utils.RunMicroService("auth", func(s server.Server, hdlr interface{},
		opts ...server.HandlerOption) error {
		return auth.RegisterAuthHandler(s, hdlr.(auth.AuthHandler), opts...)
	}, new(srv))
}
