package server

import (
	"fmt"

	"github.com/garnet-geo/garnet-geoapi/internal/env"
	"github.com/garnet-geo/garnet-geoapi/internal/handlers"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func InitServer() {
	router := gin.Default()

	reposRouter := router.Group("/repository/:id")

	authRouter := reposRouter.Group("")
	authRouter.Use(AuthMiddleware())

	accessCheckRouter := authRouter.Group("")
	accessCheckRouter.Use(RepositoryAccessCheckMiddleWare())

	accessCheckRouter.GET("/connection_info", handlers.MetadataGetConnectionInfo)

	log.Debugln("Created gin routing")

	port := fmt.Sprint(env.GetServerHttpPort())

	log.Debugln("Port from environment: " + port)
	router.Run(":" + port)

	log.Info("Server started on port " + port)
}
