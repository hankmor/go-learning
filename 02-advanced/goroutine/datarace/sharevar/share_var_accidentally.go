package main

import "os"


func main() {
	bs := []byte("some data")
	ParallelWrite(bs)
}

// 意外共享变量导致数据竞态
// ParallelWrite writes data to file1 and file2, returns the errors.
func ParallelWrite(data []byte) chan error {
	res := make(chan error, 2)
	f1, err := os.Create("file1")
	if err != nil {
		res <- err
	} else {
		go func() {
			// This err is shared with the main goroutine,
			// so the write races with the write below.
			// _, err = f1.Write(data)
			// create a new error to avoid share variable
			_, err := f1.Write(data)
			res <- err
			f1.Close()
		}()
	}
	f2, err := os.Create("file2") // The second conflicting write to err.
	if err != nil {
		res <- err
	} else {
		go func() {
			// _, err = f2.Write(data)
			_, err := f2.Write(data)
			res <- err
			f2.Close()
		}()
	}
	return res
}
