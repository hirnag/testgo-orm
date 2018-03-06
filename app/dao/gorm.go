package dao

import (
	"github.com/jinzhu/gorm"
	modGorm "../models-gorm"
)

func GormSelect(db *gorm.DB) []modGorm.User {
	users := []modGorm.User{}
	//db.Raw("select * from users").Scan(&users)
	db.Find(&users)
	return users
}

func GormRawSelect(db *gorm.DB) []modGorm.User {
	users := []modGorm.User{}
	db.Raw("select * from users").Scan(&users)
	return users
}
