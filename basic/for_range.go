package main

import (
	"fmt"
	"time"
)

func main() {
	// ====== 从数组遍历获取一个指针元素切片集合
	// for range是语法糖，内部调用还是for循环，初始化会拷贝带遍历的列表（如array，slice，map）
	// 如果直接对v取地址，最终只会拿到一个地址，而对应的值就是最后遍历的那个元素所附给v的值

	// 错误
	arr := [2]int{1, 2}
	var res []*int
	for _, v := range arr {
		res = append(res, &v)
	}
	fmt.Println(*res[0], *res[1]) // expect 1 2, but is 2 2
	sl := []int{1, 2}
	res = []*int{}
	for _, v := range sl {
		res = append(res, &v)
	}
	fmt.Println(*res[0], *res[1]) // expect 1 2, but is 2 2
	mp := map[int]int{1: 1, 2: 2}
	res = []*int{}
	for _, v := range mp {
		res = append(res, &v)
	}
	fmt.Println(*res[0], *res[1]) // expect 1 2, but is 1 1 or 2 2
	// 正确
	res = []*int{}
	for _, v := range arr {
		tmp := v // 使用局部变量拷贝 v
		res = append(res, &tmp)
	}
	fmt.Println(*res[0], *res[1]) // 1 2

	// ===== 遍历时修改，循环会停止吗？
	// 遍历时，先将 v 拷贝一份，所以修改不会反映到原始 v 中
	v := []int{1, 2, 3}
	for i := range v {
		v = append(v, i)
	}
	fmt.Println(v)

	// ===== 遍历大数组，会先拷贝数组，极大的浪费内存空间
	// 但是遍历大的 slice、map 没有问题，因为它们是引用类型
	// 假设值都为1，这里只赋值3个
	arr1 := [102400]int{1, 1, 1}
	for i, n := range arr1 {
		// just ignore i and n for simplify the example
		_ = i
		_ = n
	}
	// 解决办法是寻址原数组，只拷贝地址
	for i, n := range &arr1 {
		_ = i
		_ = n
	}

	// ===== 循环中开启 goroutine，由于拷贝m，goroutine 中每次得到的都是最后一个元素
	var m = []int{1, 2, 3}
	for i := range m {
		go func() {
			fmt.Print(i) // 222
		}()
	}
	fmt.Println()
	// 解决办法是传递变量
	for i := range m {
		go func(i int) {
			fmt.Print(i)
		}(i)
	}
	fmt.Println()
	time.Sleep(time.Millisecond)
}
