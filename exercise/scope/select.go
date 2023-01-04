package main

import "fmt"

func main() {
	c1 := make(chan int)
	// c1 := make(chan int, 1) // 如果缓冲数量为1，效果不同
	c2 := make(chan int, 1) // c2 具有 1 个缓冲数量
	c2 <- 11                // 向 c2 写入数据

	select {
	case c1 <- 1: // 向c1写入数据，由于c1没有缓冲，故无法选择此分支
		fmt.Println("SendStmt case has been chosen")
	case i := <-c2: // 从c2读取数据，可以读取，程序会选择此分支
		_ = i
		fmt.Println("RecvStmt case has been chosen")
	default:
		fmt.Println("Default case has been chosen")
	}
	/*output:
	RecvStmt case has been chosen

	如果将 c1 改为缓冲数量为1，则会选择第一个分支，输出：SendStmt case has been chosen
	*/
}
