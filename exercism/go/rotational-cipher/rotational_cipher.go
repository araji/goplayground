package rotationalcipher

import (
	"strings"
)


// RotationalCipher same as simpleCipher except for case sensitivity and keeping non alpha.
// as described in strings.Map example :-).
func RotationalCipher(message string, shiftKey int ) string {
	
	rot13:= func(r rune) rune {	
		switch {
		case r >= 'A' && r <= 'Z':
			return 'A' + (r-'A'+ rune(shiftKey))%26
		case r >= 'a' && r <= 'z':
			return 'a' + (r-'a'+ rune(shiftKey) )%26
		}
		return r
	}
	
	return strings.Map(rot13,message)
}