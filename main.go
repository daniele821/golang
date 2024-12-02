package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func isValid(vals []int) bool {
	length := len(vals)
	if length <= 1 {
		return true
	}
	direction := vals[1] - vals[0]
	decreasing := false
	if direction < 0 {
		decreasing = true
	}
	for index := range length - 1 {
		prev := vals[index]
		next := vals[index+1]
		dir := next - prev
		if decreasing {
			dir *= -1
		}
		if dir < 1 || dir > 3 {
			return false
		}
	}
	return true
}

func main() {
	file_byte, _ := os.ReadFile("input2.txt")
	file := string(file_byte)
	count := 0
	for _, line := range strings.Split(file, "\n") {
		line := strings.TrimSpace(line)
		if line == "" {
			continue
		}
		splitted := strings.Fields(line)
		vals := []int{}
		for _, elem := range splitted {
			value, _ := strconv.Atoi(elem)
			vals = append(vals, value)
		}
		if isValid(vals) {
			fmt.Println(vals)
			count += 1
		} else {
			for index := range len(vals) {
				if isValid(slices.Concat(vals[:index], vals[index+1:])) {
					fmt.Println(slices.Concat(vals[:index], vals[index+1:]))
					count += 1
					break
				}
			}
		}
	}

	fmt.Println(count)
}
