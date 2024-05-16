package main

// chan 的发送和关闭不同步也会导致数据竞态

func main() {
	c := make(chan struct{}) // or buffered channel

	// The race detector cannot derive the happens before relation
	// for the following send and close operations. These two operations
	// are unsynchronized and happen concurrently.
	// 通道 c 的发送和关闭在两个 goroutine 中，两个操作不同步会导致 data race
	// go func() {
	// 	c <- struct{}{}
	// }()
	// close(c)

	// 根据 Go 内存模型，通道上的发送发生在该通道的相应接收完成之前。要同步发送和关闭操作，请使用接收操作来保证发送在关闭之前完成：
	go func() {
		c <- struct{}{}
	}()
	<-c // 先接收在关闭
	close(c)
}
