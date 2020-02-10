package reverse

import (
	"bytes"
)
func Reverse2(in string) string {
	s1 := []rune(in)
	s2 := []rune{}
	l := len(s1)
	for i:=0; i< l ; i++ {
	    s2= append(s2,s1[l-i-1:l-i][0])
	}
	return string(s2)
}


//https://exercism.io/tracks/go/exercises/reverse-string/solutions/34d0b3964fac4977aa4a743ff6adfbf8
// String returns the input string in reverse mode
func Reverse(input string) string {
	var output bytes.Buffer

	runes := []rune(input)
	for i := len(runes) - 1; i >= 0; i-- {
		output.WriteRune(runes[i])
	}

	return output.String()
}