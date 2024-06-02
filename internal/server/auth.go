package server

import (
	"encoding/json"
	"net/http"

	"github.com/garnet-geo/garnet-geoapi/internal/consts"
	"github.com/garnet-geo/garnet-geoapi/internal/env"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader(consts.AuthorizationHeader)
		res, err := useCheck(authHeader)
		if err != nil {
			log.Debugln("Err using check", err)
			ctx.JSON(500, "internal server error")
			ctx.Abort()
			return
		}

		log.Debugln("Got response", res.StatusCode)
		if res.StatusCode != 200 {
			log.Debugln("Non 200 response, unauthorized")
			ctx.JSON(401, "unauthorized")
			ctx.Abort()
			return
		}

		defer res.Body.Close()
		type UserIdResponse struct {
			UserId string `json:"user_id"`
		}
		var uidRes UserIdResponse
		if err := json.NewDecoder(res.Body).Decode(&uidRes); err != nil {
			log.Warningln("Err decoding response", err)
			ctx.JSON(500, "internal server error")
			ctx.Abort()
			return
		}

		log.Debugln("Got user id", uidRes.UserId, "setting to context")
		ctx.Set(consts.UserIDContextKey, uidRes.UserId)
	}
}

func useCheck(authHeader string) (*http.Response, error) {
	url := env.GetUserApiUrl() + "/check"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Debugln("Err creating request", err)
		return nil, err
	}
	req.Header.Add("Authorization", authHeader)
	log.Debugln("Sending auth request to", url, "with auth header", authHeader)

	client := &http.Client{}
	res, err := client.Do(req)
	return res, err
}
