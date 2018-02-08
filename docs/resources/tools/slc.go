package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	searchDir := os.Args[1]

	err := filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		if f.IsDir() {
			fmt.Println(f.Name(), path)
		}

		return err
	})

	if err != nil {
		panic(err)
	}

	//for _, file := range fileList {
	//
	//	}
}
