package handlers

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func MetadataGetConnectionInfo(c *gin.Context) {
	repoId := c.Param("id")
	if repoId == "" {
		c.JSON(400, "Invalid repository ID")
		return
	}

	connectionInfo, err := GetDatabaseConnectionModel(repoId)
	if err != nil {
		c.JSON(500, err)
		return
	}
	log.Debugln("Got database connection info", connectionInfo)
	c.JSON(200, connectionInfo)
}
