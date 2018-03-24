package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

type messyfolders struct {
	Path    string
	Size    int64
	Mode    os.FileMode
	ModTime time.Time
	IsDir   bool
}

func main() {
	searchDir := os.Args[1]

	fileList := []messyfolders{}
	err := filepath.Walk(searchDir, func(path string, finfo os.FileInfo, err error) error {
		fileList = append(fileList, messyfolders{path, finfo.Size(), finfo.Mode(), finfo.ModTime(), finfo.IsDir()})
		return nil
	})

	if err != nil {
		fmt.Println(err)
	}

	output, err := json.Marshal(fileList)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(output))
}
