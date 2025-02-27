package crypto

import (
	"fmt"
	"slices"
)

func LinGen(x0, a, b, m int) int {
	return (x0*a + b) % m
}

func FindPeriod(x0, a, b, m int) []int {
	period := []int{}

	index := 0

	for {
		x1 := LinGen(x0, a, b, m)
		if slices.Contains(period, x1) {
			return period
		}
		period = append(period, x1)
		fmt.Printf("%d -> %d\n", index, x1)
		index += 1
	}
}
