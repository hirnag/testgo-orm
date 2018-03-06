package models_gorp

type User struct {
	Id        int        `db:"id, primarykey"`
	Name      string     `db:"name"`
}