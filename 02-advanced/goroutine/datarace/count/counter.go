package main

import (
	"fmt"
	"sync"
)

func main() {
	// race()
	fixRace()
}

func race() {
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Print(i)
			wg.Done()
		}()
	}
	wg.Wait()
}

func fixRace() {
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(j int) {
			fmt.Print(j) // Good. Read local copy of the loop counter.
			wg.Done()
		}(i) // copy var i to j
	}
	wg.Wait()
}
