package main

import (
	"fmt"
	"slices"
)

func main() {
	slice := []string{"a", "b", "b", "a"}
	slices.Sort(slice)
	slice = slices.Compact(slice)
	fmt.Println(slice)
}
