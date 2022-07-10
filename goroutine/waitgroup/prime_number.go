package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

// PrintPrimeNumber 打印 5000 以内的素数，耗时的操作，观察多个 goroutine 切换的情况
func PrintPrimeNumber() {
	runtime.GOMAXPROCS(1)

	wg.Add(2)

	// 打印素数，耗时操作
	go showPrimeNumber("A")
	go showPrimeNumber("B")

	println("Waiting to finish")
	wg.Wait()
	println("Terminating program")
}

func showPrimeNumber(prefix string) {
	defer wg.Done()

out:
	for i := 2; i < 5000; i++ {
		// 从2到这个数-1，如果有能整除的，说明不是素数，继续到外层循环
		for n := 2; n < i; n++ {
			if i%n == 0 {
				continue out
			}
		}
		fmt.Printf("%s: %d\n", prefix, i)
	}
	fmt.Printf("%s completed\n", prefix)
}

func main() {
	PrintPrimeNumber()
}
