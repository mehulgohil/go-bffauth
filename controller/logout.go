package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
	"github.com/mehulgohil/go-bffauth.git/config"
	"github.com/mehulgohil/go-bffauth.git/interfaces"
	"net/http"
	"net/url"
)

type LogoutHandler struct {
	RedisClient interfaces.IRedisLayer
}

func (l *LogoutHandler) Logout(ctx iris.Context) {
	session := sessions.Get(ctx)

	session.Destroy()
	logoutUrl, err := url.Parse("https://" + config.EnvVariables.Auth0Domain + "/v2/logout")
	if err != nil {
		ctx.StopWithError(http.StatusInternalServerError, err)
		return
	}

	returnTo, err := url.Parse(config.EnvVariables.ShortifyFrontendDomain)
	if err != nil {
		ctx.StopWithError(http.StatusInternalServerError, err)
		return
	}

	parameters := url.Values{}
	parameters.Add("returnTo", returnTo.String())
	parameters.Add("client_id", config.EnvVariables.Auth0ClientID)
	logoutUrl.RawQuery = parameters.Encode()

	ctx.Redirect(logoutUrl.String(), http.StatusTemporaryRedirect)
}
