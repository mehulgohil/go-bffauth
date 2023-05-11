package controller

import (
	"encoding/json"
	"errors"
	"github.com/kataras/iris/v12"
	"github.com/mehulgohil/go-bffauth.git/config"
	"github.com/mehulgohil/go-bffauth.git/interfaces"
	"io"
	"net/http"
)

type WriterHandler struct {
	RedisClient interfaces.IRedisLayer
}

func (w *WriterHandler) WriterRedirect(ctx iris.Context) {
	userCookie := ctx.GetCookie("logged_id_email")
	if userCookie == "" {
		ctx.StopWithError(iris.StatusUnauthorized, errors.New("please make sure user is logged in"))
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest(ctx.Request().Method, config.EnvVariables.ShortifyWriterAPI, ctx.Request().Body)
	if err != nil {
		ctx.StopWithError(500, err)
		return
	}

	req.Header.Add("Authorization", "Bearer "+tokenMap[userCookie])
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		ctx.StopWithError(500, err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		ctx.StopWithError(500, err)
		return
	}

	var respBody map[string]interface{}
	err = json.Unmarshal(body, &respBody)
	if err != nil {
		ctx.StopWithError(500, err)
		return
	}

	ctx.StopWithJSON(res.StatusCode, respBody)
}
