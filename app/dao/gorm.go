package dao

import (
	"github.com/jinzhu/gorm"
	modGorm "../models-gorm"
)

func GormSelect(db *gorm.DB) []modGorm.User {
	users := []modGorm.User{}
	db.Find(&users)
	return users
}

func GormRawSelect(db *gorm.DB) []modGorm.User {
	users := []modGorm.User{}
	db.Raw("select * from users").Scan(&users)
	return users
}

func GormInsertOne(db *gorm.DB, user modGorm.User) {
	db.Create(user)
}

func GormInserts(db *gorm.DB, bts []modGorm.Bigtable) {
	db.Create(bts)
}
