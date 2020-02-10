package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func worker(s string, FreqMap chan <- FreqMap)  {
	FreqMap <- Frequency(s) 
}

func ConcurrentFrequency(s []string) FreqMap {
	freqMapChannel := make(chan FreqMap)
	m := FreqMap{}
	for _, t := range s {
		go worker(t, freqMapChannel)
	}
	for i := 0 ; i < len(s);  i++ {
		mx := <- freqMapChannel
		for k, v := range mx {
			m[k] += v
		}
		
	}
	return m
}