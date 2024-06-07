package main

import (
	"fmt"
	"unicode/utf16"
)

func main() {
	str := "➡️Click Me⬅️ or ➡️Click Me⬅️"
	utf16Runes := utf16.Encode([]rune(str))
	utf8Runes := utf16.Decode(utf16Runes)
	fmt.Println(string(utf8Runes))
}
