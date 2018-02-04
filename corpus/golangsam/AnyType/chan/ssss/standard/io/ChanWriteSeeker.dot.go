// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// MakeWriteSeekerChan returns a new open channel
// (simply a 'chan io.WriteSeeker' that is).
//
// Note: No 'WriteSeeker-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myWriteSeekerPipelineStartsHere := MakeWriteSeekerChan()
//	// ... lot's of code to design and build Your favourite "myWriteSeekerWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myWriteSeekerPipelineStartsHere <- drop
//	}
//	close(myWriteSeekerPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeWriteSeekerBuffer) the channel is unbuffered.
//
func MakeWriteSeekerChan() chan io.WriteSeeker {
	return make(chan io.WriteSeeker)
}

// ChanWriteSeeker returns a channel to receive all inputs before close.
func ChanWriteSeeker(inp ...io.WriteSeeker) chan io.WriteSeeker {
	out := make(chan io.WriteSeeker)
	go func() {
		defer close(out)
		for i := range inp {
			out <- inp[i]
		}
	}()
	return out
}

// ChanWriteSeekerSlice returns a channel to receive all inputs before close.
func ChanWriteSeekerSlice(inp ...[]io.WriteSeeker) chan io.WriteSeeker {
	out := make(chan io.WriteSeeker)
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

// ChanWriteSeekerFuncNok returns a channel to receive all results of act until nok before close.
func ChanWriteSeekerFuncNok(act func() (io.WriteSeeker, bool)) <-chan io.WriteSeeker {
	out := make(chan io.WriteSeeker)
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

// ChanWriteSeekerFuncErr returns a channel to receive all results of act until err != nil before close.
func ChanWriteSeekerFuncErr(act func() (io.WriteSeeker, error)) <-chan io.WriteSeeker {
	out := make(chan io.WriteSeeker)
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

// JoinWriteSeeker sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinWriteSeeker(out chan<- io.WriteSeeker, inp ...io.WriteSeeker) chan struct{} {
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

// JoinWriteSeekerSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinWriteSeekerSlice(out chan<- io.WriteSeeker, inp ...[]io.WriteSeeker) chan struct{} {
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

// JoinWriteSeekerChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinWriteSeekerChan(out chan<- io.WriteSeeker, inp <-chan io.WriteSeeker) chan struct{} {
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

// DoneWriteSeeker returns a channel to receive one signal before close after inp has been drained.
func DoneWriteSeeker(inp <-chan io.WriteSeeker) chan struct{} {
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

// DoneWriteSeekerSlice returns a channel which will receive a slice
// of all the WriteSeekers received on inp channel before close.
// Unlike DoneWriteSeeker, a full slice is sent once, not just an event.
func DoneWriteSeekerSlice(inp <-chan io.WriteSeeker) chan []io.WriteSeeker {
	done := make(chan []io.WriteSeeker)
	go func() {
		defer close(done)
		WriteSeekerS := []io.WriteSeeker{}
		for i := range inp {
			WriteSeekerS = append(WriteSeekerS, i)
		}
		done <- WriteSeekerS
	}()
	return done
}

// DoneWriteSeekerFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneWriteSeekerFunc(inp <-chan io.WriteSeeker, act func(a io.WriteSeeker)) chan struct{} {
	done := make(chan struct{})
	if act == nil {
		act = func(a io.WriteSeeker) { return }
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

// PipeWriteSeekerBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeWriteSeekerBuffer(inp <-chan io.WriteSeeker, cap int) chan io.WriteSeeker {
	out := make(chan io.WriteSeeker, cap)
	go func() {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}()
	return out
}

// PipeWriteSeekerFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeWriteSeekerMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeWriteSeekerFunc(inp <-chan io.WriteSeeker, act func(a io.WriteSeeker) io.WriteSeeker) chan io.WriteSeeker {
	out := make(chan io.WriteSeeker)
	if act == nil {
		act = func(a io.WriteSeeker) io.WriteSeeker { return a }
	}
	go func() {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}()
	return out
}

// PipeWriteSeekerFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeWriteSeekerFork(inp <-chan io.WriteSeeker) (chan io.WriteSeeker, chan io.WriteSeeker) {
	out1 := make(chan io.WriteSeeker)
	out2 := make(chan io.WriteSeeker)
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

// WriteSeekerTube is the signature for a pipe function.
type WriteSeekerTube func(inp <-chan io.WriteSeeker, out <-chan io.WriteSeeker)

// WriteSeekerDaisy returns a channel to receive all inp after having passed thru tube.
func WriteSeekerDaisy(inp <-chan io.WriteSeeker, tube WriteSeekerTube) (out <-chan io.WriteSeeker) {
	cha := make(chan io.WriteSeeker)
	go tube(inp, cha)
	return cha
}

// WriteSeekerDaisyChain returns a channel to receive all inp after having passed thru all tubes.
func WriteSeekerDaisyChain(inp <-chan io.WriteSeeker, tubes ...WriteSeekerTube) (out <-chan io.WriteSeeker) {
	cha := inp
	for i := range tubes {
		cha = WriteSeekerDaisy(cha, tubes[i])
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
