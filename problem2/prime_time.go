package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	if n <= 3 {
		return n > 1
	}
	if (n%2 == 0) || (n%3 == 0) {
		return false
	}
	i := 5
	stop := int(math.Sqrt(float64(n)))
	for i <= stop {
		if (n%i == 0) || (n%(i+2) == 0) {
			return false
		}
		i += 6
	}
	return true
}

func main() {
	isPrime(0)

  listener, err := net.Listen("tcp", ":9000")
}
