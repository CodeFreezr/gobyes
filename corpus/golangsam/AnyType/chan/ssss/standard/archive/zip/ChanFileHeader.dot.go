// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package zip

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	zip "archive/zip"
)

// MakeFileHeaderChan returns a new open channel
// (simply a 'chan *zip.FileHeader' that is).
//
// Note: No 'FileHeader-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myFileHeaderPipelineStartsHere := MakeFileHeaderChan()
//	// ... lot's of code to design and build Your favourite "myFileHeaderWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myFileHeaderPipelineStartsHere <- drop
//	}
//	close(myFileHeaderPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeFileHeaderBuffer) the channel is unbuffered.
//
func MakeFileHeaderChan() chan *zip.FileHeader {
	return make(chan *zip.FileHeader)
}

// ChanFileHeader returns a channel to receive all inputs before close.
func ChanFileHeader(inp ...*zip.FileHeader) chan *zip.FileHeader {
	out := make(chan *zip.FileHeader)
	go func() {
		defer close(out)
		for i := range inp {
			out <- inp[i]
		}
	}()
	return out
}

// ChanFileHeaderSlice returns a channel to receive all inputs before close.
func ChanFileHeaderSlice(inp ...[]*zip.FileHeader) chan *zip.FileHeader {
	out := make(chan *zip.FileHeader)
	go func() {
		defer close(out)
		for i := range inp {
			for j := range inp[i] {
				out <- inp[i][j]
			}
		}
	}()
	return out
}

// ChanFileHeaderFuncNil returns a channel to receive all results of act until nil before close.
func ChanFileHeaderFuncNil(act func() *zip.FileHeader) <-chan *zip.FileHeader {
	out := make(chan *zip.FileHeader)
	go func() {
		defer close(out)
		for {
			res := act() // Apply action
			if res == nil {
				return
			}
			out <- res
		}
	}()
	return out
}

// ChanFileHeaderFuncNok returns a channel to receive all results of act until nok before close.
func ChanFileHeaderFuncNok(act func() (*zip.FileHeader, bool)) <-chan *zip.FileHeader {
	out := make(chan *zip.FileHeader)
	go func() {
		defer close(out)
		for {
			res, ok := act() // Apply action
			if !ok {
				return
			}
			out <- res
		}
	}()
	return out
}

// ChanFileHeaderFuncErr returns a channel to receive all results of act until err != nil before close.
func ChanFileHeaderFuncErr(act func() (*zip.FileHeader, error)) <-chan *zip.FileHeader {
	out := make(chan *zip.FileHeader)
	go func() {
		defer close(out)
		for {
			res, err := act() // Apply action
			if err != nil {
				return
			}
			out <- res
		}
	}()
	return out
}

// JoinFileHeader sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinFileHeader(out chan<- *zip.FileHeader, inp ...*zip.FileHeader) chan struct{} {
	done := make(chan struct{})
	go func() {
		defer close(done)
		for i := range inp {
			out <- inp[i]
		}
		done <- struct{}{}
	}()
	return done
}

// JoinFileHeaderSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinFileHeaderSlice(out chan<- *zip.FileHeader, inp ...[]*zip.FileHeader) chan struct{} {
	done := make(chan struct{})
	go func() {
		defer close(done)
		for i := range inp {
			for j := range inp[i] {
				out <- inp[i][j]
			}
		}
		done <- struct{}{}
	}()
	return done
}

// JoinFileHeaderChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinFileHeaderChan(out chan<- *zip.FileHeader, inp <-chan *zip.FileHeader) chan struct{} {
	done := make(chan struct{})
	go func() {
		defer close(done)
		for i := range inp {
			out <- i
		}
		done <- struct{}{}
	}()
	return done
}

// DoneFileHeader returns a channel to receive one signal before close after inp has been drained.
func DoneFileHeader(inp <-chan *zip.FileHeader) chan struct{} {
	done := make(chan struct{})
	go func() {
		defer close(done)
		for i := range inp {
			_ = i // Drain inp
		}
		done <- struct{}{}
	}()
	return done
}

// DoneFileHeaderSlice returns a channel which will receive a slice
// of all the FileHeaders received on inp channel before close.
// Unlike DoneFileHeader, a full slice is sent once, not just an event.
func DoneFileHeaderSlice(inp <-chan *zip.FileHeader) chan []*zip.FileHeader {
	done := make(chan []*zip.FileHeader)
	go func() {
		defer close(done)
		FileHeaderS := []*zip.FileHeader{}
		for i := range inp {
			FileHeaderS = append(FileHeaderS, i)
		}
		done <- FileHeaderS
	}()
	return done
}

// DoneFileHeaderFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneFileHeaderFunc(inp <-chan *zip.FileHeader, act func(a *zip.FileHeader)) chan struct{} {
	done := make(chan struct{})
	if act == nil {
		act = func(a *zip.FileHeader) { return }
	}
	go func() {
		defer close(done)
		for i := range inp {
			act(i) // Apply action
		}
		done <- struct{}{}
	}()
	return done
}

// PipeFileHeaderBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeFileHeaderBuffer(inp <-chan *zip.FileHeader, cap int) chan *zip.FileHeader {
	out := make(chan *zip.FileHeader, cap)
	go func() {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}()
	return out
}

// PipeFileHeaderFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeFileHeaderMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeFileHeaderFunc(inp <-chan *zip.FileHeader, act func(a *zip.FileHeader) *zip.FileHeader) chan *zip.FileHeader {
	out := make(chan *zip.FileHeader)
	if act == nil {
		act = func(a *zip.FileHeader) *zip.FileHeader { return a }
	}
	go func() {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}()
	return out
}

// PipeFileHeaderFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeFileHeaderFork(inp <-chan *zip.FileHeader) (chan *zip.FileHeader, chan *zip.FileHeader) {
	out1 := make(chan *zip.FileHeader)
	out2 := make(chan *zip.FileHeader)
	go func() {
		defer close(out1)
		defer close(out2)
		for i := range inp {
			out1 <- i
			out2 <- i
		}
	}()
	return out1, out2
}

// FileHeaderTube is the signature for a pipe function.
type FileHeaderTube func(inp <-chan *zip.FileHeader, out <-chan *zip.FileHeader)

// FileHeaderDaisy returns a channel to receive all inp after having passed thru tube.
func FileHeaderDaisy(inp <-chan *zip.FileHeader, tube FileHeaderTube) (out <-chan *zip.FileHeader) {
	cha := make(chan *zip.FileHeader)
	go tube(inp, cha)
	return cha
}

// FileHeaderDaisyChain returns a channel to receive all inp after having passed thru all tubes.
func FileHeaderDaisyChain(inp <-chan *zip.FileHeader, tubes ...FileHeaderTube) (out <-chan *zip.FileHeader) {
	cha := inp
	for i := range tubes {
		cha = FileHeaderDaisy(cha, tubes[i])
	}
	return cha
}

/*
func sendOneInto(snd chan<- int) {
	defer close(snd)
	snd <- 1 // send a 1
}

func sendTwoInto(snd chan<- int) {
	defer close(snd)
	snd <- 1 // send a 1
	snd <- 2 // send a 2
}

var fun = func(left chan<- int, right <-chan int) { left <- 1 + <-right }

func main() {
	leftmost := make(chan int)
	right := daisyChain(leftmost, fun, 10000) // the chain - right to left!
	go sendTwoInto(right)
	fmt.Println(<-leftmost)
}
*/
