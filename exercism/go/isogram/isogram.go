package isogram

import (
	"unicode"
)

func IsIsogram(s string) bool {
	var amap = make(map[rune]int)
	for _, char := range s {
		if !unicode.IsLetter(char) {
			continue
		}
		char = unicode.ToLower(char)
		_, ok := amap[char]
		if ok {
			return false
		} else {
			amap[char] = 1
		}
	}
	return true
}
