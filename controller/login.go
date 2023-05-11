package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
	"github.com/mehulgohil/go-bffauth.git/authenticator"
	"github.com/mehulgohil/go-bffauth.git/config"
	"golang.org/x/oauth2"
	"net/http"
)

type LoginHandler struct {
	Auth *authenticator.Authenticator
}

func (l *LoginHandler) Login(ctx iris.Context) {
	// Save the state inside the session.
	session := sessions.Get(ctx)
	session.Set("state", state)
	ctx.Redirect(l.Auth.AuthCodeURL(state, oauth2.SetAuthURLParam("audience", config.EnvVariables.Auth0Audience)), http.StatusTemporaryRedirect)
}
