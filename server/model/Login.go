package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Login struct {
	gorm.Model
	Name string `gorm:"unique; not null"`
	Pass string `gorm:"not null"`
}
