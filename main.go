package main

import (
	"exp/internal/calc"
	"fmt"
	"time"
)

func main() {
	i := 0x80_00_00_00 + 1
	start := time.Now()
	fmt.Println(i, calc.Factorize(i))
	fmt.Println(time.Since(start))
}
