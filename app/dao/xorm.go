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
func XormRawSelect(db *xorm.Engine) []map[string][]byte {
	rows, err := db.Query("select * from users")
	if err != nil { panic(err) }
	return rows
}
func XormInsertOne(db *xorm.Engine, user modXorm.Users) int64 {
	affected, err := db.InsertOne(user)
	if err != nil { panic(err) }
	return affected
}
func XormInsertLarge(db *xorm.Engine, bigtable modXorm.Bigtable) int64 {
	affected, err := db.InsertOne(bigtable)
	if err != nil { panic(err) }
	return affected
}
func XormInsertBulk(db *xorm.Engine, products []modXorm.Products) int64 {
	affected, err := db.Insert(products)
	if err != nil { panic(err) }
	return affected
}
