package main

import (
	"exp/calc/fact"
	"fmt"
)

func main() {
	i := 0xff_ff_ff_ff
	fmt.Println(i, fact.Factorize(i))
}
