package basic

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func Test1(t *testing.T) {
	// 101 -> 10
	println(5 >> 1)
	i := rand.Int63n(time.Now().UnixMilli())
	println(i)
	// i := 257
	bs := IntToBytes(int64(i))
	fmt.Printf("%v\n", bs)
	i = BytesToInt(bs)
	fmt.Printf("%v\n", i)
}

func TestTruncateHigh(t *testing.T) {
	src := rand.Int63n(time.Now().UnixMilli())
	println(src)
	a := BytesToInt(IntToBytes(src))
	println(a)
	i := TruncateHigh(src)
	println(i)
}