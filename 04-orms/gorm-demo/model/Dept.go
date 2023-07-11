package model

import "gorm.io/gorm"

type Dept struct {
	gorm.Model
	Name string `gorm:"type:varchar(32);default:0;comment:'名称'"`
}
