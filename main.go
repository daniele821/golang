package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file_byte, _ := os.ReadFile("input.txt")
	file := string(file_byte)
	arr1 := []string{}
	arr2 := []string{}
	for _, line := range strings.Split(file, "\n") {
		line := strings.TrimSpace(line)
		if line != "" {
			splitted := strings.Fields(line)
			if len(splitted) != 2 {
				panic("wrong amount of elems: ")
			}
			arr1 = append(arr1, splitted[0])
			arr2 = append(arr2, splitted[1])
		}
	}
	slices.Sort(arr1)
	slices.Sort(arr2)
	res := 0
	for index, elem1 := range arr1 {
		elem2 := arr2[index]
		elem1_int, _ := strconv.Atoi(elem1)
		elem2_int, _ := strconv.Atoi(elem2)
		tmp := elem2_int - elem1_int
		if tmp < 0 {
			tmp *= -1
		}
		res += tmp
	}
	fmt.Println(res)
}
