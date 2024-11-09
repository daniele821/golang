package main

import (
	"exp/calc/fact"
	"fmt"
	"time"
)

func main() {
	i := 0x80_00_00_00 + 1
	start := time.Now()
	fmt.Println(i, fact.Factorize(i))
	fmt.Println(time.Since(start))
}
