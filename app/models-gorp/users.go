package models_gorp

type User struct {
	ID        int        `db:"id, primarykey"`
	Name      string     `db:"name"`
}