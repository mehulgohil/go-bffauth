package middleware

import (
	"errors"
	"github.com/kataras/iris/v12"
	"github.com/mehulgohil/go-bffauth.git/interfaces"
)

type MiddlewareHandler struct {
	RedisClient interfaces.IRedisLayer
}

// IsAuthenticated is a middleware that checks if
// the user has already been authenticated previously.
func (m *MiddlewareHandler) IsAuthenticated(ctx iris.Context) {
	userCookie := ctx.GetCookie("logged_id_email")
	if userCookie == "" {
		ctx.StopWithError(iris.StatusUnauthorized, errors.New("please make sure user is logged in"))
		return
	}

	value, err := m.RedisClient.HGetKeyValue(userCookie + "_profile")
	if err != nil || value == nil {
		ctx.StopWithError(iris.StatusUnauthorized, errors.New("please make sure user is logged in"))
		return
	}

	ctx.SetUser(value)
	ctx.Next()
}
