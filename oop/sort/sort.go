package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
)

// sort demo
// sort 包

func main() {
	println("==== sort string ====")
	strs := []string{"haha", "xixi", "hehe"}
	sortString(strs)
	fmt.Printf("sorted: %v\n", strs) // [haha hehe xixi]

	println("==== sort int ====")
	ints := []int{3, 2, 1, 9, 8, 8, 7, 7}
	sortInt(ints)
	fmt.Printf("sorted: %v\n", ints) // [1 2 3 7 7 8 8 9]

	println("==== sort float ====")
	fs := []float64{1.11111, 1.11112, 1.111105, 1.2}
	sortFloat(fs)
	fmt.Printf("sorted: %v\n", fs) // [1.111105 1.11111 1.11112 1.2]

	sortUser()

	// 自定义排序，有时候排序的对象并不是slice
	sortCustom()
}

func sortUser() {
	users := []*user{
		{name: "huzhou", age: 18, score: 99.5},
		{name: "belonk", age: 20, score: 88.5},
		{name: "zhangsan", age: 17, score: 88.55},
	}
	println("==== sort user by name asc ====")
	sort.Sort(sortByName(users)) // 按名称排序
	printUsers(users)
	/*
		NAME      AGE  SCORE
		----      ---  -----
		belonk    20   88.5
		huzhou    18   99.5
		zhangsan  17   88.55
	*/
	println("==== sort user by name desc ====")
	// 按照名称排序，不需要在单独定义反序类型了，sort已经提供了Reverse类型
	sort.Sort(sort.Reverse(sortByName(users)))
	printUsers(users)
	/*
		NAME      AGE  SCORE
		----      ---  -----
		zhangsan  17   88.55
		huzhou    18   99.5
		belonk    20   88.5
	*/

	println("==== sort user by score asc ====")
	sort.Sort(sortByScore(users))
	printUsers(users)
	/*
		NAME      AGE  SCORE
		----      ---  -----
		belonk    20   88.5
		zhangsan  17   88.55
		huzhou    18   99.5
	*/
	println("==== sort user by score desc ====")
	sort.Sort(sort.Reverse(sortByScore(users)))
	printUsers(users)
	/*
		NAME      AGE  SCORE
		----      ---  -----
		huzhou    18   99.5
		zhangsan  17   88.55
		belonk    20   88.5
	*/
}

func printUsers(users []*user) {
	// 使用特定格式输出
	const cfmt = "%v\t%v\t%v\n"
	// 使用 tabwriter 包按照特定表格格式输出
	w := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(w, cfmt, "NAME", "AGE", "SCORE")
	fmt.Fprintf(w, cfmt, "----", "---", "-----")
	for _, u := range users {
		fmt.Fprintf(w, cfmt, u.name, u.age, u.score)
	}
	w.Flush() // 计算宽度并输出
}

type user struct {
	name  string
	age   int
	score float32
}

func (u user) String() string {
	return fmt.Sprintf("name = %s, age = %d, score = %.2f", u.name, u.age, u.score)
}

// 定义一个按照name排序的slice对象，实现了 sort.Interface 接口
type sortByName []*user

func (s sortByName) Len() int { // 排序对象的长度
	return len(s)
}

func (s sortByName) Less(i, j int) bool { // 排序的比较方式,必须可以比较
	return s[i].name < s[j].name
}

func (s sortByName) Swap(i, j int) { // 交换元素
	s[i], s[j] = s[j], s[i]
}

// 定义按照分数排序的对象
type sortByScore []*user

func (s sortByScore) Len() int { // 排序对象的长度
	return len(s)
}

func (s sortByScore) Less(i, j int) bool { // 排序的比较方式,必须可以比较
	return s[i].score < s[j].score
}

func (s sortByScore) Swap(i, j int) { // 交换元素
	s[i], s[j] = s[j], s[i]
}

func sortString(strs []string) {
	sort.Strings(strs)
}

func sortInt(ints []int) {
	sort.Ints(ints)
}

func sortFloat(fs []float64) {
	sort.Float64s(fs)
}

// 自定义user集，可以自定义排序
type customUsers struct {
	users []*user
	less  func(x, y *user) bool // 这是一个方法，用于判断两个user的排序方式
}

func (c customUsers) Len() int {
	return len(c.users)
}

// Less 调用自定义的方法进行排序
func (c customUsers) Less(i, j int) bool {
	return c.less(c.users[i], c.users[j])
}

func (c customUsers) Swap(i, j int) {
	c.users[i], c.users[j] = c.users[j], c.users[i]
}

func sortCustom() {
	println("==== custom sort ====")
	users := []*user{
		{name: "huzhou", age: 18, score: 99.5},
		{name: "huzhou", age: 16, score: 100},
		{name: "zhangsan", age: 17, score: 99.5},
		{name: "abbc", age: 17, score: 99.5},
	}

	// 自定义对象，排序时按照 score desc，age asc，name asc的顺序
	c := &customUsers{users, func(x, y *user) bool {
		if x.score != y.score { // score desc
			return x.score > y.score
		}
		if x.age != y.age { // age asc
			return x.age < y.age
		}
		if x.name != y.name { // name desc
			return x.name < y.name
		}
		return false
	}}
	sort.Sort(c)
	printUsers(c.users)
	/*
		NAME      AGE  SCORE
		----      ---  -----
		huzhou    16   100
		abbc      17   99.5
		zhangsan  17   99.5
		huzhou    18   99.5
	*/
}
