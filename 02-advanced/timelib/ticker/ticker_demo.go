package main

import (
	"fmt"
	"time"
)

func main() {
	// tick()
	reset()
}

func tick() {
	// create new ticker which will tick each second
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	done := make(chan bool)
	// start to goroutine to wait 10s then write to done
	go func() {
		time.Sleep(10 * time.Second)
		done <- true
	}()
	// listen done chan and ticker, will print each ticking time before done
	for {
		select {
		case <-done:
			fmt.Println("Done!")
			return
		case t := <-ticker.C:
			// will print current time 10 times each second
			fmt.Println("Current time: ", t)
		}
	}
}

func reset() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	done := make(chan bool)
	go func() {
		time.Sleep(10 * time.Second)
		done <- true
	}()
	// after 2 seconds later, ticker will be reset and tick each 2 seconds
	go func() {
		time.Sleep(2 * time.Second)
		ticker.Reset(2 * time.Second)
	}()

	for {
		select {
		case <-done:
			fmt.Println("Done!")
			return
		case t := <-ticker.C:
			fmt.Println("time:", t)
		}
	}
}
