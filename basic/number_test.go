package main

import (
	"encoding/hex"
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

	fmt.Printf("%x\n", []byte("100"))
	toString := hex.EncodeToString([]byte("17"))
	println(toString)
}

func TestTruncateHigh(t *testing.T) {
	src := rand.Int63n(time.Now().UnixMilli())
	println(src)
	a := BytesToInt(IntToBytes(src))
	println(a)
	i := TruncateHigh(src)
	println(i)
}

func TestRemainLow(t *testing.T) {
	// src := rand.Int63n(time.Now().UnixMilli())
	src := time.Now().UnixMilli()
	println(src)
	// i := TruncateHigh(src)
	// println(i)
	// i = RemainLow(src, 18)
	// println(i)
	time.Sleep(time.Second)
	src = time.Now().UnixMilli()
	println(src)
	time.Sleep(time.Second)
	src = time.Now().UnixMilli()
	println(src)

	time.Sleep(time.Minute)
	src = time.Now().UnixMilli()
	println(src)

	time.Sleep(time.Minute)
	src = time.Now().UnixMilli()
	println(src)
}
