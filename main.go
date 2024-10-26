package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"time"
)

func readersEqual(file1, file2 string) bool {
	const chunkSize = 64000
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

func FileCmp(file1, file2 string) bool {
	const chunkSize int = 4096

	// shortcuts: check file metadata
	stat1, err := os.Stat(file1)
	if err != nil {
		return false
	}

	stat2, err := os.Stat(file2)
	if err != nil {
		return false
	}

	// are inputs are literally the same file?
	if os.SameFile(stat1, stat2) {
		return true
	}

	// do inputs at least have the same size?
	if stat1.Size() != stat2.Size() {
		return false
	}

	// long way: compare contents
	f1, err := os.Open(file1)
	if err != nil {
		return false
	}
	defer f1.Close()

	f2, err := os.Open(file2)
	if err != nil {
		return false
	}
	defer f2.Close()

	b1 := make([]byte, chunkSize)
	b2 := make([]byte, chunkSize)
	for {
		n1, err1 := io.ReadFull(f1, b1)
		n2, err2 := io.ReadFull(f2, b2)

		// https://pkg.go.dev/io#Reader
		// > Callers should always process the n > 0 bytes returned
		// > before considering the error err. Doing so correctly
		// > handles I/O errors that happen after reading some bytes
		// > and also both of the allowed EOF behaviors.

		if !bytes.Equal(b1[:n1], b2[:n2]) {
			return false
		}

		if (err1 == io.EOF && err2 == io.EOF) || (err1 == io.ErrUnexpectedEOF && err2 == io.ErrUnexpectedEOF) {
			return true
		}

		// some other error, like a dropped network connection or a bad transfer
		if err1 != nil {
			return false
		}
		if err2 != nil {
			return false
		}
	}
}

func main() {
	a := "/home/daniele/.local/share/fonts/FiraCode/FiraCodeNerdFont-Light.ttf"
	b := "/personal/repos/daniele821/dotfiles/backup/.local/share/fonts/FiraCode/FiraCodeNerdFont-Bold.ttf"
	start := time.Now()
	fmt.Print(readersEqual(a, b))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Print(FileCmp(a, b))
	fmt.Println(time.Since(start))
}
