// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// MakeWriteCloserChan returns a new open channel
// (simply a 'chan io.WriteCloser' that is).
//
// Note: No 'WriteCloser-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myWriteCloserPipelineStartsHere := MakeWriteCloserChan()
//	// ... lot's of code to design and build Your favourite "myWriteCloserWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myWriteCloserPipelineStartsHere <- drop
//	}
//	close(myWriteCloserPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeWriteCloserBuffer) the channel is unbuffered.
//
func MakeWriteCloserChan() chan io.WriteCloser {
	return make(chan io.WriteCloser)
}

// ChanWriteCloser returns a channel to receive all inputs before close.
func ChanWriteCloser(inp ...io.WriteCloser) chan io.WriteCloser {
	out := make(chan io.WriteCloser)
	go func() {
		defer close(out)
		for i := range inp {
			out <- inp[i]
		}
	}()
	return out
}

// ChanWriteCloserSlice returns a channel to receive all inputs before close.
func ChanWriteCloserSlice(inp ...[]io.WriteCloser) chan io.WriteCloser {
	out := make(chan io.WriteCloser)
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

// ChanWriteCloserFuncNok returns a channel to receive all results of act until nok before close.
func ChanWriteCloserFuncNok(act func() (io.WriteCloser, bool)) <-chan io.WriteCloser {
	out := make(chan io.WriteCloser)
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

// ChanWriteCloserFuncErr returns a channel to receive all results of act until err != nil before close.
func ChanWriteCloserFuncErr(act func() (io.WriteCloser, error)) <-chan io.WriteCloser {
	out := make(chan io.WriteCloser)
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

// JoinWriteCloser sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinWriteCloser(out chan<- io.WriteCloser, inp ...io.WriteCloser) chan struct{} {
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

// JoinWriteCloserSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinWriteCloserSlice(out chan<- io.WriteCloser, inp ...[]io.WriteCloser) chan struct{} {
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

// JoinWriteCloserChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinWriteCloserChan(out chan<- io.WriteCloser, inp <-chan io.WriteCloser) chan struct{} {
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

// DoneWriteCloser returns a channel to receive one signal before close after inp has been drained.
func DoneWriteCloser(inp <-chan io.WriteCloser) chan struct{} {
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

// DoneWriteCloserSlice returns a channel which will receive a slice
// of all the WriteClosers received on inp channel before close.
// Unlike DoneWriteCloser, a full slice is sent once, not just an event.
func DoneWriteCloserSlice(inp <-chan io.WriteCloser) chan []io.WriteCloser {
	done := make(chan []io.WriteCloser)
	go func() {
		defer close(done)
		WriteCloserS := []io.WriteCloser{}
		for i := range inp {
			WriteCloserS = append(WriteCloserS, i)
		}
		done <- WriteCloserS
	}()
	return done
}

// DoneWriteCloserFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneWriteCloserFunc(inp <-chan io.WriteCloser, act func(a io.WriteCloser)) chan struct{} {
	done := make(chan struct{})
	if act == nil {
		act = func(a io.WriteCloser) { return }
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

// PipeWriteCloserBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeWriteCloserBuffer(inp <-chan io.WriteCloser, cap int) chan io.WriteCloser {
	out := make(chan io.WriteCloser, cap)
	go func() {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}()
	return out
}

// PipeWriteCloserFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeWriteCloserMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeWriteCloserFunc(inp <-chan io.WriteCloser, act func(a io.WriteCloser) io.WriteCloser) chan io.WriteCloser {
	out := make(chan io.WriteCloser)
	if act == nil {
		act = func(a io.WriteCloser) io.WriteCloser { return a }
	}
	go func() {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}()
	return out
}

// PipeWriteCloserFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeWriteCloserFork(inp <-chan io.WriteCloser) (chan io.WriteCloser, chan io.WriteCloser) {
	out1 := make(chan io.WriteCloser)
	out2 := make(chan io.WriteCloser)
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

// WriteCloserTube is the signature for a pipe function.
type WriteCloserTube func(inp <-chan io.WriteCloser, out <-chan io.WriteCloser)

// WriteCloserDaisy returns a channel to receive all inp after having passed thru tube.
func WriteCloserDaisy(inp <-chan io.WriteCloser, tube WriteCloserTube) (out <-chan io.WriteCloser) {
	cha := make(chan io.WriteCloser)
	go tube(inp, cha)
	return cha
}

// WriteCloserDaisyChain returns a channel to receive all inp after having passed thru all tubes.
func WriteCloserDaisyChain(inp <-chan io.WriteCloser, tubes ...WriteCloserTube) (out <-chan io.WriteCloser) {
	cha := inp
	for i := range tubes {
		cha = WriteCloserDaisy(cha, tubes[i])
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
