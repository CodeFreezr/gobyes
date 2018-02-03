// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// MakeReadSeekerChan returns a new open channel
// (simply a 'chan io.ReadSeeker' that is).
//
// Note: No 'ReadSeeker-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myReadSeekerPipelineStartsHere := MakeReadSeekerChan()
//	// ... lot's of code to design and build Your favourite "myReadSeekerWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myReadSeekerPipelineStartsHere <- drop
//	}
//	close(myReadSeekerPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeReadSeekerBuffer) the channel is unbuffered.
//
func MakeReadSeekerChan() (out chan io.ReadSeeker) {
	return make(chan io.ReadSeeker)
}

// ChanReadSeeker returns a channel to receive all inputs before close.
func ChanReadSeeker(inp ...io.ReadSeeker) (out <-chan io.ReadSeeker) {
	cha := make(chan io.ReadSeeker)
	go func(out chan<- io.ReadSeeker, inp ...io.ReadSeeker) {
		defer close(out)
		for i := range inp {
			out <- inp[i]
		}
	}(cha, inp...)
	return cha
}

// ChanReadSeekerSlice returns a channel to receive all inputs before close.
func ChanReadSeekerSlice(inp ...[]io.ReadSeeker) (out <-chan io.ReadSeeker) {
	cha := make(chan io.ReadSeeker)
	go func(out chan<- io.ReadSeeker, inp ...[]io.ReadSeeker) {
		defer close(out)
		for i := range inp {
			for j := range inp[i] {
				out <- inp[i][j]
			}
		}
	}(cha, inp...)
	return cha
}

// ChanReadSeekerFuncNok returns a channel to receive all results of act until nok before close.
func ChanReadSeekerFuncNok(act func() (io.ReadSeeker, bool)) (out <-chan io.ReadSeeker) {
	cha := make(chan io.ReadSeeker)
	go func(out chan<- io.ReadSeeker, act func() (io.ReadSeeker, bool)) {
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

// ChanReadSeekerFuncErr returns a channel to receive all results of act until err != nil before close.
func ChanReadSeekerFuncErr(act func() (io.ReadSeeker, error)) (out <-chan io.ReadSeeker) {
	cha := make(chan io.ReadSeeker)
	go func(out chan<- io.ReadSeeker, act func() (io.ReadSeeker, error)) {
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

// JoinReadSeeker sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinReadSeeker(out chan<- io.ReadSeeker, inp ...io.ReadSeeker) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- io.ReadSeeker, inp ...io.ReadSeeker) {
		defer close(done)
		for i := range inp {
			out <- inp[i]
		}
		done <- struct{}{}
	}(cha, out, inp...)
	return cha
}

// JoinReadSeekerSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinReadSeekerSlice(out chan<- io.ReadSeeker, inp ...[]io.ReadSeeker) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- io.ReadSeeker, inp ...[]io.ReadSeeker) {
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

// JoinReadSeekerChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinReadSeekerChan(out chan<- io.ReadSeeker, inp <-chan io.ReadSeeker) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- io.ReadSeeker, inp <-chan io.ReadSeeker) {
		defer close(done)
		for i := range inp {
			out <- i
		}
		done <- struct{}{}
	}(cha, out, inp)
	return cha
}

// DoneReadSeeker returns a channel to receive one signal before close after inp has been drained.
func DoneReadSeeker(inp <-chan io.ReadSeeker) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, inp <-chan io.ReadSeeker) {
		defer close(done)
		for i := range inp {
			_ = i // Drain inp
		}
		done <- struct{}{}
	}(cha, inp)
	return cha
}

// DoneReadSeekerSlice returns a channel which will receive a slice
// of all the ReadSeekers received on inp channel before close.
// Unlike DoneReadSeeker, a full slice is sent once, not just an event.
func DoneReadSeekerSlice(inp <-chan io.ReadSeeker) (done <-chan []io.ReadSeeker) {
	cha := make(chan []io.ReadSeeker)
	go func(inp <-chan io.ReadSeeker, done chan<- []io.ReadSeeker) {
		defer close(done)
		ReadSeekerS := []io.ReadSeeker{}
		for i := range inp {
			ReadSeekerS = append(ReadSeekerS, i)
		}
		done <- ReadSeekerS
	}(inp, cha)
	return cha
}

// DoneReadSeekerFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneReadSeekerFunc(inp <-chan io.ReadSeeker, act func(a io.ReadSeeker)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a io.ReadSeeker) { return }
	}
	go func(done chan<- struct{}, inp <-chan io.ReadSeeker, act func(a io.ReadSeeker)) {
		defer close(done)
		for i := range inp {
			act(i) // Apply action
		}
		done <- struct{}{}
	}(cha, inp, act)
	return cha
}

// PipeReadSeekerBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeReadSeekerBuffer(inp <-chan io.ReadSeeker, cap int) (out <-chan io.ReadSeeker) {
	cha := make(chan io.ReadSeeker, cap)
	go func(out chan<- io.ReadSeeker, inp <-chan io.ReadSeeker) {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}(cha, inp)
	return cha
}

// PipeReadSeekerFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeReadSeekerMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeReadSeekerFunc(inp <-chan io.ReadSeeker, act func(a io.ReadSeeker) io.ReadSeeker) (out <-chan io.ReadSeeker) {
	cha := make(chan io.ReadSeeker)
	if act == nil {
		act = func(a io.ReadSeeker) io.ReadSeeker { return a }
	}
	go func(out chan<- io.ReadSeeker, inp <-chan io.ReadSeeker, act func(a io.ReadSeeker) io.ReadSeeker) {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}(cha, inp, act)
	return cha
}

// PipeReadSeekerFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeReadSeekerFork(inp <-chan io.ReadSeeker) (out1, out2 <-chan io.ReadSeeker) {
	cha1 := make(chan io.ReadSeeker)
	cha2 := make(chan io.ReadSeeker)
	go func(out1, out2 chan<- io.ReadSeeker, inp <-chan io.ReadSeeker) {
		defer close(out1)
		defer close(out2)
		for i := range inp {
			out1 <- i
			out2 <- i
		}
	}(cha1, cha2, inp)
	return cha1, cha2
}

// ReadSeekerTube is the signature for a pipe function.
type ReadSeekerTube func(inp <-chan io.ReadSeeker, out <-chan io.ReadSeeker)

// ReadSeekerDaisy returns a channel to receive all inp after having passed thru tube.
func ReadSeekerDaisy(inp <-chan io.ReadSeeker, tube ReadSeekerTube) (out <-chan io.ReadSeeker) {
	cha := make(chan io.ReadSeeker)
	go tube(inp, cha)
	return cha
}

// ReadSeekerDaisyChain returns a channel to receive all inp after having passed thru all tubes.
func ReadSeekerDaisyChain(inp <-chan io.ReadSeeker, tubes ...ReadSeekerTube) (out <-chan io.ReadSeeker) {
	cha := inp
	for i := range tubes {
		cha = ReadSeekerDaisy(cha, tubes[i])
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
