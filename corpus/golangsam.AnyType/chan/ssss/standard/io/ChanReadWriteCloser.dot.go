// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// MakeReadWriteCloserChan returns a new open channel
// (simply a 'chan io.ReadWriteCloser' that is).
//
// Note: No 'ReadWriteCloser-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myReadWriteCloserPipelineStartsHere := MakeReadWriteCloserChan()
//	// ... lot's of code to design and build Your favourite "myReadWriteCloserWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myReadWriteCloserPipelineStartsHere <- drop
//	}
//	close(myReadWriteCloserPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeReadWriteCloserBuffer) the channel is unbuffered.
//
func MakeReadWriteCloserChan() chan io.ReadWriteCloser {
	return make(chan io.ReadWriteCloser)
}

// ChanReadWriteCloser returns a channel to receive all inputs before close.
func ChanReadWriteCloser(inp ...io.ReadWriteCloser) chan io.ReadWriteCloser {
	out := make(chan io.ReadWriteCloser)
	go func() {
		defer close(out)
		for i := range inp {
			out <- inp[i]
		}
	}()
	return out
}

// ChanReadWriteCloserSlice returns a channel to receive all inputs before close.
func ChanReadWriteCloserSlice(inp ...[]io.ReadWriteCloser) chan io.ReadWriteCloser {
	out := make(chan io.ReadWriteCloser)
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

// ChanReadWriteCloserFuncNok returns a channel to receive all results of act until nok before close.
func ChanReadWriteCloserFuncNok(act func() (io.ReadWriteCloser, bool)) <-chan io.ReadWriteCloser {
	out := make(chan io.ReadWriteCloser)
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

// ChanReadWriteCloserFuncErr returns a channel to receive all results of act until err != nil before close.
func ChanReadWriteCloserFuncErr(act func() (io.ReadWriteCloser, error)) <-chan io.ReadWriteCloser {
	out := make(chan io.ReadWriteCloser)
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

// JoinReadWriteCloser sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinReadWriteCloser(out chan<- io.ReadWriteCloser, inp ...io.ReadWriteCloser) chan struct{} {
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

// JoinReadWriteCloserSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinReadWriteCloserSlice(out chan<- io.ReadWriteCloser, inp ...[]io.ReadWriteCloser) chan struct{} {
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

// JoinReadWriteCloserChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinReadWriteCloserChan(out chan<- io.ReadWriteCloser, inp <-chan io.ReadWriteCloser) chan struct{} {
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

// DoneReadWriteCloser returns a channel to receive one signal before close after inp has been drained.
func DoneReadWriteCloser(inp <-chan io.ReadWriteCloser) chan struct{} {
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

// DoneReadWriteCloserSlice returns a channel which will receive a slice
// of all the ReadWriteClosers received on inp channel before close.
// Unlike DoneReadWriteCloser, a full slice is sent once, not just an event.
func DoneReadWriteCloserSlice(inp <-chan io.ReadWriteCloser) chan []io.ReadWriteCloser {
	done := make(chan []io.ReadWriteCloser)
	go func() {
		defer close(done)
		ReadWriteCloserS := []io.ReadWriteCloser{}
		for i := range inp {
			ReadWriteCloserS = append(ReadWriteCloserS, i)
		}
		done <- ReadWriteCloserS
	}()
	return done
}

// DoneReadWriteCloserFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneReadWriteCloserFunc(inp <-chan io.ReadWriteCloser, act func(a io.ReadWriteCloser)) chan struct{} {
	done := make(chan struct{})
	if act == nil {
		act = func(a io.ReadWriteCloser) { return }
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

// PipeReadWriteCloserBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeReadWriteCloserBuffer(inp <-chan io.ReadWriteCloser, cap int) chan io.ReadWriteCloser {
	out := make(chan io.ReadWriteCloser, cap)
	go func() {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}()
	return out
}

// PipeReadWriteCloserFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeReadWriteCloserMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeReadWriteCloserFunc(inp <-chan io.ReadWriteCloser, act func(a io.ReadWriteCloser) io.ReadWriteCloser) chan io.ReadWriteCloser {
	out := make(chan io.ReadWriteCloser)
	if act == nil {
		act = func(a io.ReadWriteCloser) io.ReadWriteCloser { return a }
	}
	go func() {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}()
	return out
}

// PipeReadWriteCloserFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeReadWriteCloserFork(inp <-chan io.ReadWriteCloser) (chan io.ReadWriteCloser, chan io.ReadWriteCloser) {
	out1 := make(chan io.ReadWriteCloser)
	out2 := make(chan io.ReadWriteCloser)
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

// ReadWriteCloserTube is the signature for a pipe function.
type ReadWriteCloserTube func(inp <-chan io.ReadWriteCloser, out <-chan io.ReadWriteCloser)

// ReadWriteCloserDaisy returns a channel to receive all inp after having passed thru tube.
func ReadWriteCloserDaisy(inp <-chan io.ReadWriteCloser, tube ReadWriteCloserTube) (out <-chan io.ReadWriteCloser) {
	cha := make(chan io.ReadWriteCloser)
	go tube(inp, cha)
	return cha
}

// ReadWriteCloserDaisyChain returns a channel to receive all inp after having passed thru all tubes.
func ReadWriteCloserDaisyChain(inp <-chan io.ReadWriteCloser, tubes ...ReadWriteCloserTube) (out <-chan io.ReadWriteCloser) {
	cha := inp
	for i := range tubes {
		cha = ReadWriteCloserDaisy(cha, tubes[i])
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
