// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// MakeReaderAtChan returns a new open channel
// (simply a 'chan io.ReaderAt' that is).
//
// Note: No 'ReaderAt-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myReaderAtPipelineStartsHere := MakeReaderAtChan()
//	// ... lot's of code to design and build Your favourite "myReaderAtWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myReaderAtPipelineStartsHere <- drop
//	}
//	close(myReaderAtPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeReaderAtBuffer) the channel is unbuffered.
//
func MakeReaderAtChan() chan io.ReaderAt {
	return make(chan io.ReaderAt)
}

// ChanReaderAt returns a channel to receive all inputs before close.
func ChanReaderAt(inp ...io.ReaderAt) chan io.ReaderAt {
	out := make(chan io.ReaderAt)
	go func() {
		defer close(out)
		for i := range inp {
			out <- inp[i]
		}
	}()
	return out
}

// ChanReaderAtSlice returns a channel to receive all inputs before close.
func ChanReaderAtSlice(inp ...[]io.ReaderAt) chan io.ReaderAt {
	out := make(chan io.ReaderAt)
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

// ChanReaderAtFuncNok returns a channel to receive all results of act until nok before close.
func ChanReaderAtFuncNok(act func() (io.ReaderAt, bool)) <-chan io.ReaderAt {
	out := make(chan io.ReaderAt)
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

// ChanReaderAtFuncErr returns a channel to receive all results of act until err != nil before close.
func ChanReaderAtFuncErr(act func() (io.ReaderAt, error)) <-chan io.ReaderAt {
	out := make(chan io.ReaderAt)
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

// JoinReaderAt sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinReaderAt(out chan<- io.ReaderAt, inp ...io.ReaderAt) chan struct{} {
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

// JoinReaderAtSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinReaderAtSlice(out chan<- io.ReaderAt, inp ...[]io.ReaderAt) chan struct{} {
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

// JoinReaderAtChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinReaderAtChan(out chan<- io.ReaderAt, inp <-chan io.ReaderAt) chan struct{} {
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

// DoneReaderAt returns a channel to receive one signal before close after inp has been drained.
func DoneReaderAt(inp <-chan io.ReaderAt) chan struct{} {
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

// DoneReaderAtSlice returns a channel which will receive a slice
// of all the ReaderAts received on inp channel before close.
// Unlike DoneReaderAt, a full slice is sent once, not just an event.
func DoneReaderAtSlice(inp <-chan io.ReaderAt) chan []io.ReaderAt {
	done := make(chan []io.ReaderAt)
	go func() {
		defer close(done)
		ReaderAtS := []io.ReaderAt{}
		for i := range inp {
			ReaderAtS = append(ReaderAtS, i)
		}
		done <- ReaderAtS
	}()
	return done
}

// DoneReaderAtFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneReaderAtFunc(inp <-chan io.ReaderAt, act func(a io.ReaderAt)) chan struct{} {
	done := make(chan struct{})
	if act == nil {
		act = func(a io.ReaderAt) { return }
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

// PipeReaderAtBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeReaderAtBuffer(inp <-chan io.ReaderAt, cap int) chan io.ReaderAt {
	out := make(chan io.ReaderAt, cap)
	go func() {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}()
	return out
}

// PipeReaderAtFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeReaderAtMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeReaderAtFunc(inp <-chan io.ReaderAt, act func(a io.ReaderAt) io.ReaderAt) chan io.ReaderAt {
	out := make(chan io.ReaderAt)
	if act == nil {
		act = func(a io.ReaderAt) io.ReaderAt { return a }
	}
	go func() {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}()
	return out
}

// PipeReaderAtFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeReaderAtFork(inp <-chan io.ReaderAt) (chan io.ReaderAt, chan io.ReaderAt) {
	out1 := make(chan io.ReaderAt)
	out2 := make(chan io.ReaderAt)
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

// ReaderAtTube is the signature for a pipe function.
type ReaderAtTube func(inp <-chan io.ReaderAt, out <-chan io.ReaderAt)

// ReaderAtDaisy returns a channel to receive all inp after having passed thru tube.
func ReaderAtDaisy(inp <-chan io.ReaderAt, tube ReaderAtTube) (out <-chan io.ReaderAt) {
	cha := make(chan io.ReaderAt)
	go tube(inp, cha)
	return cha
}

// ReaderAtDaisyChain returns a channel to receive all inp after having passed thru all tubes.
func ReaderAtDaisyChain(inp <-chan io.ReaderAt, tubes ...ReaderAtTube) (out <-chan io.ReaderAt) {
	cha := inp
	for i := range tubes {
		cha = ReaderAtDaisy(cha, tubes[i])
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
