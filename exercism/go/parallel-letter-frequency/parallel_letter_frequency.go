package letter
import (

	"sync"

)

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int
var lock = sync.RWMutex{}
var mc = FreqMap{}



// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func worker(s string, done chan bool)  {
	lock.Lock()
	defer lock.Unlock()
	for _, r := range s {
		mc[r]++
	}
	done <- true
}

func ConcurrentFrequency(s []string) FreqMap {
	//freqMapChannel := make(chan FreqMap)
	done := make(chan bool)
	for _, t := range s {
		go worker(t, done)
	}
	for  i:=0 ; i< len(s) ; i ++ {
		<- done
	}
	return mc
}