package main

import "fmt"

// 并发地枚举所有素数：仅能被自己和1整数的数
//
// 埃拉托斯特尼素数筛除算方法：从最小的 2 开始，筛除其倍数的数，下一个未被筛除的数（这里是3）就是素数，
// 再用这个素数 3 去筛除其倍数的数……不断重复，直到全部筛除。

// 从2开始累加，不断生成整数并写入 ch 通道中
func Generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i
	}
}

// 不断从生成整数的通道 in 中读取数据并基于基础 prime 筛选素数，不能被其整除的数则为素数，并将其写入 out 通道
func Filter(in <-chan int, out chan<- int, prime int) {
	for {
		n := <-in
		if n%prime != 0 { // 不能被整数则为素数，写回通道
			out <- n
		}
	}
}

func main() {
	ch := make(chan int)
	go Generate(ch) // 开启单独的 goroutine 生成从 2 开始并累加的整数
	for {
		prime := <-ch // 获取基础，从 2 开始
		fmt.Println("sieve: ", prime)
		out := make(chan int)     // 用于接收已经筛选的素数
		go Filter(ch, out, prime) // 筛选素数，prime 为基数，从 2 开始，将已经筛选的素数作为基数，如 2 3 5 7 ...，然后开启单独的goroutine不断筛掉该基数的倍数
		ch = out                  // 已经筛选出来的通道 out 赋值为 ch，将通道的素数作为基数，继续筛选
	}
}
