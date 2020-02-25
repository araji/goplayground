package etl

import (
	"strings"
)

func Transform(input map[int][]string )  (out map[string]int ) {
	out = make(map[string]int,0)
	if len(input) == 0 {
		return nil

	}
	for k,vslice := range input {
		for _,v := range vslice {
			out[strings.ToLower(v)] =k
		}
	}
	return
}	
