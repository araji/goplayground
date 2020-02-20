package flatten

//Flatten flatten slices
func Flatten(sin interface{}) []interface{} {
	sout := make([]interface{}, 0, 1)
	for _, val := range sin.([]interface{}) {
		if v, ok := val.([]interface{}); ok {
			sout = append(sout, Flatten(v)...)
		} else {
			if val != nil {
				sout = append(sout, val)
			}
		}
	}
	return sout
}
