package howto

import (
	"fmt"
	"github.com/koobyte/go-learning/howto/morestrings"
)

func Run() {
	fmt.Println("Hello, go!")
	fmt.Println(morestrings.ReverseRunes("Hello, go!"))
}
