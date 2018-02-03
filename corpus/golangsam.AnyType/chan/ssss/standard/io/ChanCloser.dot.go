// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// MakeCloserChan returns a new open channel
// (simply a 'chan io.Closer' that is).
//
// Note: No 'Closer-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myCloserPipelineStartsHere := MakeCloserChan()
//	// ... lot's of code to design and build Your favourite "myCloserWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myCloserPipelineStartsHere <- drop
//	}
//	close(myCloserPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeCloserBuffer) the channel is unbuffered.
//
func MakeCloserChan() chan io.Closer {
	return make(chan io.Closer)
}

// ChanCloser returns a channel to receive all inputs before close.
func ChanCloser(inp ...io.Closer) chan io.Closer {
	out := make(chan io.Closer)
	go func() {
		defer close(out)
		for i := range inp {
			out <- inp[i]
		}
	}()
	return out
}

// ChanCloserSlice returns a channel to receive all inputs before close.
func ChanCloserSlice(inp ...[]io.Closer) chan io.Closer {
	out := make(chan io.Closer)
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

// ChanCloserFuncNok returns a channel to receive all results of act until nok before close.
func ChanCloserFuncNok(act func() (io.Closer, bool)) <-chan io.Closer {
	out := make(chan io.Closer)
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

// ChanCloserFuncErr returns a channel to receive all results of act until err != nil before close.
func ChanCloserFuncErr(act func() (io.Closer, error)) <-chan io.Closer {
	out := make(chan io.Closer)
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

// JoinCloser sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinCloser(out chan<- io.Closer, inp ...io.Closer) chan struct{} {
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

// JoinCloserSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinCloserSlice(out chan<- io.Closer, inp ...[]io.Closer) chan struct{} {
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

// JoinCloserChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinCloserChan(out chan<- io.Closer, inp <-chan io.Closer) chan struct{} {
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

// DoneCloser returns a channel to receive one signal before close after inp has been drained.
func DoneCloser(inp <-chan io.Closer) chan struct{} {
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

// DoneCloserSlice returns a channel which will receive a slice
// of all the Closers received on inp channel before close.
// Unlike DoneCloser, a full slice is sent once, not just an event.
func DoneCloserSlice(inp <-chan io.Closer) chan []io.Closer {
	done := make(chan []io.Closer)
	go func() {
		defer close(done)
		CloserS := []io.Closer{}
		for i := range inp {
			CloserS = append(CloserS, i)
		}
		done <- CloserS
	}()
	return done
}

// DoneCloserFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneCloserFunc(inp <-chan io.Closer, act func(a io.Closer)) chan struct{} {
	done := make(chan struct{})
	if act == nil {
		act = func(a io.Closer) { return }
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

// PipeCloserBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeCloserBuffer(inp <-chan io.Closer, cap int) chan io.Closer {
	out := make(chan io.Closer, cap)
	go func() {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}()
	return out
}

// PipeCloserFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeCloserMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeCloserFunc(inp <-chan io.Closer, act func(a io.Closer) io.Closer) chan io.Closer {
	out := make(chan io.Closer)
	if act == nil {
		act = func(a io.Closer) io.Closer { return a }
	}
	go func() {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}()
	return out
}

// PipeCloserFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeCloserFork(inp <-chan io.Closer) (chan io.Closer, chan io.Closer) {
	out1 := make(chan io.Closer)
	out2 := make(chan io.Closer)
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

// CloserTube is the signature for a pipe function.
type CloserTube func(inp <-chan io.Closer, out <-chan io.Closer)

// CloserDaisy returns a channel to receive all inp after having passed thru tube.
func CloserDaisy(inp <-chan io.Closer, tube CloserTube) (out <-chan io.Closer) {
	cha := make(chan io.Closer)
	go tube(inp, cha)
	return cha
}

// CloserDaisyChain returns a channel to receive all inp after having passed thru all tubes.
func CloserDaisyChain(inp <-chan io.Closer, tubes ...CloserTube) (out <-chan io.Closer) {
	cha := inp
	for i := range tubes {
		cha = CloserDaisy(cha, tubes[i])
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
