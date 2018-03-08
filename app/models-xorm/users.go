package model

type Users struct {
	Id   int    `xorm:"not null pk INT(11)"`
	Name string `xorm:"VARCHAR(255)"`
}
