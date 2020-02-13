// Package leap should have a package comment that summarizes what it's about.
package leap

// IsLeapYear test if provided year is a leap year.
/**
on every year that is evenly divisible by 4
  except every year that is evenly divisible by 100
	unless the year is also evenly divisible by 400
*/
func IsLeapYear(year int ) bool {
	if ! (year % 4 == 0 ) {
		return false
	}
	if (  (year % 100 ==0 ) && ! (year % 400 == 0 )) {
		return false
	}

	return true
}
