package models

import (
	"database/sql"
)

// Tables for list of table
var Tables []interface{}

// LimitStatement for limit statement
var LimitStatement = "LIMIT ? OFFSET ?"

// MDL for model
type MDL struct {
	MysqlClient *sql.DB
}

// NewMDL create new model
func NewMDL(db *sql.DB) MDL {
	return MDL{
		MysqlClient: db,
	}
}

func init() {
	// logrus.Info("model init")
}

// SetMysqlClient set mysql db client
func (m *MDL) SetMysqlClient(db *sql.DB) {
	m.MysqlClient = db
}

// Close pg client connection
func (m *MDL) Close() {
	m.MysqlClient.Close()
}

// AddTable for append new schema
func (m *MDL) AddTable(table ...interface{}) {
	Tables = append(Tables, table...)
}
