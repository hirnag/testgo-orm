package dao

import (
	"github.com/go-xorm/xorm"
	modXorm "../models-xorm"
)

func XormSelect(db *xorm.Engine) []modXorm.Users {
	users := []modXorm.Users{}
	err := db.Find(&users)
	if err != nil { panic(err) }
	return users
}
