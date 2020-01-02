package migrations

import (
	"database/sql"
	"fmt"
)

func ExecQuery(db *sql.DB, s string) {
	res, err := db.Exec(s)

	if err != nil {
		panic(err.Error())
	}

	if res != nil {
		fmt.Printf("Query Execution %s Success !! \n", s)
	}
}
