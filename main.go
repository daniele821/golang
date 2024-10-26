package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"time"
)

const chunkSize = 64000

func readersEqual(file1, file2 string) bool {
	r, _ := os.Open(file1)
	t, _ := os.Open(file2)
	rBuf := make([]byte, chunkSize)
	tBuf := make([]byte, chunkSize)

	for {
		readFromR, errR := r.Read(rBuf)
		if errR != nil && !errors.Is(errR, io.EOF) {
			return false
		}

		readFromT := 0
		tCmpBuf := tBuf[:readFromR]

		if readFromR == 0 && errors.Is(errR, io.EOF) {
			readFromT, errT := t.Read(tBuf[:1])
			if readFromT == 0 && errors.Is(errT, io.EOF) {
				return true
			} else {
				return false
			}
		}

		for readFromR > readFromT {
			nextReadFromT, errT := t.Read(tCmpBuf[readFromT:])
			if errT != nil && !errors.Is(errT, io.EOF) {
				return false
			}
			prevReadFromT := readFromT
			readFromT = prevReadFromT + nextReadFromT
			if !bytes.Equal(rBuf[prevReadFromT:readFromT], tCmpBuf[prevReadFromT:readFromT]) {
				return false
			}
			if errors.Is(errR, io.EOF) && errors.Is(errT, io.EOF) {
				return true
			}
			if errors.Is(errR, io.EOF) || errors.Is(errT, io.EOF) {
				return false
			}
		}
	}
}

func main() {
	start := time.Now()
	a := "/home/daniele/.local/share/fonts/FiraCode/FiraCodeNerdFont-Bold.ttf"
	b := "/home/daniele/.local/share/fonts/FiraCode/FiraCodeNerdFont-Bold.ttf"
	fmt.Println(readersEqual(a, b))
	fmt.Println(time.Since(start))
}
