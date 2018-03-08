package main

import (
	"database/sql"
	"./dao"
	"testing"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/jinzhu/gorm"
	"github.com/go-xorm/xorm"
	"github.com/go-gorp/gorp"
	"github.com/jmoiron/sqlx"
	modNatv "./models-native"
	//modGorm "./models-gorm"
	modXorm "./models-xorm"
	"reflect"
	"strconv"
	"fmt"
	"time"
)

// local variable for database object
var db *sql.DB
var gormDb *gorm.DB
var xormDb *xorm.Engine
var gorpDb *gorp.DbMap
var sqlxDb *sqlx.DB

// local variable for other
var usid int
var pdid int
var xormUsr []modXorm.Users
var xormBts []modXorm.Bigtable
var xormPds [][]modXorm.Products

func init() {
	var err error
	// -- DB接続
	db, err = dao.GetConnection()
	if err != nil { panic(err) }
	// Boil
	boil.SetDB(db)
	// GORM
	gormDb, err = gorm.Open(dao.DriverName, dao.GetDataSourceName())
	if err != nil { panic(err) }
	// GORP
	gorpDb = &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	// XORM
	xormDb, err = xorm.NewEngine(dao.DriverName, dao.GetDataSourceName())
	if err != nil { panic(err) }
	// Sqlx
	sqlxDb, err = sqlx.Connect(dao.DriverName, dao.GetDataSourceName())
	if err != nil { panic(err) }

	// get latest user id and set
	users := []modNatv.User{}
	err = sqlxDb.Select(&users, "select id from users order by id desc limit 1")
	if err != nil { panic(err) }
	usid = users[0].ID

	// delete extra record
	deleteProducts()

	// set struct for xorm
	setStruct4Xorm()

}
// -- testing init
func TestInit(t *testing.T) {
}

// -- select
func BenchmarkRawSelectOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dao.RawSelect(db)
	}
}
func BenchmarkGormSelectOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dao.GormSelect(gormDb)
	}
}
func BenchmarkGormRawSelectOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dao.GormRawSelect(gormDb)
	}
}
func BenchmarkGorpSelectOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dao.GorpSelect(*gorpDb)
	}
}
func BenchmarkBoilSelectOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dao.BoilSelect(boil.GetDB())
	}
}
func BenchmarkXormSelectOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dao.XormSelect(xormDb)
	}
}
func BenchmarkXormRawSelectOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dao.XormRawSelect(xormDb)
	}
}
func BenchmarkSqlxSelectOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dao.SqlxSelect(*sqlxDb)
	}
}

// -- insert
var x1 int
func BenchmarkXormInsertOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dao.XormInsertOne(xormDb, xormUsr[x1])
		x1++
	}
}
var x2 int
func BenchmarkXormInsertLarge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dao.XormInsertLarge(xormDb, xormBts[x2])
		x2++
	}
}
var x3 int
func BenchmarkXormInsertBulk(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dao.XormInsertBulk(xormDb, xormPds[x3])
		x3++
	}
}

// local function ------------------
func setStruct4Xorm() {
	// set struct for xorm
	pdid = 1
	for x := 0; x < 10000; x++ {
		// create users
		usid++
		user := modXorm.Users{}
		user.Id = usid
		user.Name = "user" + strconv.Itoa(usid)
		xormUsr = append(xormUsr, user)

		// create bigtables
		bt := modXorm.Bigtable{}
		v := reflect.ValueOf(&bt).Elem()
		v.CanSet()
		v.Field(0).SetInt(int64(x))
		for i := 1; i < 100; i++ {
			v.Field(i).SetInt(int64(i))
		}
		xormBts = append(xormBts, bt)
	}
	for x := 0; x < 50; x++ {
		// create products list
		pds := []modXorm.Products{}
		term := pdid + 10000 // 一度にinsertするレコード数
		for i := pdid; i < term; i++ {
			pd := modXorm.Products{}
			pd.Id = pdid
			pd.ProductId = "P" + fmt.Sprintf("%07d", pdid)
			pd.ProductName = "name_" + strconv.Itoa(pdid)
			pd.Price = strconv.Itoa(pdid * (10 + 1/10))
			pd.RegisterTime = time.Now().Local()
			pds = append(pds, pd)
			pdid = i + 1
		}
		xormPds = append(xormPds, pds)
	}
}

func deleteProducts() {
	db.Exec("truncate table products")
}
//func TestRawResult(t *testing.T) {
//	rows := dao.RawSelect(db)
//	defer rows.Close()
//	users := []modNatv.User{}
//	for rows.Next() {
//		user := modNatv.User{}
//		rows.Scan(&user.ID, &user.Name)
//		users = append(users, user)
//	}
//	res := ""
//	for _, v := range users {
//		res += strconv.Itoa(v.ID)
//		res += ","
//		res += v.Name
//		res += "|"
//	}
//	fmt.Println(res)
//}