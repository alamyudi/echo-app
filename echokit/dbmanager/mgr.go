package dbmanager

import "github.com/alamyudi/echo-app/echokit/models"

// DBManager for business layers
type DBManager struct {
	MDL models.MDL
}

// NewManager init manager
func NewManager(mdl models.MDL) *DBManager {
	mng := new(DBManager)
	mng.MDL = mdl
	return mng
}
