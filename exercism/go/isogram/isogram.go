package isogram

import (
	"unicode"
)

// IsIsogram word or phrase without a repeating letter 
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

/**
Best solution seen :
 https://exercism.io/tracks/go/exercises/isogram/solutions/2e0f7084226541b5af92d9895d1a9fa3

 for i, c := range s {
		if unicode.IsLetter(c) && strings.ContainsRune(s[i+1:], c) {
			return false
		}
	}
*/