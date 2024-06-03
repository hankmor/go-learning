package main

import (
	"fmt"
	"time"
)

func main() {
	// newTimer()
	// afterFunc()
	reset()
}

func newTimer() {
	timer := time.NewTimer(time.Second * 5)
	defer timer.Stop()

	done := make(chan bool)
	go func() {
		time.Sleep(time.Second * 6)
		done <- true
	}()

	for {
		select {
		case <-done:
			fmt.Println("Done!")
			return
		case t := <-timer.C:
			fmt.Println("current time:", t)
		}
	}
}

func afterFunc() {
	var timer *time.Timer
	timer = time.AfterFunc(time.Second*5, func() {
		// t := <-timer.C // NOTE: DO NOT To USE IT, not used and will be nil
		// fmt.Println("current time:", t)
		fmt.Println("after 5 seconds later...")
	})

	time.Sleep(6 * time.Second)
	timer.Stop()
	fmt.Println("Done!")
}

func reset() {
	timer := time.NewTimer(time.Second)

	done := make(chan bool)
	go func() {
		time.Sleep(6 * time.Second)
		done <- true
	}()

	go func() {
		time.Sleep(2 * time.Second)
		// if Stop reports that the timer expired before being stopped—the channel explicitly drained
		// if !timer.Stop() { // already stopped
			// <-timer.C
		// }
		// reset to 1 seconds
		timer.Reset(2 * time.Second)
		fmt.Println("reset timer")
	}()
	for {
		select {
		case <-done:
			fmt.Println("done!")
			return
		case t := <-timer.C:
			fmt.Println("current time:", t)
			// reset timer to 2 seconds
			// timer.Reset(time.Second) // the same result to ticker
		}
	}
}
