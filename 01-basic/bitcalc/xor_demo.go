package main

import "fmt"

// 抑或运算定律：
// 1、交换律：x ^ y = y ^ x
// 2、集合律：x ^ (y ^ z) = (x ^ y) ^ z
// 3、任何数与自身进行异或运算，结果都为 0：x ^ x = 0
// 4、任何数与 0 进行异或运算, 结果为其自身：x ^ 0 = x
func main() {
	fmt.Println(swap(3, 4))

	text := "hello,this is a plain text，来段中文"
	key := "test-key"
	cipher := simpleEncrypt(text, key)
	fmt.Println("chiper: ", cipher)
	fmt.Println("raw: ", simpleDecrypt(cipher, key))

	fmt.Println(leecode([]int{1, 2, 3, 4, 4, 3, 2}))
}

// 交换两个变量的值
func swap(x, y int) (int, int) {
	// 异或运算性质: 三个值中的任意两个值进行异或运算，都可以得出第三个值
	fmt.Println(37 ^ 53) // 16
	fmt.Println(37 ^ 16) // 53
	fmt.Println(53 ^ 16) // 37

	// 设最初x、y为x0、y0
	x ^= y // x = x ^ y，x相当于临时变量c
	fmt.Println(x)
	y = x ^ y // 结果为x0，相当于 y = c ^ y = x0
	fmt.Println(y)
	x ^= y // 结果为y0，因为 c ^ x0 = y0
	return x, y
}

// 简单加密
func simpleEncrypt(text string, key string) string {
	bs := []byte(text)
	keybs := []byte(key)
	var n = 0
	for i, b := range bs {
		if n+1 == len(keybs) { // key长度不够，则从头开始
			n = 0
		}
		bs[i] = b ^ keybs[n]
		n++
	}
	return string(bs)
}

// 简单解密
func simpleDecrypt(cipher string, key string) string {
	bs := []byte(cipher)
	keybs := []byte(key)
	var n = 0
	for i, b := range bs {
		if n+1 == len(keybs) {
			n = 0
		}
		bs[i] = b ^ keybs[n]
		n++
	}
	return string(bs)
}

// 给你一个 非空 整数数组 nums ，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
// 要求: 必须设计并实现线性时间复杂度的算法来解决此问题，且该算法只使用常量额外空间。
// # 示例 1
// 输入：nums = [2,2,1]
// 输出：1
//
// # 示例 2
// 输入：nums = [4,1,2,1,2]
// 输出：4
//
// # 示例 3
// 输入：nums = [1]
// 输出：1
func leecode(nums []int) int {
	// 利用异或运算的性质:
	// 1. 任何数和 0 做异或运算，结果仍然是原来的数
	// 2. 任何数和其自身做异或运算，结果是 0
	// 3. 异或运算满足交换律和结合律，a^b^a = b^a^a = b^(a^a) = b^0 = b
	var x = 0
	for _, n := range nums {
		x ^= n
	}
	return x
}
