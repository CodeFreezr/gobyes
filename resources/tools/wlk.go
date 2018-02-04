package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

type gofi struct {
	Name    string
	Path    string
	User    string
	Size    int64
	Mode    os.FileMode
	ModTime time.Time
	IsDir   bool
}

func main() {
	searchDir := os.Args[1]
	ext := ".go"
	//fileList := make([]string, 0)

	fileList := []gofi{}
	err := filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		if !f.IsDir() {
			reg, err := regexp.MatchString(ext, f.Name())
			if err == nil && reg {
				r := strings.Split(path, ".")[0]
				//fileList = append(fileList, path)
				fileList = append(fileList, gofi{f.Name(), path, r, f.Size(), f.Mode(), f.ModTime(), f.IsDir()})
			}
		} else {
			//fmt.Println(f.Name(), path)
		}

		return err
	})

	if err != nil {
		panic(err)
	}

	for _, file := range fileList {
		fmt.Println("File: ", file.Name)
		fmt.Println("Path: ", file.Path)
		fmt.Println("User: ", file.User)

	}
}
