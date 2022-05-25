package db

import (
	"time"
)

type User struct {
	Id      int64     `json:"id"` // Id类型为uint64，则自动主键
	Name    string    `xorm:"varchar(25) notnull comment('姓名')" json:"name"`
	Age     int8      `xorm:"int comment('年龄')" json:"age"`
	Gender  int8      `xorm:"smallint(1) comment('性别')" json:"gender"`
	Created time.Time `xorm:"datetime created comment('创建时间')" json:"created"`
	Version uint      `xorm:"version" json:"-"` // 乐观锁
}

// TableName 通过TableName方法自定义表名称
func (e *User) TableName() string {
	return "sys_user"
}

func (e *User) Insert() (int64, error) {
	return E.InsertOne(e)
}

func (e *User) Update() (int64, error) {
	return E.ID(e.Id).Update(e)
}

func (e *User) Delete() (int64, error) {
	return E.ID(e.Id).Delete(e)
}

func (e *User) GetOne(id int64) (*User, error) {
	user := new(User)
	_, err := E.ID(id).Get(user)
	return user, err
}

func (e *User) GetAll() ([]*User, error) {
	users := make([]*User, 0)
	err := E.Where("true").Find(&users)
	return users, err
}
