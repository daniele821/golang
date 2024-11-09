package fact

import (
	"fmt"
	"math/bits"
	"strings"
)

type PrimeFactors map[int]int

func Factorize(num int) PrimeFactors {
	if num < 2 {
		return nil
	}
	factors := map[int]int{}

	// fast calculate power of 2
	powTwo := bits.TrailingZeros(uint(num))
	if powTwo > 0 {
		num >>= powTwo
		factors[2] = powTwo
	}

	// calculate all odd factors
	// no need to check if divisor is prime:
	// if divisor isn't prime, its prime factors
	// are smaller and already included
	divisor := 3
	for num > 1 {
		if num%divisor == 0 {
			factors[divisor] += 1
			num /= divisor
		} else {
			divisor += 2
		}
	}

	return factors
}

func (p PrimeFactors) String() string {
	builder := []string{}
	for key, val := range p {
		if val == 1 {
			builder = append(builder, fmt.Sprintf("%d", key))
		} else {
			builder = append(builder, fmt.Sprintf("%d^%d", key, val))
		}
	}
	return strings.Join(builder, " * ")
}
