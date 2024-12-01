package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file_byte, _ := os.ReadFile("input.txt")
	file := string(file_byte)
	arr1 := []int{}
	arr2 := map[int]int{}
	for _, line := range strings.Split(file, "\n") {
		line := strings.TrimSpace(line)
		if line != "" {
			splitted := strings.Fields(line)
			if len(splitted) != 2 {
				panic("wrong amount of elems: ")
			}
			tmp1, _ := strconv.Atoi(splitted[0])
			arr1 = append(arr1, tmp1)
			tmp2, _ := strconv.Atoi(splitted[1])
			if val, ok := arr2[tmp2]; ok {
				arr2[tmp2] = val + 1
			} else {
				arr2[tmp2] = 1
			}
		}
	}
	res := 0

	for _, val1 := range arr1 {
		val2 := arr2[val1]
		res += val1 * val2
	}

	fmt.Println(res)
}
