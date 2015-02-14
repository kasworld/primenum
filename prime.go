package primenum

import (
	"math"
)

func isprime(n int64, parr []int64) bool {
	to := int64(math.Sqrt(float64(n)))
	for _, v := range parr {
		if n%v == 0 {
			return false
		}
		if v > to {
			break
		}
	}
	return true
}

func MakePrimes(n int64) []int64 {
	foundPrimes := []int64{2}
	var i int64

	for i = 3; i < n; i += 2 {
		if isprime(i, foundPrimes) {
			foundPrimes = append(foundPrimes, i)
		}
	}
	return foundPrimes
}
