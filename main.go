package main

import (
	"fmt"
	"slices"
)

func main() {
	test(1, 2)
}

func test(a ...int) {
	fmt.Println(slices.Contains(a, 1))
	fmt.Printf("%T", a)
}
