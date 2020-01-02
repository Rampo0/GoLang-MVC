package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (db *sql.DB) {

	db, err := sql.Open(config().Driver, config().User+":"+config().Pass+"@/"+config().Name)

	if err != nil {
		panic(err.Error())
	}

	return
}
