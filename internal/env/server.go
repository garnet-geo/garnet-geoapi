package env

import "github.com/garnet-geo/garnet-geoapi/internal/consts"

func GetServerHttpPort() int {
	return GetIntegerEnv(consts.ServerHttpPortEnv)
}

func GetUserApiUrl() string {
	return GetStringEnv(consts.UserApiUrlEnv)
}
