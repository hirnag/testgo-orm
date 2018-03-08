package dao

import (
	"github.com/jmoiron/sqlx"
	modXorm "../models-xorm"
)

func SqlxSelect(db sqlx.DB) []modXorm.Users {
	users := []modXorm.Users{}
	err := db.Select(&users, "select * from users")
	if err != nil { panic(err) }
	return users
}
