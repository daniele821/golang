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

func readersEqual(r io.Reader, t io.Reader) (bool, error) {
	rBuf := make([]byte, chunkSize)
	tBuf := make([]byte, chunkSize)

	for {
		readFromR, errR := r.Read(rBuf)
		if errR != nil && !errors.Is(errR, io.EOF) {
			return false, errR
		}

		readFromT := 0
		tCmpBuf := tBuf[:readFromR]

		if readFromR == 0 && errors.Is(errR, io.EOF) {
			readFromT, errT := t.Read(tBuf[:1])
			if readFromT == 0 && errors.Is(errT, io.EOF) {
				return true, nil
			} else {
				return false, errT
			}
		}

		for readFromR > readFromT {
			nextReadFromT, errT := t.Read(tCmpBuf[readFromT:])
			if errT != nil && !errors.Is(errT, io.EOF) {
				return false, errT
			}
			prevReadFromT := readFromT
			readFromT = prevReadFromT + nextReadFromT
			if !bytes.Equal(rBuf[prevReadFromT:readFromT], tCmpBuf[prevReadFromT:readFromT]) {
				return false, nil
			}
			if errors.Is(errR, io.EOF) && errors.Is(errT, io.EOF) {
				return true, nil
			}
			if errors.Is(errR, io.EOF) || errors.Is(errT, io.EOF) {
				return false, nil
			}
		}
	}
}
func main() {
	start := time.Now()
	a, _ := os.Open("/home/daniele/.local/share/fonts/FiraCode/FiraCodeNerdFont-Bold.ttf")
	b, _ := os.Open("/home/daniele/.local/share/fonts/FiraCode/FiraCodeNerdFont-Bold.ttf")
	fmt.Println(readersEqual(a, b))
	fmt.Println(time.Since(start))
}
