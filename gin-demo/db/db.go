package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

const (
	username string = "root"
	password string = "12345678"
	database string = "test"
	host     string = "127.0.0.1"
	port     int16  = 3306
	showSql  bool   = true
)

var E *xorm.Engine

func init() {
	// sql driver 连接
	// db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", username, password, host, port, database))
	// if err != nil {
	// 	panic(err)
	// }
	// db.SetConnMaxLifetime(2 * time.Minute)
	// db.SetMaxOpenConns(5)
	// db.SetMaxIdleConns(5)
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", username, password, host, port, database)
	log.Println("Connecting to mysql: ", url)

	var err error
	E, err = xorm.NewEngine("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&charset=utf8mb4", username, password, host, port, database))
	if err != nil {
		panic(err)
	}

	E.SetMaxOpenConns(5)
	E.SetMaxIdleConns(5)
	E.SetConnMaxLifetime(2 * time.Minute)
	// 是否显示sql
	E.ShowSQL(showSql)
	// 实体名称映射规则
	E.SetMapper(names.GonicMapper{})

	// 初始化表
	user := new(User)
	E.CreateTables(user)

	// E.Sync(user)
	log.Printf("mysql Connected, url: %s:%d/%s\n", host, port, database)
}
