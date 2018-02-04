// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// MakeLimitedReaderChan returns a new open channel
// (simply a 'chan *io.LimitedReader' that is).
//
// Note: No 'LimitedReader-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myLimitedReaderPipelineStartsHere := MakeLimitedReaderChan()
//	// ... lot's of code to design and build Your favourite "myLimitedReaderWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myLimitedReaderPipelineStartsHere <- drop
//	}
//	close(myLimitedReaderPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeLimitedReaderBuffer) the channel is unbuffered.
//
func MakeLimitedReaderChan() chan *io.LimitedReader {
	return make(chan *io.LimitedReader)
}

// ChanLimitedReader returns a channel to receive all inputs before close.
func ChanLimitedReader(inp ...*io.LimitedReader) chan *io.LimitedReader {
	out := make(chan *io.LimitedReader)
	go func() {
		defer close(out)
		for i := range inp {
			out <- inp[i]
		}
	}()
	return out
}

// ChanLimitedReaderSlice returns a channel to receive all inputs before close.
func ChanLimitedReaderSlice(inp ...[]*io.LimitedReader) chan *io.LimitedReader {
	out := make(chan *io.LimitedReader)
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

// ChanLimitedReaderFuncNok returns a channel to receive all results of act until nok before close.
func ChanLimitedReaderFuncNok(act func() (*io.LimitedReader, bool)) <-chan *io.LimitedReader {
	out := make(chan *io.LimitedReader)
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

// ChanLimitedReaderFuncErr returns a channel to receive all results of act until err != nil before close.
func ChanLimitedReaderFuncErr(act func() (*io.LimitedReader, error)) <-chan *io.LimitedReader {
	out := make(chan *io.LimitedReader)
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

// JoinLimitedReader sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinLimitedReader(out chan<- *io.LimitedReader, inp ...*io.LimitedReader) chan struct{} {
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

// JoinLimitedReaderSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinLimitedReaderSlice(out chan<- *io.LimitedReader, inp ...[]*io.LimitedReader) chan struct{} {
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

// JoinLimitedReaderChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinLimitedReaderChan(out chan<- *io.LimitedReader, inp <-chan *io.LimitedReader) chan struct{} {
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

// DoneLimitedReader returns a channel to receive one signal before close after inp has been drained.
func DoneLimitedReader(inp <-chan *io.LimitedReader) chan struct{} {
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

// DoneLimitedReaderSlice returns a channel which will receive a slice
// of all the LimitedReaders received on inp channel before close.
// Unlike DoneLimitedReader, a full slice is sent once, not just an event.
func DoneLimitedReaderSlice(inp <-chan *io.LimitedReader) chan []*io.LimitedReader {
	done := make(chan []*io.LimitedReader)
	go func() {
		defer close(done)
		LimitedReaderS := []*io.LimitedReader{}
		for i := range inp {
			LimitedReaderS = append(LimitedReaderS, i)
		}
		done <- LimitedReaderS
	}()
	return done
}

// DoneLimitedReaderFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneLimitedReaderFunc(inp <-chan *io.LimitedReader, act func(a *io.LimitedReader)) chan struct{} {
	done := make(chan struct{})
	if act == nil {
		act = func(a *io.LimitedReader) { return }
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

// PipeLimitedReaderBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeLimitedReaderBuffer(inp <-chan *io.LimitedReader, cap int) chan *io.LimitedReader {
	out := make(chan *io.LimitedReader, cap)
	go func() {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}()
	return out
}

// PipeLimitedReaderFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeLimitedReaderMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeLimitedReaderFunc(inp <-chan *io.LimitedReader, act func(a *io.LimitedReader) *io.LimitedReader) chan *io.LimitedReader {
	out := make(chan *io.LimitedReader)
	if act == nil {
		act = func(a *io.LimitedReader) *io.LimitedReader { return a }
	}
	go func() {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}()
	return out
}

// PipeLimitedReaderFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeLimitedReaderFork(inp <-chan *io.LimitedReader) (chan *io.LimitedReader, chan *io.LimitedReader) {
	out1 := make(chan *io.LimitedReader)
	out2 := make(chan *io.LimitedReader)
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

// LimitedReaderTube is the signature for a pipe function.
type LimitedReaderTube func(inp <-chan *io.LimitedReader, out <-chan *io.LimitedReader)

// LimitedReaderDaisy returns a channel to receive all inp after having passed thru tube.
func LimitedReaderDaisy(inp <-chan *io.LimitedReader, tube LimitedReaderTube) (out <-chan *io.LimitedReader) {
	cha := make(chan *io.LimitedReader)
	go tube(inp, cha)
	return cha
}

// LimitedReaderDaisyChain returns a channel to receive all inp after having passed thru all tubes.
func LimitedReaderDaisyChain(inp <-chan *io.LimitedReader, tubes ...LimitedReaderTube) (out <-chan *io.LimitedReader) {
	cha := inp
	for i := range tubes {
		cha = LimitedReaderDaisy(cha, tubes[i])
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
