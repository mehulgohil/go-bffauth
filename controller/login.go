package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/mehulgohil/go-bffauth.git/authenticator"
	"github.com/mehulgohil/go-bffauth.git/config"
	"golang.org/x/oauth2"
	"net/http"
)

type LoginHandler struct {
	Auth *authenticator.Authenticator
}

func (l *LoginHandler) Login(ctx iris.Context) {
	ctx.Redirect(l.Auth.AuthCodeURL(state, oauth2.SetAuthURLParam("audience", config.EnvVariables.Auth0Audience)), http.StatusTemporaryRedirect)
}
