package middleware

import (
	"errors"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
)

// IsAuthenticated is a middleware that checks if
// the user has already been authenticated previously.
func IsAuthenticated(ctx iris.Context) {
	session := sessions.Get(ctx)
	profileToken := session.Get("profile")
	if profileToken == nil {
		ctx.StopWithError(iris.StatusUnauthorized, errors.New("please make sure user is logged in"))
	} else {
		ctx.Next()
	}
}
