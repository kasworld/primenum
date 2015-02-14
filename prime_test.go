package primenum

import (
	"testing"
)

func TestMakePrime(t *testing.T) {
	primes := MakePrimes(0xffff)
	t.Logf("%v", len(primes))
	t.Logf("%v", primes[len(primes)-50:])
}

func TestGetFactor(t *testing.T) {
	primes := MakePrimes(0xffff)
	for i := 0; i < 10; i++ {
		factors := GetFactor(primes, 50)
		t.Logf("factors %v", factors)
		for j := 0; j < 3; j++ {
			for k := j + 1; k < 3; k++ {
				if factors[j] == factors[k] {
					t.Fail()
				}
			}
		}
	}
}
