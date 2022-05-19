package main

import "fmt"

func Reverse(s string) string {
	bs := []byte(s)
	for i, j := 0, len(bs)-1; i < len(bs)/2; i, j = i+1, j-1 {
		bs[i], bs[j] = bs[j], bs[i]
	}
	return string(bs)
}

func main() {
	input := "The quick brown fox jumped over the lazy dog"
	rev := Reverse(input)
	doubleRev := Reverse(rev)
	fmt.Printf("original: %q\n", input)
	fmt.Printf("reversed: %q\n", rev)
	fmt.Printf("reversed again: %q\n", doubleRev)
}
