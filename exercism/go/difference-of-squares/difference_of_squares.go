package diffsquares

func SquareOfSum(n int) int {
	sos := 0
	for i:=0 ; i<=n ; i ++ {
		sos += i 
	}
	return sos*sos
}
func SumOfSquares(n int) int {
	sos := 0
	for i:=0 ; i<=n ; i ++ {
		sos += i *i
	}
	return sos
}

func Difference(n int) int {
//	return SquareOfSum(n) - SumOfSquares(n )
	s1 := 0
	s2 := 0
	for i:=0 ; i<=n ; i ++ {
		s1 += i 
		s2 += i*i
	}

	return (s1*s1) -s2
}

//very nice : https://exercism.io/tracks/go/exercises/difference-of-squares/solutions/289541bbda34fc1bdb5c099b
//excellent : https://exercism.io/tracks/go/exercises/difference-of-squares/solutions/17c2257af6cb445d9aa5da1a3ccf4c3a
