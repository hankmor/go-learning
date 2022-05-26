package basic

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"testing"
)

func TestHex(t *testing.T) {
	src := []byte("Hello")
	encodedStr := hex.EncodeToString(src)
	fmt.Printf("%s\n", encodedStr)
	decodeString, _ := hex.DecodeString(encodedStr)
	fmt.Printf("%s\n", decodeString)

	println("=====")

	src = []byte("1000")
	fmt.Printf("%s\n", strconv.FormatInt(1000, 16))
	fmt.Printf("%x\n", "1000")
	encodedStr = hex.EncodeToString(src)
	fmt.Printf("%s\n", encodedStr)
	decodeString, _ = hex.DecodeString(encodedStr)
	fmt.Printf("%s\n", decodeString)
}
