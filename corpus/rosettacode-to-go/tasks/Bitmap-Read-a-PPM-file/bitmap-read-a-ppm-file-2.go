package main

// Files required to build supporting package raster are found in:
// * This task (immediately above)
// * Bitmap
// * Grayscale image
// * Write a PPM file

import (
	"fmt"
	"raster"
)

func main() {
	// (A file with this name is output by the Go solution to the task
	// "Bitmap/Read an image through a pipe," but of course any 8-bit
	//  P6 PPM file should work.)
	b, err := raster.ReadPpmFile("pipein.ppm")
	if err != nil {
		fmt.Println(err)
		return
	}
	b = b.Grmap().Bitmap()
	err = b.WritePpmFile("grayscale.ppm")
	if err != nil {
		fmt.Println(err)
	}
}

//\Bitmap-Read-a-PPM-file\bitmap-read-a-ppm-file-2.go
