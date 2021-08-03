package hard

import (
	"fmt"
)

func sieve(upperBound int64) []int64 {
	_sieveSize := upperBound + 10

	var bs [10000010]bool

	primes := make([]int64, 0, 1000)

	bs[0] = true
	bs[1] = true

	for i := int64(0); i <= _sieveSize; i++ {
		if !bs[i] {
			for j := i * i; j <= _sieveSize; j += i {
				bs[j] = true
			}

			primes = append(primes, i)
		}
	}
	return primes
}

func main() {
	N := 100
	primes := sieve(N)
	fmt.Println(primes)
}