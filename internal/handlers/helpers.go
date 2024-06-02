package handlers

import (
	"github.com/garnet-geo/garnet-geoapi/internal/db"
	log "github.com/sirupsen/logrus"
)

func GetDatabaseConnectionModel(repoId string) (ConnectionInfoModel, error) {
	row := db.Conn.QueryRow(db.Context(),
		"SELECT dolt_user, dolt_password, dolt_address, dolt_status "+
			"FROM repositories WHERE id = $1;", repoId)

	var connectionInfo ConnectionInfoModel
	err := row.Scan(&connectionInfo.User, &connectionInfo.Password, &connectionInfo.Address, &connectionInfo.Status)
	if err != nil {
		log.Debugln("Error getting database connection info", err)
		return connectionInfo, err
	}

	log.Debugf("Got database connection info %s:%s@%s with status %s",
		connectionInfo.User,
		connectionInfo.Password,
		connectionInfo.Address,
		connectionInfo.Status)

	connectionInfo.Database = "dolt"
	return connectionInfo, nil
}
