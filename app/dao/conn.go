package dao

import (
	"database/sql"
	"github.com/volatiletech/sqlboiler/bdb/drivers"
	"github.com/volatiletech/sqlboiler/boil"
	modBoil "../models"
	"fmt"
)

const (
	DbUser = "root"
	DbPassword = "password"
	DbName = "test"
	DbHost = "localhost"
	DbPort = 3306
	DbSslmode = "false"
)

var db boil.Executor

func Connect() {
	// DB接続
	db, err := sql.Open("mysql", drivers.MySQLBuildQueryString(
		DbUser, DbPassword,DbName,DbHost,DbPort,DbSslmode))
	if err != nil {
		return
	}
	defer db.Close()

	boil.SetDB(db)
	users, err := modBoil.UsersG().All()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(modBoil.GetUsers(users))


}
