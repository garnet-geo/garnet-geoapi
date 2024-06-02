package env

import "github.com/garnet-geo/garnet-geoapi/internal/consts"

func GetDatabaseUrl() string {
	return GetStringEnv(consts.DatabaseUrlEnv)
}
