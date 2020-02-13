package space
import (
	"math"
)
type Planet string

/*
- Mercury: orbital period 0.2408467 Earth years
- Venus: orbital period 0.61519726 Earth years
- Earth: orbital period 1.0 Earth years, 365.25 Earth days, or 31557600 seconds
- Mars: orbital period 1.8808158 Earth years
- Jupiter: orbital period 11.862615 Earth years
- Saturn: orbital period 29.447498 Earth years
- Uranus: orbital period 84.016846 Earth years
- Neptune: orbital period 164.79132 Earth years
*/
var orbitals = map[Planet] float64{
		"Mercury": 0.2408467, 
		"Venus" : 0.61519726,
		"Earth": 1.0,
		"Mars": 1.8808158,
		"Jupiter": 11.862615 ,
		"Saturn": 29.447498,
		"Uranus": 84.016846,
		"Neptune": 164.79132,
	}

func 	Age(t float64 , planet Planet ) float64{
	age := t  /(31557600 * orbitals[planet] )
	age =  math.Floor((age * 100 ) + .5 ) / 100

	return age

}