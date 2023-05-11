package controller

import (
	"errors"
	"github.com/kataras/iris/v12"
	"github.com/mehulgohil/go-bffauth.git/config"
	"github.com/mehulgohil/go-bffauth.git/interfaces"
	"net/http"
	"net/url"
)

type LogoutHandler struct {
	RedisClient interfaces.IRedisLayer
}

func (l *LogoutHandler) Logout(ctx iris.Context) {
	userCookie := ctx.GetCookie("logged_id_email")
	if userCookie == "" {
		ctx.StopWithError(iris.StatusUnauthorized, errors.New("please make sure user is logged in"))
		return
	}

	// delete token key
	err := l.RedisClient.DeleteKey(userCookie + "_token")
	if err != nil {
		ctx.StopWithError(http.StatusInternalServerError, err)
		return
	}

	// delete profile key
	err = l.RedisClient.DeleteKey(userCookie + "_profile")
	if err != nil {
		ctx.StopWithError(http.StatusInternalServerError, err)
		return
	}

	logoutUrl, err := url.Parse("https://" + config.EnvVariables.Auth0Domain + "/v2/logout")
	if err != nil {
		ctx.StopWithError(http.StatusInternalServerError, err)
		return
	}

	returnTo, err := url.Parse(config.EnvVariables.FrontendURL)
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
