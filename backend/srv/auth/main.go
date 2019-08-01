package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	auth "jiaojiao/srv/auth/proto"
	"jiaojiao/utils"
	"net/http"
	"net/url"
	"os"
)

type srv struct{}

type idToken struct {
	IDToken string `json:"id_token"`
	Error   string `json:"error"`
}

/**
 * @api {rpc} /rpc Auth.Auth
 * @apiVersion 1.0.0
 * @apiGroup Service
 * @apiName Auth.Auth
 * @apiDescription Check OAuth code.
 *
 * @apiParam {string} code OAuth code.
 * @apiSuccess {int32} status -1 for invalid param <br> 1 for success <br> 2 for invalid code <br> 3 for frozen user
 * @apiSuccess {string} token verified token when status=1
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

		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{Transport: tr}
		resp, err := client.PostForm(utils.GetStringConfig("sys_config", "token_url"), params)
		if utils.LogContinue(err, utils.Error) {
			return err
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if utils.LogContinue(err, utils.Error) {
			return err
		}

		id := idToken{}
		err = json.Unmarshal(body, &id)
		if utils.LogContinue(err, utils.Error) {
			return err
		}

		if id.Error != "" { // invalid code
			rsp.Status = auth.AuthResponse_INVALID_CODE
		} else {
			t, err := utils.JWTVerify(id.IDToken, os.Getenv("JJ_CLIENTSECRET"))
			if err != nil {
				rsp.Status = auth.AuthResponse_INVALID_CODE
			} else {
				rsp.Status = auth.AuthResponse_SUCCESS
				rsp.Token = id.IDToken
				rsp.StudentID = utils.JWTParse(t, "code").(string)
				rsp.StudentName = utils.JWTParse(t, "name").(string)
			}
		}
	} else {
		rsp.Status = auth.AuthResponse_INVALID_PARAM
	}
	return nil
}

func main() {
	service := utils.InitMicroService("auth")
	utils.LogPanic(auth.RegisterAuthHandler(service.Server(), new(srv)))
	utils.RunMicroService(service)
}
