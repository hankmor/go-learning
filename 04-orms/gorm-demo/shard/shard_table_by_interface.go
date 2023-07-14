package shard

import (
	"fmt"
	"gorm-demo/model"
	"strconv"
)

// ==============================
// 分表测试，使用interface和多态
// ==============================

func TestShardModelByInterface() {
	var mdl model.UserInterface
	for i := 0; i < 100; i++ {
		// model为接口的具体实现
		mdl = getModelByInterface(i)
		u, err := model.GetUserByNameIt(mdl, "name"+strconv.Itoa(i))
		if err != nil {
			panic(err)
		}
		fmt.Println(u.GetID(), u.GetName(), u.GetAge())
	}

	// 测试 update
	id := 1
	mdl = getModelByInterface(id)
	u, err := model.GetUserIt(mdl, uint(id))
	fmt.Println("before update:", u.GetAge())
	u1, err := mdl.UpdateIt(uint(id), 20)
	if err != nil {
		panic(err)
	}
	fmt.Println("after update:", u1.GetID(), u1.GetName(), u1.GetAge())
}

func getModelByInterface(i int) model.UserInterface {
	switch i % n {
	case 0:
		return &model.User1{}
	case 1:
		return &model.User2{}
	}
	return nil
}
