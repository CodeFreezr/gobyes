// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package zip

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	zip "archive/zip"
)

// MakeFileChan returns a new open channel
// (simply a 'chan *zip.File' that is).
//
// Note: No 'File-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myFilePipelineStartsHere := MakeFileChan()
//	// ... lot's of code to design and build Your favourite "myFileWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myFilePipelineStartsHere <- drop
//	}
//	close(myFilePipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeFileBuffer) the channel is unbuffered.
//
func MakeFileChan() (out chan *zip.File) {
	return make(chan *zip.File)
}

// ChanFile returns a channel to receive all inputs before close.
func ChanFile(inp ...*zip.File) (out <-chan *zip.File) {
	cha := make(chan *zip.File)
	go func(out chan<- *zip.File, inp ...*zip.File) {
		defer close(out)
		for i := range inp {
			out <- inp[i]
		}
	}(cha, inp...)
	return cha
}

// ChanFileSlice returns a channel to receive all inputs before close.
func ChanFileSlice(inp ...[]*zip.File) (out <-chan *zip.File) {
	cha := make(chan *zip.File)
	go func(out chan<- *zip.File, inp ...[]*zip.File) {
		defer close(out)
		for i := range inp {
			for j := range inp[i] {
				out <- inp[i][j]
			}
		}
	}(cha, inp...)
	return cha
}

// ChanFileFuncNil returns a channel to receive all results of act until nil before close.
func ChanFileFuncNil(act func() *zip.File) (out <-chan *zip.File) {
	cha := make(chan *zip.File)
	go func(out chan<- *zip.File, act func() *zip.File) {
		defer close(out)
		for {
			res := act() // Apply action
			if res == nil {
				return
			}
			out <- res
		}
	}(cha, act)
	return cha
}

// ChanFileFuncNok returns a channel to receive all results of act until nok before close.
func ChanFileFuncNok(act func() (*zip.File, bool)) (out <-chan *zip.File) {
	cha := make(chan *zip.File)
	go func(out chan<- *zip.File, act func() (*zip.File, bool)) {
		defer close(out)
		for {
			res, ok := act() // Apply action
			if !ok {
				return
			}
			out <- res
		}
	}(cha, act)
	return cha
}

// ChanFileFuncErr returns a channel to receive all results of act until err != nil before close.
func ChanFileFuncErr(act func() (*zip.File, error)) (out <-chan *zip.File) {
	cha := make(chan *zip.File)
	go func(out chan<- *zip.File, act func() (*zip.File, error)) {
		defer close(out)
		for {
			res, err := act() // Apply action
			if err != nil {
				return
			}
			out <- res
		}
	}(cha, act)
	return cha
}

// JoinFile sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinFile(out chan<- *zip.File, inp ...*zip.File) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- *zip.File, inp ...*zip.File) {
		defer close(done)
		for i := range inp {
			out <- inp[i]
		}
		done <- struct{}{}
	}(cha, out, inp...)
	return cha
}

// JoinFileSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinFileSlice(out chan<- *zip.File, inp ...[]*zip.File) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- *zip.File, inp ...[]*zip.File) {
		defer close(done)
		for i := range inp {
			for j := range inp[i] {
				out <- inp[i][j]
			}
		}
		done <- struct{}{}
	}(cha, out, inp...)
	return cha
}

// JoinFileChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinFileChan(out chan<- *zip.File, inp <-chan *zip.File) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- *zip.File, inp <-chan *zip.File) {
		defer close(done)
		for i := range inp {
			out <- i
		}
		done <- struct{}{}
	}(cha, out, inp)
	return cha
}

// DoneFile returns a channel to receive one signal before close after inp has been drained.
func DoneFile(inp <-chan *zip.File) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, inp <-chan *zip.File) {
		defer close(done)
		for i := range inp {
			_ = i // Drain inp
		}
		done <- struct{}{}
	}(cha, inp)
	return cha
}

// DoneFileSlice returns a channel which will receive a slice
// of all the Files received on inp channel before close.
// Unlike DoneFile, a full slice is sent once, not just an event.
func DoneFileSlice(inp <-chan *zip.File) (done <-chan []*zip.File) {
	cha := make(chan []*zip.File)
	go func(inp <-chan *zip.File, done chan<- []*zip.File) {
		defer close(done)
		FileS := []*zip.File{}
		for i := range inp {
			FileS = append(FileS, i)
		}
		done <- FileS
	}(inp, cha)
	return cha
}

// DoneFileFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneFileFunc(inp <-chan *zip.File, act func(a *zip.File)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a *zip.File) { return }
	}
	go func(done chan<- struct{}, inp <-chan *zip.File, act func(a *zip.File)) {
		defer close(done)
		for i := range inp {
			act(i) // Apply action
		}
		done <- struct{}{}
	}(cha, inp, act)
	return cha
}

// PipeFileBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeFileBuffer(inp <-chan *zip.File, cap int) (out <-chan *zip.File) {
	cha := make(chan *zip.File, cap)
	go func(out chan<- *zip.File, inp <-chan *zip.File) {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}(cha, inp)
	return cha
}

// PipeFileFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeFileMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeFileFunc(inp <-chan *zip.File, act func(a *zip.File) *zip.File) (out <-chan *zip.File) {
	cha := make(chan *zip.File)
	if act == nil {
		act = func(a *zip.File) *zip.File { return a }
	}
	go func(out chan<- *zip.File, inp <-chan *zip.File, act func(a *zip.File) *zip.File) {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}(cha, inp, act)
	return cha
}

// PipeFileFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeFileFork(inp <-chan *zip.File) (out1, out2 <-chan *zip.File) {
	cha1 := make(chan *zip.File)
	cha2 := make(chan *zip.File)
	go func(out1, out2 chan<- *zip.File, inp <-chan *zip.File) {
		defer close(out1)
		defer close(out2)
		for i := range inp {
			out1 <- i
			out2 <- i
		}
	}(cha1, cha2, inp)
	return cha1, cha2
}

// FileTube is the signature for a pipe function.
type FileTube func(inp <-chan *zip.File, out <-chan *zip.File)

// FileDaisy returns a channel to receive all inp after having passed thru tube.
func FileDaisy(inp <-chan *zip.File, tube FileTube) (out <-chan *zip.File) {
	cha := make(chan *zip.File)
	go tube(inp, cha)
	return cha
}

// FileDaisyChain returns a channel to receive all inp after having passed thru all tubes.
func FileDaisyChain(inp <-chan *zip.File, tubes ...FileTube) (out <-chan *zip.File) {
	cha := inp
	for i := range tubes {
		cha = FileDaisy(cha, tubes[i])
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
