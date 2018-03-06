package models_xorm

type Users struct {
	ID   int    `json:"id" xorm:"'id'"`
	Name string `json:"name" xorm:"'name'"`
}
