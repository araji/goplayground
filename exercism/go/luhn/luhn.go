package luhn

// Valid determines whether the digits in the given string constitute a valid luhn code
// Copied from leaderboard as this is 50 times faster than my solution .Embarrassing :-).
func Valid(s string) bool {
	var n, d, i, m int
	for i = len(s) - 1; i >= 0; i-- {
		c := s[i]
		switch {
		case c == ' ':
			continue
		case c >= '0' && c <= '9':
			m = int(c - '0')
			if d%2 == 1 {
				m <<= 1
			}
			if m > 9 {
				m -= 9
			}
			n += m
			d++
		default:
			return false
		}
	}
	return d > 1 && n%10 == 0
}