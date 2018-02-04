// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// MakePipeWriterChan returns a new open channel
// (simply a 'chan *io.PipeWriter' that is).
//
// Note: No 'PipeWriter-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myPipeWriterPipelineStartsHere := MakePipeWriterChan()
//	// ... lot's of code to design and build Your favourite "myPipeWriterWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myPipeWriterPipelineStartsHere <- drop
//	}
//	close(myPipeWriterPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipePipeWriterBuffer) the channel is unbuffered.
//
func MakePipeWriterChan() chan *io.PipeWriter {
	return make(chan *io.PipeWriter)
}

// ChanPipeWriter returns a channel to receive all inputs before close.
func ChanPipeWriter(inp ...*io.PipeWriter) chan *io.PipeWriter {
	out := make(chan *io.PipeWriter)
	go func() {
		defer close(out)
		for i := range inp {
			out <- inp[i]
		}
	}()
	return out
}

// ChanPipeWriterSlice returns a channel to receive all inputs before close.
func ChanPipeWriterSlice(inp ...[]*io.PipeWriter) chan *io.PipeWriter {
	out := make(chan *io.PipeWriter)
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

// ChanPipeWriterFuncNok returns a channel to receive all results of act until nok before close.
func ChanPipeWriterFuncNok(act func() (*io.PipeWriter, bool)) <-chan *io.PipeWriter {
	out := make(chan *io.PipeWriter)
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

// ChanPipeWriterFuncErr returns a channel to receive all results of act until err != nil before close.
func ChanPipeWriterFuncErr(act func() (*io.PipeWriter, error)) <-chan *io.PipeWriter {
	out := make(chan *io.PipeWriter)
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

// JoinPipeWriter sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinPipeWriter(out chan<- *io.PipeWriter, inp ...*io.PipeWriter) chan struct{} {
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

// JoinPipeWriterSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinPipeWriterSlice(out chan<- *io.PipeWriter, inp ...[]*io.PipeWriter) chan struct{} {
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

// JoinPipeWriterChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinPipeWriterChan(out chan<- *io.PipeWriter, inp <-chan *io.PipeWriter) chan struct{} {
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

// DonePipeWriter returns a channel to receive one signal before close after inp has been drained.
func DonePipeWriter(inp <-chan *io.PipeWriter) chan struct{} {
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

// DonePipeWriterSlice returns a channel which will receive a slice
// of all the PipeWriters received on inp channel before close.
// Unlike DonePipeWriter, a full slice is sent once, not just an event.
func DonePipeWriterSlice(inp <-chan *io.PipeWriter) chan []*io.PipeWriter {
	done := make(chan []*io.PipeWriter)
	go func() {
		defer close(done)
		PipeWriterS := []*io.PipeWriter{}
		for i := range inp {
			PipeWriterS = append(PipeWriterS, i)
		}
		done <- PipeWriterS
	}()
	return done
}

// DonePipeWriterFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DonePipeWriterFunc(inp <-chan *io.PipeWriter, act func(a *io.PipeWriter)) chan struct{} {
	done := make(chan struct{})
	if act == nil {
		act = func(a *io.PipeWriter) { return }
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

// PipePipeWriterBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipePipeWriterBuffer(inp <-chan *io.PipeWriter, cap int) chan *io.PipeWriter {
	out := make(chan *io.PipeWriter, cap)
	go func() {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}()
	return out
}

// PipePipeWriterFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipePipeWriterMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipePipeWriterFunc(inp <-chan *io.PipeWriter, act func(a *io.PipeWriter) *io.PipeWriter) chan *io.PipeWriter {
	out := make(chan *io.PipeWriter)
	if act == nil {
		act = func(a *io.PipeWriter) *io.PipeWriter { return a }
	}
	go func() {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}()
	return out
}

// PipePipeWriterFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipePipeWriterFork(inp <-chan *io.PipeWriter) (chan *io.PipeWriter, chan *io.PipeWriter) {
	out1 := make(chan *io.PipeWriter)
	out2 := make(chan *io.PipeWriter)
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

// PipeWriterTube is the signature for a pipe function.
type PipeWriterTube func(inp <-chan *io.PipeWriter, out <-chan *io.PipeWriter)

// PipeWriterDaisy returns a channel to receive all inp after having passed thru tube.
func PipeWriterDaisy(inp <-chan *io.PipeWriter, tube PipeWriterTube) (out <-chan *io.PipeWriter) {
	cha := make(chan *io.PipeWriter)
	go tube(inp, cha)
	return cha
}

// PipeWriterDaisyChain returns a channel to receive all inp after having passed thru all tubes.
func PipeWriterDaisyChain(inp <-chan *io.PipeWriter, tubes ...PipeWriterTube) (out <-chan *io.PipeWriter) {
	cha := inp
	for i := range tubes {
		cha = PipeWriterDaisy(cha, tubes[i])
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
