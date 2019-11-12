package echokit

import (
	"database/sql"

	"github.com/alamyudi/echo-app/echokit/dbmanager"
	"github.com/alamyudi/echo-app/echokit/models"
)

// EchoKit structure
type EchoKit struct {
	DBManager *dbmanager.DBManager
}

// Init init EchoKit
func Init(db *sql.DB) *EchoKit {
	nk := new(EchoKit)

	dbManager := dbmanager.NewManager(
		models.NewMDL(db),
	)
	nk.DBManager = dbManager

	return nk
}
