package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
	"github.com/mehulgohil/go-bffauth.git/authenticator"
	"github.com/mehulgohil/go-bffauth.git/config"
	"github.com/mehulgohil/go-bffauth.git/interfaces"
	"net/http"
)

type CallbackHandler struct {
	Auth        *authenticator.Authenticator
	RedisClient interfaces.IRedisLayer
}

func (c *CallbackHandler) Callback(ctx iris.Context) {
	session := sessions.Get(ctx)
	if ctx.URLParam("state") != state {
		ctx.StopWithJSON(http.StatusBadRequest, "Invalid state parameter.")
		return
	}

	// Exchange an authorization code for a token.
	token, err := c.Auth.Exchange(ctx.Request().Context(), ctx.URLParam("code"))
	if err != nil {
		ctx.StopWithJSON(http.StatusUnauthorized, "Failed to convert an authorization code into a token.")
		return
	}

	idToken, err := c.Auth.VerifyIDToken(ctx.Request().Context(), token)
	if err != nil {
		ctx.StopWithJSON(http.StatusInternalServerError, "Failed to verify ID Token.")
		return
	}

	var profile map[string]interface{}
	if err := idToken.Claims(&profile); err != nil {
		ctx.StopWithError(http.StatusInternalServerError, err)
		return
	}

	tokenMap[profile["email"].(string)] = token.AccessToken

	session.Set("profile", profile)

	ctx.SetCookieKV("logged_id_email", profile["email"].(string), iris.CookieHTTPOnly(false))

	// Redirect to logged in page.
	ctx.Redirect(config.EnvVariables.ShortifyFrontendDomain, http.StatusTemporaryRedirect)
}