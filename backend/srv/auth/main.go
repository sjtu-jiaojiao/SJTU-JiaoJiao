package main

import (
	"context"
	"encoding/json"
	"errors"
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

func (a *srv) Auth(ctx context.Context, req *auth.AuthRequest, rsp *auth.AuthResponse) error {
	if req.Code != "" {
		rsp.Status = 1

		params := url.Values{}
		params.Add("grant_type", "authorization_code")
		params.Add("code", req.Code)
		params.Add("redirect_uri", utils.GetDeployHost("url")+"/"+
			utils.GetStringConfig("api_config", "version")+"/auth")
		params.Add("client_id", os.Getenv("JJ_CLIENTID"))
		params.Add("client_secret", os.Getenv("JJ_CLIENTSECRET"))
		resp, err := http.PostForm(utils.GetStringConfig("sys_config", "token_url"), params)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		id := idToken{}
		err = json.Unmarshal(body, &id)
		if err != nil {
			return err
		}
		if id.Error != "" {
			return errors.New(id.Error)
		}
		rsp.Token = id.IDToken
	} else {
		rsp.Status = 2
	}
	return nil
}

func main() {
	service := utils.InitMicroService("auth")
	auth.RegisterAuthHandler(service.Server(), new(srv))
	utils.RunMicroService(service)
}
