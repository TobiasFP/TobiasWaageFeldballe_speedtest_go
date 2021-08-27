package speedtest

//go:generate gopherjs build --minify

// This is an experiment to see if gopherjs can reasonably generate js code from go source
// so that we can have a single-source solution for keys and addresses.
// Use "go generate" to build this.

import (
	"fmt"
)

// Taken from http://www.hildstrom.com/projects/langcomp/index.html
func SpeedTest(iterations int) float64 {
	var (
		element      int       = 0
		iteration    int       = 0
		innerloop    int       = 0
		sum          float64   = 0.0
		array_length int       = 100000000
		array        []float64 = make([]float64, array_length)
	)
	for element = 0; element < array_length; element++ {
		array[element] = float64(element)
	}
	for iteration = 0; iteration < iterations; iteration++ {
		for innerloop = 0; innerloop < array_length; innerloop++ {
			sum += array[(iteration+innerloop)%array_length]
		}
	}
	fmt.Println(sum)
	array = nil
	return sum
}
