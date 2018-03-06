package dao

import (
	"github.com/volatiletech/sqlboiler/boil"
	"fmt"
	modBoil "../models-boil"
)

func BoilSelect(db boil.Executor) {

	boil.SetDB(db)
	_, err := modBoil.UsersG().All()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

}