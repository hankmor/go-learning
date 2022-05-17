package morestrings

// ReverseRunes 翻转字符
func ReverseRunes(s string) string {
	rs := []rune(s)
	for i, j := 0, len(rs)-1; i < len(rs)/2; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	return string(rs)
}
