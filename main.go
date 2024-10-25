package main

import (
	"flag"
	"fmt"
	"io/fs"
	"path/filepath"
)

func visit(path string, di fs.DirEntry, err error) error {
	if err != nil {
		fmt.Println(err)
		panic("FUUUUUCK")
	}
	fmt.Println(path, " | ", di, " | ", err)
	return nil
}

func main() {
	flag.Parse()
	root := flag.Arg(0)
	err := filepath.WalkDir(root, visit)
	fmt.Printf("filepath.WalkDir() returned %v\n", err)
}
