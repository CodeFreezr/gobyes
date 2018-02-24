package main

import "os"

func main() {
	os.Rename("input.txt", "output.txt")
	os.Rename("docs", "mydocs")
	os.Rename("/input.txt", "/output.txt")
	os.Rename("/docs", "/mydocs")
}

//\Rename-a-file\rename-a-file.go
