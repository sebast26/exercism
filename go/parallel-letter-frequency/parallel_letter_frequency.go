package letter

import "sync"

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

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(l []string) FreqMap {
	results := make(chan FreqMap)
	var n sync.WaitGroup
	for _, line := range l {
		n.Add(1)
		go concFreq(line, &n, results)
	}
	go func() {
		n.Wait()
		close(results)
	}()

	fm := FreqMap{}
	for r := range results {
		for k, v := range r {
			fm[k] += v
		}
	}

	return fm
}

func concFreq(line string, n *sync.WaitGroup, results chan<- FreqMap) {
	defer n.Done()
	results <- Frequency(line)
}
