package main

import "fmt"

func print(box [3][3]int) {
	for _, line := range box {
		for _, cell := range line {
			fmt.Print(cell, " ")
		}
		fmt.Println()
	}
}

func oneline(box [3][3]int) [9]int {
	boxLine := [9]int{}
	for a, line := range box {
		for b, cell := range line {
			boxLine[a*3+b] = cell
		}
	}
	return boxLine
}

func main() {
	box := [3][3]int{}
	print(box)
}
