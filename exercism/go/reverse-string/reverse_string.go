package reverse

/**import (
	"fmt"

)
*/
func Reverse(in string) string {
	s1 := []rune(in)
	s2 := []rune{}
	l := len(s1)
	for i:=0; i< l ; i++ {
		///s2= append(s2,s1[l-i-1:l-i]...)
		// or
	    s2= append(s2,s1[l-i-1:l-i][0])
		
	}
	return string(s2)
}



//Fails on multibyte support 
func BadReverse(in string) string {
	l := len(in)
	out := ""
	for i:=0; i< l ; i++ {
		c:= in[l-i-1:l-i]
		out += c
	}
	return out
}