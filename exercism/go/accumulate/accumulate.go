package accumulate

// Accumulate apply operation to each element of input
func Accumulate(s []string, op func(string) string) []string {
	out := make([]string, 0, 1)
	for _, c := range s {
		out = append(out, op(c))
	}
	return out
}
