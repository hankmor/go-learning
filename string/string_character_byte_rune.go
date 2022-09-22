package main

import (
	"fmt"
	"unicode/utf8"
)

// go中的字符串、字节、字符 character 和符文 rune 的区别
// demo code from: https://go.dev/blog/strings

func main() {
	characterAndRune()
	fmt.Println("================")

	// 由八个字节组成的字符串，字节 byte 的取值范围为 00 ~ FF
	const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

	fmt.Println(len(sample)) // 8

	demo1(sample)
	fmt.Println("================")
	demo2(sample)
	fmt.Println("================")
	demo3(sample)

	fmt.Println("\n================")
	printSpecialString()

	fmt.Println("\n================")
	printCodePoint()
	printCodePointWithUtf8()
}

func printCodePointWithUtf8() {
	const nihongo = "日本語"
	fmt.Println(len(nihongo)) // 9

	// 不使用 for...range 语句，而是用utf8包来解码rune
	for i, j := 0, 0; i < len(nihongo); i += j {
		// 将字符串解码为一个 rune，返回解码后的 rune 以及用到的字节宽度
		runeValue, index := utf8.DecodeRuneInString(nihongo[i:])
		fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
		j = index
	}
	// Output:
	// U+65E5 '日' starts at byte position 3
	// U+672C '本' starts at byte position 3
	// U+8A9E '語' starts at byte position 3

	const nihongo1 = "日本\xbc語"
	fmt.Println(len(nihongo1)) // 10

	// 不使用 for...range 语句，而是用utf8包来解码rune
	for i, j := 0, 0; i < len(nihongo1); i += j {
		// 将字符串解码为一个 rune，返回解码后的 rune 以及用到的字节宽度
		runeValue, index := utf8.DecodeRuneInString(nihongo1[i:])
		fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
		j = index
	}
	// Output:
	// U+65E5 '日' starts at byte position 3
	// U+672C '本' starts at byte position 3
	// U+FFFD '�' starts at byte position 1
	// U+8A9E '語' starts at byte position 3
}

// 输出 unicode 码点和byte序列的关系
func printCodePoint() {
	const nihongo = "日本語"
	fmt.Println(len(nihongo)) // 9

	// 使用 for range 时，go 可以将字节序列解码为 UTF-8 的 rune
	for index, runeValue := range nihongo {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	}
	// Output:
	// U+65E5 '日' starts at byte position 0
	// U+672C '本' starts at byte position 3
	// U+8A9E '語' starts at byte position 6

	// const nihongo1 = "日本😄語"
	const nihongo1 = "日本\xbd語"
	// invalid utf8 字符串
	fmt.Println(utf8.ValidString("\xbd")) // false
	fmt.Println(len(nihongo1))            // 10

	// 使用 for range 时，go 可以将字节序列解码为 UTF-8 的 rune
	for index, runeValue := range nihongo1 {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	}
	// Output:
	// U+65E5 '日' starts at byte position 0
	// U+672C '本' starts at byte position 3
	// U+FFFD '�' starts at byte position 6
	// U+8A9E '語' starts at byte position 7
}

func printSpecialString() {
	const s = '⌘' // 单个字节的 char，其实是一个 rune

	fmt.Printf("%x\n", s) // 十六进制格式：0x2318
	fmt.Printf("%T\n", s) // 类型是 int32, 也就是 rune

	const placeOfInterest = `⌘`

	fmt.Printf("plain string: ")
	fmt.Printf("%s\n", placeOfInterest) // ⌘
	fmt.Printf("%x\n", placeOfInterest) // 十六进制格式：e28c98
	fmt.Printf("%T", placeOfInterest)   // 类型是 string
	fmt.Printf("\n")

	fmt.Printf("quoted string: ")
	fmt.Printf("%+q", placeOfInterest) // "\u2318"
	fmt.Printf("\n")

	fmt.Printf("hex bytes: ")
	for i := 0; i < len(placeOfInterest); i++ {
		fmt.Printf("%x ", placeOfInterest[i]) // e2 8c 98，3个字节组成
	}
	fmt.Printf("\n")
}

func characterAndRune() {
	var s = "你好"
	fmt.Println("len: ", len(s)) // 6

	fmt.Println("loop with bytes: ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	fmt.Println("loop with rang: ")
	for _, c := range s {
		fmt.Printf("%x ", c) // 字节值，unicode 的 U+4f60 代表 "你"，U+597d 表示 "好"
		fmt.Printf("%c ", c) // 字符值，可以正确输出中文
	}
	fmt.Println()

	// 与上边的range等效
	fmt.Println("rune loop: ")
	rs := []rune(s)
	for _, c := range rs {
		fmt.Printf("%c ", c) // rune, 正确输出中文
	}
	fmt.Println()
}

func demo1(sample string) {
	fmt.Println("Println:")
	fmt.Println(sample)
	// 输出: ��=� ⌘

	fmt.Println("Byte loop:")
	for i := 0; i < len(sample); i++ {
		fmt.Printf("%x ", sample[i]) // 按照十六进制格式输出单个字节
	}
	// 输出: bd b2 3d bc 20 e2 8c 98

	fmt.Printf("\n")

	fmt.Println("Printf with %x:")
	fmt.Printf("%x\n", sample) // 整个字符串输出为十六进制格式
	// 输出：bd b2 3d bc 20 e2 8c 98

	fmt.Println("Printf with % x:")
	fmt.Printf("% x\n", sample) // 多个一个空格
	// 输出：bdb23dbc20e28c98

	fmt.Println("Printf with %q:")
	fmt.Printf("%q\n", sample) // 转义字符串中不可打印的字符序列
	// 输出："\xbd\xb2=\xbc ⌘"

	fmt.Println("Printf with %+q:")
	fmt.Printf("%+q\n", sample) // 输出不仅转义不可打印的序列，而且转义任何非 ASCII 字节，并解释 UTF-8 字符
	// 输出："\xbd\xb2=\xbc \u2318"
}

func demo2(s string) {
	// 使用字节切片来打印
	sample := []byte(s)
	fmt.Println("Println:")
	fmt.Println(sample)
	// 输出: [189 178 61 188 32 226 140 152]

	fmt.Println("Byte loop:")
	for i := 0; i < len(sample); i++ {
		fmt.Printf("%x ", sample[i]) // 按照十六进制格式输出单个字节
	}
	// 输出: bd b2 3d bc 20 e2 8c 98

	fmt.Printf("\n")

	fmt.Println("Printf with %x:")
	fmt.Printf("%x\n", sample) // 整个字符串输出为十六进制格式
	// 输出：bd b2 3d bc 20 e2 8c 98

	fmt.Println("Printf with % x:")
	fmt.Printf("% x\n", sample) // 多个一个空格
	// 输出：bdb23dbc20e28c98

	fmt.Println("Printf with %q:")
	fmt.Printf("%q\n", sample) // 转义字符串中不可打印的字符序列
	// 输出："\xbd\xb2=\xbc ⌘"

	fmt.Println("Printf with %+q:")
	fmt.Printf("%+q\n", sample) // 输出不仅转义不可打印的序列，而且转义任何非 ASCII 字节，并解释 UTF-8 字符
	// 输出："\xbd\xb2=\xbc \u2318"
}

func demo3(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%q", s[i])
	}
	// 输出：'½''²''=''¼'' ''â''\u008c''\u0098'
}
