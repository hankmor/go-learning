package main

import (
	"github.com/koobyte/go-learning/basic"
	"github.com/koobyte/go-learning/fuzz"
	"github.com/koobyte/go-learning/generic"
	"github.com/koobyte/go-learning/gowiki"
	howto "github.com/koobyte/go-learning/howto/hello"
	"github.com/koobyte/go-learning/oop"
)

func main() {
	howto.Run()
	basic.Run()
	fuzz.Run()
	oop.Run()
	generic.Run()
	json.Run()
	gowiki.Run()
	gowiki.Run()

}
