package main

import (
	"./morestrings"
	"fmt"
)

func main() {
	fmt.Println("Hello, go!")
	fmt.Println(morestrings.ReverseRunes("Hello, go!"))
}