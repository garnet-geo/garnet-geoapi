package server

import (
	"github.com/garnet-geo/garnet-geoapi/internal/consts"
	"github.com/garnet-geo/garnet-geoapi/internal/db"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func RepositoryAccessCheckMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Debugln("Checking access for", ctx.Request.URL.Path)

		userId := ctx.MustGet(consts.UserIDContextKey).(string)
		log.Debugln("Got user id", userId)

		repositoryId := ctx.Param("id")
		if repositoryId == "" {
			log.Debugln("Not repository id specified")
			ctx.JSON(400, "Invalid repository id")
			ctx.Abort()
			return
		}

		row := db.Conn.QueryRow(db.Context(),
			"SELECT u.id FROM ( "+
				"SELECT d.id FROM domains d "+
				"INNER JOIN repositories r ON d.id = r.\"domain\" "+
				"WHERE r.id = $1 "+
				") d INNER JOIN users u ON d.id = u.\"domain\";",
			repositoryId)
		var repoOwnerId string
		err := row.Scan(&repoOwnerId)
		if err != nil {
			log.Debugln("Error getting repo owner", err)
			ctx.JSON(500, "internal server error")
			ctx.Abort()
			return
		}

		if repoOwnerId != userId {
			log.Debugln("User", userId, "is not owner of repository", repositoryId)
			ctx.JSON(403, "forbidden")
			ctx.Abort()
			return
		}
	}
}
