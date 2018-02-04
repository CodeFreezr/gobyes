// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bufio

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"bufio"
	"io"
	"os"
)

var _ = map[string]string{
	"Reader":     "*bufio.Reader",     // func NewReader(w io.Reader) *Reader
	"":           "",                  // func NewReaderSize(w io.Reader, size int) *Reader
	"ReadWriter": "*bufio.ReadWriter", // func NewReadWriter(r *Reader, w *Writer) *ReadWriter
	"Scanner":    "*bufio.Scanner",    // func NewScanner(r io.Reader) *Scanner
	"SplitFunc":  "*bufio.SplitFunc",  // type SplitFunc func(data []byte, atEOF bool) (advance int, token []byte, err error)
	"Writer":     "*bufio.Writer",     // func NewWriter(w io.Writer) *Writer
	".":          "",                  // func NewWriterSize(w io.Writer, size int) *Writer
}

// Readers - (w io.Reader) *Reader
func Readers(inp <-chan io.Reader) (out <-chan *bufio.Reader) {
	cha := make(chan *bufio.Reader)
	go func(inp <-chan io.Reader, out chan<- *bufio.Reader) {
		defer close(out)
		for i := range inp {
			out <- bufio.NewReader(i)
		}
	}(inp, cha)
	return cha
}

// ReaderSize - (w io.Reader, size int) *Reader
func ReaderSize(inp <-chan io.Reader, size int) (out <-chan *bufio.Reader) {
	cha := make(chan *bufio.Reader)
	go func(inp <-chan io.Reader, size int, out chan<- *bufio.Reader) {
		defer close(out)
		for i := range inp {
			out <- bufio.NewReaderSize(i, size)
		}
	}(inp, size, cha)
	return cha
}

// ReadWriters - (r *Reader, w *Writer) *ReadWriter
func ReadWriters(inp1 <-chan *bufio.Reader, inp2 <-chan *bufio.Writer) (out <-chan *bufio.ReadWriter) {
	cha := make(chan *bufio.ReadWriter)
	go func(inp1 <-chan *bufio.Reader, inp2 <-chan *bufio.Writer, out chan<- *bufio.ReadWriter) {
		defer close(out)
		for i1 := range inp1 {
			if i2, ok := <-inp2; ok {
				out <- bufio.NewReadWriter(i1, i2)
			} else {
				break
			}
		}
	}(inp1, inp2, cha)
	return cha
}

// Scanners - (r io.Reader) *Scanner
func Scanners(inp <-chan io.Reader) (out <-chan *bufio.Scanner) {
	cha := make(chan *bufio.Scanner)
	go func(inp <-chan io.Reader, out chan<- *bufio.Scanner) {
		defer close(out)
		for i := range inp {
			out <- bufio.NewScanner(i)
		}
	}(inp, cha)
	return cha
}

// Writers - (w io.Writer) *Writer
func Writers(inp <-chan io.Writer) (out <-chan *bufio.Writer) {
	cha := make(chan *bufio.Writer)
	go func(inp <-chan io.Writer, out chan<- *bufio.Writer) {
		defer close(out)
		for i := range inp {
			out <- bufio.NewWriter(i)
		}
	}(inp, cha)
	return cha
}

// WriterSize - (w io.Writer, size int) *Writer
func WriterSize(inp <-chan io.Writer, size int) (out <-chan *bufio.Writer) {
	cha := make(chan *bufio.Writer)
	go func(inp <-chan io.Writer, size int, out chan<- *bufio.Writer) {
		defer close(out)
		for i := range inp {
			out <- bufio.NewWriterSize(i, size)
		}
	}(inp, size, cha)
	return cha
}

// PipeLines - an experiment
func PipeLines(inp <-chan os.File) (out <-chan string) {
	cha := make(chan string)
	go func(inp <-chan os.File, out chan<- string) {
		defer close(out)
		var s *bufio.Scanner
		for file := range inp {
			f, _ := os.Open(file.Name())
			s = bufio.NewScanner(f)
			for s.Scan() {
				out <- s.Text()
			}
			f.Close()
		}
	}(inp, cha)
	return cha
}
