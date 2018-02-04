// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// MakePipeReaderChan returns a new open channel
// (simply a 'chan *io.PipeReader' that is).
//
// Note: No 'PipeReader-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myPipeReaderPipelineStartsHere := MakePipeReaderChan()
//	// ... lot's of code to design and build Your favourite "myPipeReaderWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myPipeReaderPipelineStartsHere <- drop
//	}
//	close(myPipeReaderPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipePipeReaderBuffer) the channel is unbuffered.
//
func MakePipeReaderChan() chan *io.PipeReader {
	return make(chan *io.PipeReader)
}

// ChanPipeReader returns a channel to receive all inputs before close.
func ChanPipeReader(inp ...*io.PipeReader) chan *io.PipeReader {
	out := make(chan *io.PipeReader)
	go func() {
		defer close(out)
		for i := range inp {
			out <- inp[i]
		}
	}()
	return out
}

// ChanPipeReaderSlice returns a channel to receive all inputs before close.
func ChanPipeReaderSlice(inp ...[]*io.PipeReader) chan *io.PipeReader {
	out := make(chan *io.PipeReader)
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

// ChanPipeReaderFuncNok returns a channel to receive all results of act until nok before close.
func ChanPipeReaderFuncNok(act func() (*io.PipeReader, bool)) <-chan *io.PipeReader {
	out := make(chan *io.PipeReader)
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

// ChanPipeReaderFuncErr returns a channel to receive all results of act until err != nil before close.
func ChanPipeReaderFuncErr(act func() (*io.PipeReader, error)) <-chan *io.PipeReader {
	out := make(chan *io.PipeReader)
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

// JoinPipeReader sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinPipeReader(out chan<- *io.PipeReader, inp ...*io.PipeReader) chan struct{} {
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

// JoinPipeReaderSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinPipeReaderSlice(out chan<- *io.PipeReader, inp ...[]*io.PipeReader) chan struct{} {
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

// JoinPipeReaderChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinPipeReaderChan(out chan<- *io.PipeReader, inp <-chan *io.PipeReader) chan struct{} {
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

// DonePipeReader returns a channel to receive one signal before close after inp has been drained.
func DonePipeReader(inp <-chan *io.PipeReader) chan struct{} {
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

// DonePipeReaderSlice returns a channel which will receive a slice
// of all the PipeReaders received on inp channel before close.
// Unlike DonePipeReader, a full slice is sent once, not just an event.
func DonePipeReaderSlice(inp <-chan *io.PipeReader) chan []*io.PipeReader {
	done := make(chan []*io.PipeReader)
	go func() {
		defer close(done)
		PipeReaderS := []*io.PipeReader{}
		for i := range inp {
			PipeReaderS = append(PipeReaderS, i)
		}
		done <- PipeReaderS
	}()
	return done
}

// DonePipeReaderFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DonePipeReaderFunc(inp <-chan *io.PipeReader, act func(a *io.PipeReader)) chan struct{} {
	done := make(chan struct{})
	if act == nil {
		act = func(a *io.PipeReader) { return }
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

// PipePipeReaderBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipePipeReaderBuffer(inp <-chan *io.PipeReader, cap int) chan *io.PipeReader {
	out := make(chan *io.PipeReader, cap)
	go func() {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}()
	return out
}

// PipePipeReaderFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipePipeReaderMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipePipeReaderFunc(inp <-chan *io.PipeReader, act func(a *io.PipeReader) *io.PipeReader) chan *io.PipeReader {
	out := make(chan *io.PipeReader)
	if act == nil {
		act = func(a *io.PipeReader) *io.PipeReader { return a }
	}
	go func() {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}()
	return out
}

// PipePipeReaderFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipePipeReaderFork(inp <-chan *io.PipeReader) (chan *io.PipeReader, chan *io.PipeReader) {
	out1 := make(chan *io.PipeReader)
	out2 := make(chan *io.PipeReader)
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

// PipeReaderTube is the signature for a pipe function.
type PipeReaderTube func(inp <-chan *io.PipeReader, out <-chan *io.PipeReader)

// PipeReaderDaisy returns a channel to receive all inp after having passed thru tube.
func PipeReaderDaisy(inp <-chan *io.PipeReader, tube PipeReaderTube) (out <-chan *io.PipeReader) {
	cha := make(chan *io.PipeReader)
	go tube(inp, cha)
	return cha
}

// PipeReaderDaisyChain returns a channel to receive all inp after having passed thru all tubes.
func PipeReaderDaisyChain(inp <-chan *io.PipeReader, tubes ...PipeReaderTube) (out <-chan *io.PipeReader) {
	cha := inp
	for i := range tubes {
		cha = PipeReaderDaisy(cha, tubes[i])
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
