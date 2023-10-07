package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"testing"
	"time"
)

func TestReq(t *testing.T) {
	// nginx version: nginx/1.25.2
	// x nginx version: nginx/1.18.0 (Ubuntu)
	url := "http://127.0.0.1:9000" // nginx负载地址
	n := 4
	total := 0
	for {
		total++
		if total > 10 {
			break
		}
		var wg sync.WaitGroup
		wg.Add(n)
		for i := 0; i < n; i++ {
			go func(idx int) {
				r, _ := http.Get(url)
				bs, _ := io.ReadAll(r.Body)
				fmt.Println(idx, " => ", string(bs))
				wg.Done()
			}(i)
		}
		wg.Wait()
		time.Sleep(time.Second * 1)
		fmt.Println("==============")
	}
}

func TestTimeout(t *testing.T) {
	url := "http://127.0.0.1:9000"
	timeoutUrl := "http://127.0.0.1:9000/timeout"
	n := 11
	var sg sync.WaitGroup
	sg.Add(n)
	go func() {
		r, _ := http.Get(timeoutUrl)
		bs, _ := io.ReadAll(r.Body)
		fmt.Println(string(bs))
		sg.Done()
	}()
	for i := 0; i < n-1; i++ {
		go func(x int) {
			r, _ := http.Get(url)
			bs, _ := io.ReadAll(r.Body)
			fmt.Println(string(bs))
			sg.Done()
		}(i)
	}
	sg.Wait()
}
