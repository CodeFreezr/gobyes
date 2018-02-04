// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// MakeWriterToChan returns a new open channel
// (simply a 'chan io.WriterTo' that is).
//
// Note: No 'WriterTo-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myWriterToPipelineStartsHere := MakeWriterToChan()
//	// ... lot's of code to design and build Your favourite "myWriterToWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myWriterToPipelineStartsHere <- drop
//	}
//	close(myWriterToPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeWriterToBuffer) the channel is unbuffered.
//
func MakeWriterToChan() chan io.WriterTo {
	return make(chan io.WriterTo)
}

// ChanWriterTo returns a channel to receive all inputs before close.
func ChanWriterTo(inp ...io.WriterTo) chan io.WriterTo {
	out := make(chan io.WriterTo)
	go func() {
		defer close(out)
		for i := range inp {
			out <- inp[i]
		}
	}()
	return out
}

// ChanWriterToSlice returns a channel to receive all inputs before close.
func ChanWriterToSlice(inp ...[]io.WriterTo) chan io.WriterTo {
	out := make(chan io.WriterTo)
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

// ChanWriterToFuncNok returns a channel to receive all results of act until nok before close.
func ChanWriterToFuncNok(act func() (io.WriterTo, bool)) <-chan io.WriterTo {
	out := make(chan io.WriterTo)
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

// ChanWriterToFuncErr returns a channel to receive all results of act until err != nil before close.
func ChanWriterToFuncErr(act func() (io.WriterTo, error)) <-chan io.WriterTo {
	out := make(chan io.WriterTo)
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

// JoinWriterTo sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinWriterTo(out chan<- io.WriterTo, inp ...io.WriterTo) chan struct{} {
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

// JoinWriterToSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinWriterToSlice(out chan<- io.WriterTo, inp ...[]io.WriterTo) chan struct{} {
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

// JoinWriterToChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinWriterToChan(out chan<- io.WriterTo, inp <-chan io.WriterTo) chan struct{} {
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

// DoneWriterTo returns a channel to receive one signal before close after inp has been drained.
func DoneWriterTo(inp <-chan io.WriterTo) chan struct{} {
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

// DoneWriterToSlice returns a channel which will receive a slice
// of all the WriterTos received on inp channel before close.
// Unlike DoneWriterTo, a full slice is sent once, not just an event.
func DoneWriterToSlice(inp <-chan io.WriterTo) chan []io.WriterTo {
	done := make(chan []io.WriterTo)
	go func() {
		defer close(done)
		WriterToS := []io.WriterTo{}
		for i := range inp {
			WriterToS = append(WriterToS, i)
		}
		done <- WriterToS
	}()
	return done
}

// DoneWriterToFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneWriterToFunc(inp <-chan io.WriterTo, act func(a io.WriterTo)) chan struct{} {
	done := make(chan struct{})
	if act == nil {
		act = func(a io.WriterTo) { return }
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

// PipeWriterToBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeWriterToBuffer(inp <-chan io.WriterTo, cap int) chan io.WriterTo {
	out := make(chan io.WriterTo, cap)
	go func() {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}()
	return out
}

// PipeWriterToFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeWriterToMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeWriterToFunc(inp <-chan io.WriterTo, act func(a io.WriterTo) io.WriterTo) chan io.WriterTo {
	out := make(chan io.WriterTo)
	if act == nil {
		act = func(a io.WriterTo) io.WriterTo { return a }
	}
	go func() {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}()
	return out
}

// PipeWriterToFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeWriterToFork(inp <-chan io.WriterTo) (chan io.WriterTo, chan io.WriterTo) {
	out1 := make(chan io.WriterTo)
	out2 := make(chan io.WriterTo)
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

// WriterToTube is the signature for a pipe function.
type WriterToTube func(inp <-chan io.WriterTo, out <-chan io.WriterTo)

// WriterToDaisy returns a channel to receive all inp after having passed thru tube.
func WriterToDaisy(inp <-chan io.WriterTo, tube WriterToTube) (out <-chan io.WriterTo) {
	cha := make(chan io.WriterTo)
	go tube(inp, cha)
	return cha
}

// WriterToDaisyChain returns a channel to receive all inp after having passed thru all tubes.
func WriterToDaisyChain(inp <-chan io.WriterTo, tubes ...WriterToTube) (out <-chan io.WriterTo) {
	cha := inp
	for i := range tubes {
		cha = WriterToDaisy(cha, tubes[i])
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
