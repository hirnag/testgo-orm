package dao

import (
	"github.com/volatiletech/sqlboiler/queries"
	modBoil "../models"
)

func Raw() {
	err := queries.Raw(db, "select * from users where id=$1", 5).Bind(&modBoil.UserSlice{})
	if err != nil {
		return
	}
}