package migrations

import (
	"fmt"
	"go_framework/database"
)

func Post() {

	db := database.Connect()
	tableName := "Posts"
	// drop if exists
	qcheck := "DROP TABLE IF EXISTS %s"
	ExecQuery(db, fmt.Sprintf(qcheck, tableName))

	query := "CREATE TABLE %s (%s, %s, %s, %s, %s, %s)"

	value := []interface{}{
		tableName,
		"id bigint PRIMARY KEY NOT NULL AUTO_INCREMENT",
		"title varchar(100) UNIQUE",
		"content varchar(1000)",
		"image varchar(100)",
		"created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP",
		"updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP",
	}

	ExecQuery(db, fmt.Sprintf(query, value...))

	defer db.Close()
}
