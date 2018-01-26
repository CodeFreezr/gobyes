package main

import (
	"os"
	"path/filepath"
	"regexp"
)

func main() {
	searchDir := "D:/dbt/01/git/gobyes/ae6rt.golang-examples/goeg"
	ext := ".go"
	fileList := make([]string, 0)

	//fmt.Println("Directories: ----------------------")
	err := filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		if !f.IsDir() {
			reg, err := regexp.MatchString(ext, f.Name())
			if err == nil && reg {
				fileList = append(fileList, path)
			}
		} else {
			//fmt.Println(f.Name(), path)
		}

		return err
	})

	if err != nil {
		panic(err)
	}

	/*
		fmt.Println("Files: ----------------------")

		for _, file := range fileList {

			fmt.Println(file)
		}
	*/
	//return fileList, nil
}
