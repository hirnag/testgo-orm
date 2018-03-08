package dao

import (
	"github.com/go-gorp/gorp"
	modNatv "../models-native"
)

func GorpSelect(db gorp.DbMap) []interface{} {
	users, err := db.Select(modNatv.User{}, "select * from users")
	if err != nil { panic(err) }
	return users
}
