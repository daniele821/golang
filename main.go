package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func AllFilesInDir(dir, relPath string) []string {
	var files []string
	visit := func(path string, dirEntry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		fmt.Println(path)
		if dirEntry.Type().IsRegular() || dirEntry.Type()&os.ModeSymlink == os.ModeSymlink {
			if relPath != "" {
				path, err = filepath.Rel(relPath, path)
				if err != nil {
					return err
				}
			}
			files = append(files, path)
		}
		return nil
	}
	err := filepath.WalkDir(dir, visit)
	if err != nil {
		return []string{}
	}
	return files
}

func main() {
	fmt.Println(AllFilesInDir("/home/daniele/.config/kitty", ""))
}
