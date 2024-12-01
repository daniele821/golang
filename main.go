package main

import (
	"fmt"
	"os"
)

func main() {
	file_byte, _ := os.ReadFile("input.txt")
	file := string(file_byte)
	fmt.Println(file)
}
