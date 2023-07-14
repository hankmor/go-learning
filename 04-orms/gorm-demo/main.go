package main

import (
	"errors"
	"fmt"
	"gorm-demo/db"
	"gorm-demo/model"
	"gorm-demo/shard"
	"gorm.io/gorm"
	"reflect"
)

var DB *gorm.DB

func main() {
	DB = db.Conn()
	model.DB = DB
	// 迁移 schema
	_ = DB.AutoMigrate(model.User{}, model.Dept{}, model.User1{}, model.User2{})

	//testDB()
	//testCreate()
	//testSave()
	//testFirstNotFound()
	//testFirst()
	//testLastNotFound()
	//testFindOne()
	//testFindSlice()
	//testGenericMethod()

	//testFunctionalFirst()

	shard.TestShardModelByCompose()
	//shard.TestShardModelByInterface()
}

func testDB() {
	fmt.Println("===== testDB")
	fmt.Println(DB.Model(model.User{}) == DB) // false, 克隆了一个实例
	fmt.Println(DB.Where("id = ?", 0) == DB)  // false
	fmt.Println(DB.Create(struct{}{}) == DB)  // false

	db1 := DB.Model(struct{}{})                                    // 从DB克隆一个db
	fmt.Println(db1.Where("id = ?", 0) == db1)                     // true
	fmt.Println(db1.Where("id = ?", 0).Find(&model.User{}) == db1) // true
	fmt.Println(db1.Create(struct{}{}) == db1)                     // true
	fmt.Println(db1.Save(struct{}{}) == db1)                       // false，底层更新时复制了DB，底层首先执行update，如果update影响函数为0这create
	fmt.Println(db1.Updates(struct{}{}) == db1)                    // true
	fmt.Println(db1.Update("a", struct{}{}) == db1)                // true

	// 调用 db.getInstance() 复制一个实例，链式操作期间不会复制实例
}

func testCreate() {
	fmt.Println("结构体参数不寻址，直接panic: ")
	//var usr1 = model.User{Name: "张三", Age: 10}
	//r1 := DB.Create(usr1) // 直接panic: reflect: reflect.Value.Set using unaddressable value
	//fmt.Println("id:", usr1.ID)
	//fmt.Println("handleerr:", r1.Error)
	//fmt.Println("rowAffect:", r1.RowsAffected)

	fmt.Println("结构体参数寻址寻址: ")
	var usr2 = model.User{Name: "张三", Age: 10}
	r2 := DB.Create(&usr2)
	fmt.Println("id:", usr2.ID)                // 1
	fmt.Println("handleerr:", r2.Error)        // nil
	fmt.Println("rowAffect:", r2.RowsAffected) // 1
}

func testSave() {

}

func testFirstNotFound() {
	fmt.Println("===== testFirstNotFound")
	var usr *model.User
	fmt.Println("before:", usr)
	// 查询不存在的记录，结果对象不会为nil，而是零值，结果的db中err不为nil
	// 参数寻址
	fmt.Println("指针变量参数寻址：")
	r := DB.Model(model.User{}).Where("id = ?", 0).First(&usr)
	fmt.Println("after:", usr)                              // 零值 &{{0 0001-01-01 00:00:00 +0000 UTC...
	fmt.Println(errors.Is(r.Error, gorm.ErrRecordNotFound)) // true
	fmt.Println("handleerr:", r.Error)                      // record not found

	// 参数不寻址，错误用法
	fmt.Println("指针变量参数不寻址：")
	var usr1 *model.User
	r1 := DB.Model(model.User{}).Where("id = ?", 0).First(usr1)
	fmt.Println("after:", usr1)                              // nil
	fmt.Println(errors.Is(r1.Error, gorm.ErrRecordNotFound)) // false
	fmt.Println("handleerr:", r1.Error)                      // 错误信息：invalid value, should be pointer to struct or slice

	// 非指针
	fmt.Println("结构体变量参数寻址：")
	var usr2 model.User
	r2 := DB.Model(model.User{}).Where("id = ?", 0).First(&usr2)
	fmt.Println("after:", usr2)                              // 零值，非指针 {{0 0001-01-01 00:00:00 +0000 UTC...
	fmt.Println(errors.Is(r2.Error, gorm.ErrRecordNotFound)) // true
	fmt.Println("handleerr:", r2.Error)                      // record not found
	fmt.Println("结构体变量参数不寻址：")
	var usr3 model.User
	r3 := DB.Model(model.User{}).Where("id = ?", 0).First(usr3)
	fmt.Println(usr2 == usr3)                                // true
	fmt.Println("after:", usr3)                              // 零值，非指针 {{0 0001-01-01 00:00:00 +0000 UTC...
	fmt.Println(errors.Is(r3.Error, gorm.ErrRecordNotFound)) // true
	fmt.Println("handleerr:", r3.Error)                      // record not found
}

func testLastNotFound() {
	fmt.Println("===== testLastNotFound")
	var usr *model.User
	fmt.Println("before:", usr)

	// 查询不存在的记录，结果对象不会为nil，而是零值，结果的db中err不为nil

	// 参数寻址
	fmt.Println("指针变量参数寻址：")
	r := DB.Model(model.User{}).Where("id = ?", 0).Last(&usr)
	fmt.Println("after:", usr)                              // 零值 &{{0 0001-01-01 00:00:00 +0000 UTC...
	fmt.Println(errors.Is(r.Error, gorm.ErrRecordNotFound)) // true
	fmt.Println("handleerr:", r.Error)                      // record not found

	// 参数不寻址
	fmt.Println("指针变量参数不寻址，错误用法：")
	var usr1 *model.User
	// 传递空的指针参数，gorm.DB.Error 会存储错误信息，这种使用方式错误
	r1 := DB.Model(model.User{}).Where("id = ?", 0).Last(usr1)
	fmt.Println("after:", usr1)                              // nil，注意这里返回的是nil
	fmt.Println(errors.Is(r1.Error, gorm.ErrRecordNotFound)) // false
	fmt.Println("handleerr:", r1.Error)                      // 直接存在错误信息：invalid value, should be pointer to struct or slice

	// 非指针
	fmt.Println("结构体变量参数寻址：")
	var usr2 model.User
	r2 := DB.Model(model.User{}).Where("id = ?", 0).Last(&usr2)
	fmt.Println("after:", usr2)                              // 零值，非指针 {{0 0001-01-01 00:00:00 +0000 UTC...
	fmt.Println(errors.Is(r2.Error, gorm.ErrRecordNotFound)) // true
	fmt.Println("handleerr:", r2.Error)                      // record not found

	fmt.Println("结构体变量参数不寻址：")
	var usr3 model.User
	// 传递结构体，正常执行，gorm.DB.Error 是record not found
	r3 := DB.Model(model.User{}).Where("id = ?", 0).Last(usr3)
	fmt.Println(usr2 == usr3)                                // true
	fmt.Println("after:", usr3)                              // 零值，非指针 {{0 0001-01-01 00:00:00 +0000 UTC...
	fmt.Println(errors.Is(r3.Error, gorm.ErrRecordNotFound)) // true
	fmt.Println("handleerr:", r3.Error)                      // record not found
}

func testFirst() {
	id := 1 // 存在的记录
	fmt.Println("===== testFirst")
	var usr *model.User
	// 参数寻址
	fmt.Println("指针变量参数寻址：")
	r := DB.Model(model.User{}).Where("id = ?", id).First(&usr)
	fmt.Println("after:", usr) // ok, 指针
	fmt.Println("rowAffected:", r.RowsAffected)
	fmt.Println("handleerr:", r.Error)
	// 参数不寻址
	fmt.Println("指针变量参数不寻址：")
	var usr1 *model.User
	r1 := DB.Model(model.User{}).Where("id = ?", id).First(usr1)
	fmt.Println("after:", usr1)                  // nil
	fmt.Println("rowAffected:", r1.RowsAffected) // 0
	fmt.Println("handleerr:", r1.Error)          // 错误信息： invalid value, should be pointer to struct or slice
	fmt.Println(usr == usr1)                     // false

	// 非指针
	fmt.Println("结构体变量参数寻址：")
	var usr2 model.User
	r2 := DB.Model(model.User{}).Where("id = ?", id).First(&usr2)
	fmt.Println("after:", usr2) // ok
	fmt.Println("rowAffected:", r2.RowsAffected)
	fmt.Println("handleerr:", r2.Error)
	fmt.Println("结构体变量参数不寻址：直接 panic")
	//var usr3 model.User
	//r3 := DB.Model(model.User{}).Where("id = ?", id).First(usr3) // panic: reflect: reflect.Value.SetUint using unaddressable value
	//fmt.Println("after:", usr3)
	//fmt.Println("rowAffected:", r3.RowsAffected) //
}

func testFindOne() {
	id := 1
	fmt.Println("===== testFindOne")
	//var us model.User
	//r := DB.Model(model.User{}).Where("id = ?", id).Find(us) // panic: reflect: reflect.Value.Set using unaddressable value
	//fmt.Println("user:", us)
	//fmt.Println("rowAffected:", r.RowsAffected)
	//fmt.Println("handleerr:", r.Error)

	var us1 model.User
	r1 := DB.Model(model.User{}).Where("id = ?", id).Find(&us1) // panic: reflect: reflect.Value.Set using unaddressable value
	fmt.Println("user:", us1)
	fmt.Println("rowAffected:", r1.RowsAffected) // 1
	fmt.Println("handleerr:", r1.Error)          // nil

	var us2 *model.User
	r2 := DB.Model(model.User{}).Where("id = ?", id).Find(us2)
	fmt.Println("user:", us2)                    // nil
	fmt.Println("rowAffected:", r2.RowsAffected) // 0
	fmt.Println("handleerr:", r2.Error)          // invalid value, should be pointer to struct or slice

	var us3 []*model.User
	r3 := DB.Model(model.User{}).Find(&us3) // panic: reflect: reflect.Value.Set using unaddressable value
	fmt.Println("user:", us3)
	fmt.Println("rowAffected:", r3.RowsAffected)
	fmt.Println("handleerr:", r3.Error)
}

func testFindSlice() {
	fmt.Println("===== testFindSlice")
	//var us []model.User
	//DB.Model(model.User{}).Find(us) // panic: reflect: reflect.Value.Set using unaddressable value
	//fmt.Println(us)

	var us1 []model.User
	DB.Model(model.User{}).Find(&us1) // panic: reflect: reflect.Value.Set using unaddressable value
	fmt.Println(us1)                  // ok，元素为结构体

	//var us2 []*model.User
	//DB.Model(model.User{}).Find(us2) // panic: reflect: reflect.Value.Set using unaddressable value
	//fmt.Println(us2)

	var us2 []*model.User
	DB.Model(model.User{}).Find(&us2) // panic: reflect: reflect.Value.Set using unaddressable value
	fmt.Println(us2)                  // ok，指针
}

func testGenericMethod() {
	fmt.Println("===== testGenericMethod")
	//DB.Create(&model.Dept{Name: "IT"})
	var findAll = func(db *gorm.DB, entity any) any {
		//m := reflect.New(reflect.SliceOf(reflect.TypeOf(entity))).Elem().Interface()
		m := reflect.New(reflect.SliceOf(reflect.TypeOf(entity))).Interface()
		fmt.Println("type: ", reflect.TypeOf(m))
		db.Model(entity).Find(m)
		return m
	}

	us := findAll(DB, model.User{})
	fmt.Println(us)
	ds := findAll(DB, model.Dept{})
	fmt.Println(ds)
}

type QueryOption func(db *gorm.DB) (*gorm.DB, error)

func Where(query interface{}, args ...interface{}) QueryOption {
	return func(db *gorm.DB) (*gorm.DB, error) {
		ret := db.Where(query, args)
		return ret, ret.Error
	}
}

type DW struct {
	*gorm.DB
}

func (db *DW) First(out interface{}, opts ...QueryOption) error {
	// Get the GORM DB
	gdb := db.DB
	var err error = nil
	// Apply all the options
	for _, opt := range opts {
		gdb, err = opt(gdb)
		if err != nil {
			return err
		}
	}
	// Execute the `First` method and check for errors
	if err := gdb.First(out).Error; err != nil {
		return err
	}
	return nil
}

func testFunctionalFirst() {
	var DBW = &DW{DB}
	var user model.User
	err := DBW.First(&user,
		Where("name = ?", "张三"),
		Where("id = ?", 1),
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(user)
}
