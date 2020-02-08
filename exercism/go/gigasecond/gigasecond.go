// gigasecond package one giasecond later solver
package gigasecond

// import path for the time package from the standard library
import "time"

// oneGigasecond in time.Duration
const onegigasecond = time.Duration(1000000000000000000)

// AddGigasecond takes time as input and returns time plus one gigasecond
func AddGigasecond(t time.Time) time.Time {
	return t.Add(onegigasecond)
}
