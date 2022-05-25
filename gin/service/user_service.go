package service

import (
	"fmt"
	"gin-demo/db"
)

type UserService struct {
	dao db.User
}

func (us *UserService) Add(user *db.User) bool {
	rows, err := user.Insert()
	if err != nil {
		panic(err)
		return false
	}
	return rows == 1
}

func (us *UserService) Delete(id int64) bool {
	rows, err := us.dao.Delete()
	if err != nil {
		panic(err)
		return false
	}
	return rows == 1
}

func (us *UserService) Update(user *db.User) bool {
	rows, err := user.Update()
	if err != nil {
		panic(err)
		return false
	}
	return rows == 1
}

func (us *UserService) Get(id int64) *db.User {
	fmt.Printf("%T", us.dao)
	user, err := us.dao.GetOne(id)
	if err != nil {
		panic(err)
		return nil
	}
	return user
}

func (us *UserService) GetAll() []*db.User {
	users, err := us.dao.GetAll()
	if err != nil {
		panic(nil)
		return make([]*db.User, 0)
	}
	return users
}
