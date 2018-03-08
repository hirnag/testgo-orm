package dao

import (
	"database/sql"
)

func RawSelect(db *sql.DB) *sql.Rows {
	rows, err := db.Query("select * from users")
	if err != nil {
		panic(err)
	}
	return rows
}
