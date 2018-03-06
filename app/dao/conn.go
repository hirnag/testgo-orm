package dao

import (
	"github.com/volatiletech/sqlboiler/bdb/drivers"
	"database/sql"
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

func GetDataSourceName() string {
	return drivers.MySQLBuildQueryString(
		DbUser,DbPassword,DbName,DbHost,DbPort,DbSslmode)
}

func GetConnection() (*sql.DB, error) {
	return sql.Open(DriverName, GetDataSourceName())
}