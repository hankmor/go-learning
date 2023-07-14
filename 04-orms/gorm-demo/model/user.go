package model

import (
	"fmt"
	"gorm.io/gorm"
	"reflect"
)

var DB *gorm.DB

type UserInterface interface {
	GetID() uint
	GetName() string
	GetAge() int
	SayHello() error
	UpdateIt(id uint, age int) (UserInterface, error)
}

type User struct {
	gorm.Model
	Name string `gorm:"type:varchar(32);default:'';comment:'名称'"`
	Age  int    `gorm:"type:int;default:0;comment:'年龄'"`
}

func (u *User) GetID() uint {
	return u.ID
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) GetAge() int {
	return u.Age
}

func (u *User) SayHello() error {
	fmt.Println("hello,", u.Name)
	return nil
}

func (u *User) Update(id uint, age int) (any, error) {
	fmt.Println(reflect.TypeOf(u)) // 即使是User1再调用该方法，这里的u的具体类型仍然是 *model.User
	r := DB.Model(u).Where("id = ?", id).Update("age", age)
	// 无法返回具体的u的类型User1、User2
	return u, r.Error
}

func (u *User) UpdateIt(id uint, age int) (UserInterface, error) {
	r := DB.Model(u).Where("id = ?", id).Update("age", age)
	return u, r.Error
}

func Create(user any) (any, error) {
	r := DB.Model(user).Create(user)
	return user, r.Error
}

func CreateIt(user UserInterface) (UserInterface, error) {
	r := DB.Model(user).Create(user)
	return user, r.Error
}

func GetUser(user any, id uint) (any, error) {
	if reflect.TypeOf(user).Kind() != reflect.Ptr {
		return nil, fmt.Errorf("query model need be an pointer")
	}
	r := DB.Model(user).Where("id = ?", id).Find(user)
	return user, r.Error
}

func GetUserIt(user UserInterface, id uint) (UserInterface, error) {
	r := DB.Model(user).Where("id = ?", id).Find(user)
	return user, r.Error
}

func GetUserByName(user any, name string) (any, error) {
	if reflect.TypeOf(user).Kind() != reflect.Ptr {
		return nil, fmt.Errorf("query model need be an pointer")
	}
	r := DB.Model(user).Where("name = ?", name).Find(user)
	return user, r.Error
}

func GetUserByNameIt(user UserInterface, name string) (UserInterface, error) {
	r := DB.Model(user).Where("name = ?", name).Find(user)
	return user, r.Error
}

// 分表测试

type User1 struct {
	User
}

type User2 struct {
	User
}
