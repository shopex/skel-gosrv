package model

import (
	_ "database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func ConnectDB(dbinfo string) (err error) {
	db, err = sqlx.Connect("mysql", dbinfo)
	return
}
