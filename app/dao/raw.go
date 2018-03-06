package dao

import (
	"fmt"
	"database/sql"
)

func RawSelect(db *sql.DB) {

	_, err := db.Query("select * from users")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

}