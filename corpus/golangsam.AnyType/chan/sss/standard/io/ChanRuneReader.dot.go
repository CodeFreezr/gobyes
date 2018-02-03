// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// MakeRuneReaderChan returns a new open channel
// (simply a 'chan io.RuneReader' that is).
//
// Note: No 'RuneReader-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myRuneReaderPipelineStartsHere := MakeRuneReaderChan()
//	// ... lot's of code to design and build Your favourite "myRuneReaderWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myRuneReaderPipelineStartsHere <- drop
//	}
//	close(myRuneReaderPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeRuneReaderBuffer) the channel is unbuffered.
//
func MakeRuneReaderChan() (out chan io.RuneReader) {
	return make(chan io.RuneReader)
}

// ChanRuneReader returns a channel to receive all inputs before close.
func ChanRuneReader(inp ...io.RuneReader) (out <-chan io.RuneReader) {
	cha := make(chan io.RuneReader)
	go func(out chan<- io.RuneReader, inp ...io.RuneReader) {
		defer close(out)
		for i := range inp {
			out <- inp[i]
		}
	}(cha, inp...)
	return cha
}

// ChanRuneReaderSlice returns a channel to receive all inputs before close.
func ChanRuneReaderSlice(inp ...[]io.RuneReader) (out <-chan io.RuneReader) {
	cha := make(chan io.RuneReader)
	go func(out chan<- io.RuneReader, inp ...[]io.RuneReader) {
		defer close(out)
		for i := range inp {
			for j := range inp[i] {
				out <- inp[i][j]
			}
		}
	}(cha, inp...)
	return cha
}

// ChanRuneReaderFuncNok returns a channel to receive all results of act until nok before close.
func ChanRuneReaderFuncNok(act func() (io.RuneReader, bool)) (out <-chan io.RuneReader) {
	cha := make(chan io.RuneReader)
	go func(out chan<- io.RuneReader, act func() (io.RuneReader, bool)) {
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

// ChanRuneReaderFuncErr returns a channel to receive all results of act until err != nil before close.
func ChanRuneReaderFuncErr(act func() (io.RuneReader, error)) (out <-chan io.RuneReader) {
	cha := make(chan io.RuneReader)
	go func(out chan<- io.RuneReader, act func() (io.RuneReader, error)) {
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

// JoinRuneReader sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinRuneReader(out chan<- io.RuneReader, inp ...io.RuneReader) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- io.RuneReader, inp ...io.RuneReader) {
		defer close(done)
		for i := range inp {
			out <- inp[i]
		}
		done <- struct{}{}
	}(cha, out, inp...)
	return cha
}

// JoinRuneReaderSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinRuneReaderSlice(out chan<- io.RuneReader, inp ...[]io.RuneReader) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- io.RuneReader, inp ...[]io.RuneReader) {
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

// JoinRuneReaderChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinRuneReaderChan(out chan<- io.RuneReader, inp <-chan io.RuneReader) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- io.RuneReader, inp <-chan io.RuneReader) {
		defer close(done)
		for i := range inp {
			out <- i
		}
		done <- struct{}{}
	}(cha, out, inp)
	return cha
}

// DoneRuneReader returns a channel to receive one signal before close after inp has been drained.
func DoneRuneReader(inp <-chan io.RuneReader) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, inp <-chan io.RuneReader) {
		defer close(done)
		for i := range inp {
			_ = i // Drain inp
		}
		done <- struct{}{}
	}(cha, inp)
	return cha
}

// DoneRuneReaderSlice returns a channel which will receive a slice
// of all the RuneReaders received on inp channel before close.
// Unlike DoneRuneReader, a full slice is sent once, not just an event.
func DoneRuneReaderSlice(inp <-chan io.RuneReader) (done <-chan []io.RuneReader) {
	cha := make(chan []io.RuneReader)
	go func(inp <-chan io.RuneReader, done chan<- []io.RuneReader) {
		defer close(done)
		RuneReaderS := []io.RuneReader{}
		for i := range inp {
			RuneReaderS = append(RuneReaderS, i)
		}
		done <- RuneReaderS
	}(inp, cha)
	return cha
}

// DoneRuneReaderFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneRuneReaderFunc(inp <-chan io.RuneReader, act func(a io.RuneReader)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a io.RuneReader) { return }
	}
	go func(done chan<- struct{}, inp <-chan io.RuneReader, act func(a io.RuneReader)) {
		defer close(done)
		for i := range inp {
			act(i) // Apply action
		}
		done <- struct{}{}
	}(cha, inp, act)
	return cha
}

// PipeRuneReaderBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeRuneReaderBuffer(inp <-chan io.RuneReader, cap int) (out <-chan io.RuneReader) {
	cha := make(chan io.RuneReader, cap)
	go func(out chan<- io.RuneReader, inp <-chan io.RuneReader) {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}(cha, inp)
	return cha
}

// PipeRuneReaderFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeRuneReaderMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeRuneReaderFunc(inp <-chan io.RuneReader, act func(a io.RuneReader) io.RuneReader) (out <-chan io.RuneReader) {
	cha := make(chan io.RuneReader)
	if act == nil {
		act = func(a io.RuneReader) io.RuneReader { return a }
	}
	go func(out chan<- io.RuneReader, inp <-chan io.RuneReader, act func(a io.RuneReader) io.RuneReader) {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}(cha, inp, act)
	return cha
}

// PipeRuneReaderFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeRuneReaderFork(inp <-chan io.RuneReader) (out1, out2 <-chan io.RuneReader) {
	cha1 := make(chan io.RuneReader)
	cha2 := make(chan io.RuneReader)
	go func(out1, out2 chan<- io.RuneReader, inp <-chan io.RuneReader) {
		defer close(out1)
		defer close(out2)
		for i := range inp {
			out1 <- i
			out2 <- i
		}
	}(cha1, cha2, inp)
	return cha1, cha2
}

// RuneReaderTube is the signature for a pipe function.
type RuneReaderTube func(inp <-chan io.RuneReader, out <-chan io.RuneReader)

// RuneReaderDaisy returns a channel to receive all inp after having passed thru tube.
func RuneReaderDaisy(inp <-chan io.RuneReader, tube RuneReaderTube) (out <-chan io.RuneReader) {
	cha := make(chan io.RuneReader)
	go tube(inp, cha)
	return cha
}

// RuneReaderDaisyChain returns a channel to receive all inp after having passed thru all tubes.
func RuneReaderDaisyChain(inp <-chan io.RuneReader, tubes ...RuneReaderTube) (out <-chan io.RuneReader) {
	cha := inp
	for i := range tubes {
		cha = RuneReaderDaisy(cha, tubes[i])
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
