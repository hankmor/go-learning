package shard

import (
	"fmt"
	"gorm-demo/model"
	"strconv"
)

// ==============================
// 分表测试，使用User组合
// ==============================

var n = 2 // 2 张表，user1和user2

func prepare() {
	for i := 0; i < 100; i++ {
		var u any
		switch i % n {
		case 0:
			u = &model.User1{User: model.User{Name: "name" + strconv.Itoa(i), Age: i}}
		case 1:
			u = &model.User2{User: model.User{Name: "name" + strconv.Itoa(i), Age: i}}
		}
		_, err := model.Create(u)
		if err != nil {
			panic(err)
		}
	}
}
func TestShardModelByCompose() {
	//prepare()

	// get
	var mdl any
	for i := 0; i < 100; i++ {
		mdl = getModelByCompose(i)
		// model使用底层的具体类型
		u, err := model.GetUserByName(mdl, "name"+strconv.Itoa(i))
		if err != nil {
			panic(err)
		}
		// 想要获取具体的 User1或者User2，只能使用断言，如果表实例很多会很臃肿
		if u1, ok := u.(*model.User1); ok {
			fmt.Println(u1.ID, u1.Name, u1.Age)
		} else if u2, ok := u.(*model.User2); ok {
			fmt.Println(u2.ID, u2.Name, u2.Age)
		} else {
			panic("invalid user type")
		}
	}

	// 测试 update，冗长的断言
	var err error
	var a any
	var ok bool
	var u1 *model.User1
	var u2 *model.User2
	id := 1
	mdl = getModelByCompose(id)
	u, err := model.GetUser(mdl, uint(id))
	if u1, ok = u.(*model.User1); ok {
		fmt.Println("before update:", u1.Age)
		a, err = u1.Update(uint(id), 20)
		fmt.Println("after update:", a.(*model.User1).Age)
	} else if u2, ok = u.(*model.User2); ok {
		fmt.Println("before update:", u2.Age)
		a, err = u2.Update(uint(id), 20)
		fmt.Println("after update:", a.(*model.User2).Age)
	} else {
		panic("invalid user type")
	}
	if err != nil {
		panic(err)
	}
}

func getModelByCompose(i int) any {
	switch i % n {
	case 0:
		return &model.User1{}
	case 1:
		return &model.User2{}
	}
	return nil
}
