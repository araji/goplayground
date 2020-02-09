package cipher

import (
	"strings"
	"unicode"
	"regexp"
)

var (
	re  = regexp.MustCompile(`[a-zA-Z]+`)
	nre = regexp.MustCompile(`[^a-z]+`)
)
type CaesarCipher int
type ShiftCipher int
type VigenereCipher string

func abs(n int) int {
	if n < 0 {
		return -n
	} else {
		return n
	}
}

func StringShift(m string, d int) string {
	out := make([]rune, 0)
	for _, c := range m {
		char := unicode.ToLower(c)
		if char >= 97 && char <= 122 {
			shift := (((char - 97) + rune(d)) + 26) % 26
			char = 97 + shift
			out = append(out, char)
		}
	}
	return string(out)
}

func (cphr CaesarCipher) Encode(m string) string {
	return StringShift(m, int(cphr))
}

func (cphr CaesarCipher) Decode(m string) string {
	return StringShift(m, 26-int(cphr))

}

func (scphr ShiftCipher) Encode(m string) string {
	return StringShift(m, int(scphr))

}

func (scphr ShiftCipher) Decode(m string) string {
	return StringShift(m, 26-int(scphr))

}

func (v VigenereCipher) Encode(m string) string {
	m1 := re.FindAllString(strings.ToLower(m),-1)
	message := strings.Join(m1,"")
	key := []rune(strings.Repeat(string(v), 1+(len(message)/len(v))))
	out := []rune{}
	for ndx, chr := range message {
	        dec := 97 + ( (chr  - 97)  + ( key[ndx] - 97 ) ) % 26
			out = append(out, dec)
	}
	return string(out)
}

func (v VigenereCipher) Decode(m string) string {
	msg :=[]rune(m)
	key := []rune(strings.Repeat(string(v), 1+(len(msg)/len(v))))
	out := []rune{}
	for ndx, chr := range msg {
	    newchr := 97 + ( ( chr  - 97) - ( key[ndx] - 97  ) + 26 ) % 26
	    out = append(out, newchr)
	}
	return string(out)
}

func NewCaesar() Cipher {
	return CaesarCipher(3)
}

func NewShift(distance int) Cipher {
	if (abs(distance) < 1) || (abs(distance) > 25) {
		return nil
	}
	return ShiftCipher(distance)
}

func NewVigenere(key string) Cipher {
	testkey := strings.Replace(key,"a","",-1)
	if   ! (len(testkey) > 1 )  || ( nre.MatchString(testkey))  {
		return nil
	}
	
	return VigenereCipher(key)

}
