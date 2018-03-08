package dao

import (
	"github.com/volatiletech/sqlboiler/bdb/drivers"
	"database/sql"
	"github.com/go-xorm/xorm"
)

const (
	DriverName = "mysql"
	DbUser = "root"
	DbPassword = "password"
	DbName = "test"
	DbHost = "localhost"
	DbPort = 3306
	DbSslmode = "false"
)

var engine *xorm.Engine

func OpenConnection() {
	var err error
	engine, err = xorm.NewEngine(DriverName, GetDataSourceName())
	if err != nil { panic(err) }
}

func GetDataSourceName() string {
	return drivers.MySQLBuildQueryString(
		DbUser,DbPassword,DbName,DbHost,DbPort,DbSslmode)
}

func GetConnection() (*sql.DB, error) {
	return sql.Open(DriverName, GetDataSourceName())
}