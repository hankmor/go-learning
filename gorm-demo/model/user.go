package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string `gorm:"type:varchar(32);default:'';comment:'名称'"`
	Age  int    `gorm:"type:int;default:0;comment:'年龄'"`
}
