package main

import "fmt"

var LIMIT int = 5

func iter(i *int, sum int) {
	*i = (*i+sum+LIMIT-1)%LIMIT + 1
}

func main() {
	i := 1
	for range 20 {
		fmt.Print(i, " ")
		iter(&i, 3)
	}
	fmt.Println()
	for range 20 {
		fmt.Print(i, " ")
		iter(&i, -2)
	}
	fmt.Println()
}
