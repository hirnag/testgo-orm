package dao

import (
	"github.com/go-gorp/gorp"
	modGorp "../models-gorp"
)

func GorpSelect(db gorp.DbMap) []interface{} {
	users, err := db.Select(modGorp.User{}, "select * from users")
	if err != nil { panic(err) }
	return users
}
