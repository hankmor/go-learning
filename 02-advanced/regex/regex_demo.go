package main

import (
	"fmt"
	"regexp"
	"strings"
)

/*
Regex			Meaning
.				匹配任何单一字符
?				匹配前面的元素一次要么根本不匹配
+				匹配前面的元素一次或多次
*				匹配前一个元素的零次或多次
^				匹配字符串中的起始位置
$				匹配字符串中的结束位置
|				交替运算符
[abc]			匹配 a 或 b，或 c
[a-c]			a-c 的范围; 匹配 a 或 b，或 c
[^abc]			否定, 匹配除 a，或 b，或 c之外的一切
\s				匹配空格字符
\S				匹配非空格字符
\w				匹配一个单词字符；等同于 [a-zA-Z_0-9]
*/

func compile(pattern string) *regexp.Regexp {
	return regexp.MustCompile(pattern)
}

var r = compile("\\s#\\w+")
var s = "would you eat this? #shorts #sho #haha"

var r1 = compile("\\s#(\\w+)")
var r2 = compile("\\sthis\\?\\s") // this?

func MatchString() {
	// 匹配子字符串
	b := r.MatchString(s) // 是否能够匹配字符串
	fmt.Println(b)
}

func FindString() {
	a := r.FindString(s) // 查找字符串
	fmt.Println(a)       // #shorts
}

func FindAllString() {
	a := r.FindAllString(s, 3) // 查找所有匹配的字符串，第二个参数表示匹配的最大项，-1 表示匹配所有
	fmt.Println(a)             // [ #shorts  #sho  #haha]
}

func FindStringIndex() {
	a := r.FindStringIndex(s) // 返回匹配到的字符串的起止位置索引
	fmt.Println(a)            // 19 27
	fmt.Println(s[a[0]:a[1]]) // #shorts
}

func FindAllStringIndex() {
	a := r.FindAllStringIndex(s, -1) // 返回所有匹配的字符串的起止位置的索引
	fmt.Println(a)                   // [[19 27] [27 32] [32 38]]
	for _, p := range a {
		fmt.Print(s[p[0]:p[1]], " ") // #shorts  #sho  #haha
	}
	fmt.Println()
}

func FindStringSubmatch() {
	a := r1.FindStringSubmatch(s)        // 继续匹配下一个分组
	fmt.Println(a)                       // [ #shorts shorts]
	b := r1.FindAllStringSubmatch(s, -1) // 继续所有字符串，并继续匹配所有分组
	fmt.Println(b)                       // [[ #shorts shorts] [ #sho sho] [ #haha haha]]
}

func Split() {
	a := r2.Split(s, -1)
	fmt.Println(len(a)) // 2
	fmt.Println(a)      // [would you eat #shorts #sho #haha]
}

func ReplaceAllString() {
	a := r.ReplaceAllString(s, "**") // 替换匹配到的字符串
	fmt.Println(a)                   // would you eat this?******
}

func ReplaceAllStringFunc() {
	a := r.ReplaceAllStringFunc(s, strings.ToUpper) // 按照指定函数匹配替换到的字符串
	fmt.Println(a)                                  // would you eat this? #SHORTS #SHO #HAHA
}

func main() {
	MatchString()
	FindString()
	FindAllString()
	FindStringIndex()
	FindAllStringIndex()
	FindStringSubmatch()
	Split()
	ReplaceAllString()
	ReplaceAllStringFunc()
}
