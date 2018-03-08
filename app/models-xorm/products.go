package model

import (
	"time"
)

type Products struct {
	Id           int       `xorm:"not null pk autoincr INT(11)"`
	ProductId    string    `xorm:"not null unique CHAR(8)"`
	ProductName  string    `xorm:"not null VARCHAR(255)"`
	Price        string    `xorm:"not null DECIMAL(10)"`
	RegisterTime time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' DATETIME"`
	UpdateTime   time.Time `xorm:"DATETIME"`
}
