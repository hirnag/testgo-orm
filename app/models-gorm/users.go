package models_gorm

type User struct {
    ID   int    `gorm:"primary_key"`
    Name string `gorm:"type:varchar(255)"`
}

