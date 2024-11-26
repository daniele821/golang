package main

import (
	"fmt"
	"os"
)

func main() {
	i, err := os.ReadFile("/home/daniele/test")
	fmt.Println(err)
	fmt.Println((string(i)))
}
