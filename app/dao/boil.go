package dao

import (
	"github.com/volatiletech/sqlboiler/boil"
	"fmt"
	modBoil "../models-boil"
)

func BoilSelect(db boil.Executor) []*modBoil.User {
	boil.SetDB(db)
	users, err := modBoil.UsersG().All()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return users
}