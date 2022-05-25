package howto

import (
	"./morestrings"
	"fmt"
)

func Run() {
	fmt.Println("Hello, go!")
	fmt.Println(morestrings.ReverseRunes("Hello, go!"))
}
