package main

import (
	"database/sql"
	"./dao"
	"testing"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/jinzhu/gorm"
	"strconv"
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/go-gorp/gorp"
)

var db *sql.DB
var gormDb *gorm.DB
var xormDb *xorm.Engine
var gorpDb *gorp.DbMap

func init() {
	var err error
	// -- DB接続
	db, err = dao.GetConnection()
	if err != nil { panic(err) }
	// Boil
	boil.SetDB(db)
	// GORM
	//str := dao.DbUser + ":" + dao.DbPassword + "@tcp(" + dao.DbHost + ":" + strconv.Itoa(dao.DbPort) + ")/" + dao.DbName
	//gormDb, err = gorm.Open(dao.DriverName, str)
	gormDb, err = gorm.Open(dao.DriverName, dao.GetDataSourceName())
	if err != nil { panic(err) }
	// XORM
	xormDb, err = xorm.NewEngine(dao.DriverName, dao.GetDataSourceName())
	if err != nil { panic(err) }
	// GORP
	gorpDb = &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
}

func BenchmarkRawSelect(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dao.RawSelect(db)
	}
}
func BenchmarkBoilSelect(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dao.BoilSelect(boil.GetDB())
	}
}
func BenchmarkGormSelect(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dao.GormSelect(gormDb)
	}
}
func BenchmarkGormRawSelect(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dao.GormRawSelect(gormDb)
	}
}
func BenchmarkXormSelect(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dao.XormSelect(xormDb)
	}
}
func BenchmarkGorpSelect(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dao.GorpSelect(*gorpDb)
	}
}

// ------------------
//
//func TestResult(t *testing.T) {
//	users := dao.GormSelect(gormDb)
//	res := ""
//	for _, v := range users {
//		res += strconv.Itoa(v.ID)
//		res += ","
//		res += v.Name
//		res += "|"
//	}
//	fmt.Println(res)
//}