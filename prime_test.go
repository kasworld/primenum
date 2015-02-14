package primenum

import (
	"testing"
)

func TestMakePrime(t *testing.T) {
	primes := MakePrimes(0xffff)
	t.Logf("%v", len(primes))
	t.Logf("%v", primes[len(primes)-50:])
}
