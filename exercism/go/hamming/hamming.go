package hamming

import (
	"errors"
)

var HammingError = errors.New("Hamming Error")

func Distance(a, b string) (int, error) {
	mismatch := 0 
	
	if ( len(a) != len(b) ) {
		return mismatch, HammingError
	}
	
	for ndx  := range a {
		if a[ndx] != b[ndx] {
			mismatch += 1
		}
	}
	return mismatch, nil
}
