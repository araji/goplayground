package grains

import (
	"errors"
)

func Square(i int) (uint64, error) {
	
	if i <= 0  || i > 64 {
		return 0,errors.New("cannot use for square")
	} 
	return ( 1<< uint(i-1) ) , nil
	
}

func Total() uint64 {
//	if false {}
	return ( 1<<64 ) -1 
}
