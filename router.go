package main

import (
	"github.com/gorilla/securecookie"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/accesslog"
	"github.com/kataras/iris/v12/sessions"
	"github.com/mehulgohil/go-bffauth.git/authenticator"
	"github.com/mehulgohil/go-bffauth.git/config"
	"github.com/mehulgohil/go-bffauth.git/controller"
	"github.com/mehulgohil/go-bffauth.git/interfaces"
	"github.com/mehulgohil/go-bffauth.git/middleware"
	"io"
	"os"
	"sync"
)

var (
	irisRouter             *router
	routerOnce             sync.Once
	cookieNameForSessionID = "mycookiesessionnameid"
	sess                   = sessions.New(sessions.Config{Cookie: cookieNameForSessionID})
)

type IRouter interface {
	InitRouter(auth *authenticator.Authenticator, redis interfaces.IRedisLayer) *iris.Application
}

type router struct{}

func (router *router) InitRouter(auth *authenticator.Authenticator, redis interfaces.IRedisLayer) *iris.Application {
	app := iris.New()
	ac := makeAccessLog()
	app.UseRouter(ac.Handler)
	app.Use(sess.Handler())
	app.Use(useSecureCookies())

	// Our custom CORS middleware.
	crs := func(ctx iris.Context) {
		ctx.Header("Access-Control-Allow-Origin", config.EnvVariables.FrontendURL)
		ctx.Header("Access-Control-Allow-Credentials", "true")

		if ctx.Method() == iris.MethodOptions {
			ctx.Header("Access-Control-Methods",
				"POST, PUT, PATCH, GET, DELETE")

			ctx.Header("Access-Control-Allow-Headers",
				"Access-Control-Allow-Origin,Content-Type,Authorization")

			ctx.Header("Access-Control-Max-Age",
				"86400")

			ctx.StatusCode(iris.StatusNoContent)
			return
		}

		ctx.Next()
	}

	app.UseRouter(crs)

	loginHandler := controller.LoginHandler{Auth: auth}
	callbackHandler := controller.CallbackHandler{Auth: auth, RedisClient: redis}
	logoutHandler := controller.LogoutHandler{RedisClient: redis}
	backendApiHandler := controller.BackendApiHandler{RedisClient: redis}
	middlewareHandler := middleware.MiddlewareHandler{RedisClient: redis}

	app.Get("/login", loginHandler.Login)
	app.Get("/callback", callbackHandler.Callback)
	app.Get("/logout", logoutHandler.Logout)

	// Backend Api
	app.Post("/shorten", middlewareHandler.IsAuthenticated, backendApiHandler.WriterRedirect)

	return app
}

func Router() IRouter {
	if irisRouter == nil {
		routerOnce.Do(func() {
			irisRouter = &router{}
		})
	}
	return irisRouter
}

// This helps to log the request and its metadata
func makeAccessLog() *accesslog.AccessLog {
	ac := accesslog.New(io.MultiWriter(os.Stdout))
	ac.Delim = ' '
	ac.ResponseBody = false
	ac.RequestBody = false
	ac.BytesReceived = true
	ac.BytesSent = true

	return ac
}

func useSecureCookies() iris.Handler {
	var (
		hashKey  = securecookie.GenerateRandomKey(64)
		blockKey = securecookie.GenerateRandomKey(32)

		s = securecookie.New(hashKey, blockKey)
	)

	return func(ctx iris.Context) {
		ctx.AddCookieOptions(iris.CookieEncoding(s))
		ctx.Next()
	}
}
