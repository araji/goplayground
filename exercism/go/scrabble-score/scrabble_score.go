// scrabble word scoring implementation
package scrabble

import (
	"unicode"
)

// catalog : mapping letter to their values
var catalog = map[rune] int {
	'A':1,'B':3,'C':3,'D':2,'E':1,'F':4,'G':2,'H':4,'I':1,'J':8,'K':5,'L':1,
	'M':3,'N':1,'O':1,'P':3,'Q':10,'R':1,'S':1,'T':1,'U':1,'V':4,'W':4,'X':8,
	'Y':4,'Z':10,
}
// Score 
// TODO: runs @3500 ns/op bring it down to < 1000
// DONE: switching to rune and unicode brought it down to 820!!!
func Score(s string) int {
	score := 0 
	for _, c := range(s) {
		//no need to test presence , score is 0 if string not in map
		score += catalog[unicode.ToUpper(c)]
	}

	return score
}