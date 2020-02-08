package cipher

import (
	"fmt"
	"strings"
	"unicode"
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
	//TODO: hackish , find a better way
	key := strings.Repeat(string(v), 1+(len(m)/len(v)))
	fmt.Println(key)
	out := make([]rune, 0)
	for ndx, chr := range m {
		fmt.Printf("%T\n", chr)

		chr = unicode.ToLower(chr)
		fmt.Printf("%T\n", key[ndx])
		fmt.Println(key[ndx])

		fmt.Println("message char", string(chr))
		fmt.Println("cipher char", string(key[ndx]))
		fmt.Println("cipher shift", key[ndx]-97)
		fmt.Println("cipher shift", chr+rune((key[ndx]-97)))
		newchr := chr
		//newchr :=  rune( (chr + ( key[ndx] -97))  % 26 )
		fmt.Println("new char", string(rune(newchr)))
		out = append(out, newchr)
	}
	return string(out)
}

func (v VigenereCipher) Decode(m string) string {
	return m
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
	return VigenereCipher(key)

}
