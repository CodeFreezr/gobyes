// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// MakeReadWriteSeekerChan returns a new open channel
// (simply a 'chan io.ReadWriteSeeker' that is).
//
// Note: No 'ReadWriteSeeker-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myReadWriteSeekerPipelineStartsHere := MakeReadWriteSeekerChan()
//	// ... lot's of code to design and build Your favourite "myReadWriteSeekerWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myReadWriteSeekerPipelineStartsHere <- drop
//	}
//	close(myReadWriteSeekerPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeReadWriteSeekerBuffer) the channel is unbuffered.
//
func MakeReadWriteSeekerChan() (out chan io.ReadWriteSeeker) {
	return make(chan io.ReadWriteSeeker)
}

// ChanReadWriteSeeker returns a channel to receive all inputs before close.
func ChanReadWriteSeeker(inp ...io.ReadWriteSeeker) (out <-chan io.ReadWriteSeeker) {
	cha := make(chan io.ReadWriteSeeker)
	go func(out chan<- io.ReadWriteSeeker, inp ...io.ReadWriteSeeker) {
		defer close(out)
		for i := range inp {
			out <- inp[i]
		}
	}(cha, inp...)
	return cha
}

// ChanReadWriteSeekerSlice returns a channel to receive all inputs before close.
func ChanReadWriteSeekerSlice(inp ...[]io.ReadWriteSeeker) (out <-chan io.ReadWriteSeeker) {
	cha := make(chan io.ReadWriteSeeker)
	go func(out chan<- io.ReadWriteSeeker, inp ...[]io.ReadWriteSeeker) {
		defer close(out)
		for i := range inp {
			for j := range inp[i] {
				out <- inp[i][j]
			}
		}
	}(cha, inp...)
	return cha
}

// ChanReadWriteSeekerFuncNok returns a channel to receive all results of act until nok before close.
func ChanReadWriteSeekerFuncNok(act func() (io.ReadWriteSeeker, bool)) (out <-chan io.ReadWriteSeeker) {
	cha := make(chan io.ReadWriteSeeker)
	go func(out chan<- io.ReadWriteSeeker, act func() (io.ReadWriteSeeker, bool)) {
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

// ChanReadWriteSeekerFuncErr returns a channel to receive all results of act until err != nil before close.
func ChanReadWriteSeekerFuncErr(act func() (io.ReadWriteSeeker, error)) (out <-chan io.ReadWriteSeeker) {
	cha := make(chan io.ReadWriteSeeker)
	go func(out chan<- io.ReadWriteSeeker, act func() (io.ReadWriteSeeker, error)) {
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

// JoinReadWriteSeeker sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinReadWriteSeeker(out chan<- io.ReadWriteSeeker, inp ...io.ReadWriteSeeker) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- io.ReadWriteSeeker, inp ...io.ReadWriteSeeker) {
		defer close(done)
		for i := range inp {
			out <- inp[i]
		}
		done <- struct{}{}
	}(cha, out, inp...)
	return cha
}

// JoinReadWriteSeekerSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinReadWriteSeekerSlice(out chan<- io.ReadWriteSeeker, inp ...[]io.ReadWriteSeeker) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- io.ReadWriteSeeker, inp ...[]io.ReadWriteSeeker) {
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

// JoinReadWriteSeekerChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinReadWriteSeekerChan(out chan<- io.ReadWriteSeeker, inp <-chan io.ReadWriteSeeker) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- io.ReadWriteSeeker, inp <-chan io.ReadWriteSeeker) {
		defer close(done)
		for i := range inp {
			out <- i
		}
		done <- struct{}{}
	}(cha, out, inp)
	return cha
}

// DoneReadWriteSeeker returns a channel to receive one signal before close after inp has been drained.
func DoneReadWriteSeeker(inp <-chan io.ReadWriteSeeker) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, inp <-chan io.ReadWriteSeeker) {
		defer close(done)
		for i := range inp {
			_ = i // Drain inp
		}
		done <- struct{}{}
	}(cha, inp)
	return cha
}

// DoneReadWriteSeekerSlice returns a channel which will receive a slice
// of all the ReadWriteSeekers received on inp channel before close.
// Unlike DoneReadWriteSeeker, a full slice is sent once, not just an event.
func DoneReadWriteSeekerSlice(inp <-chan io.ReadWriteSeeker) (done <-chan []io.ReadWriteSeeker) {
	cha := make(chan []io.ReadWriteSeeker)
	go func(inp <-chan io.ReadWriteSeeker, done chan<- []io.ReadWriteSeeker) {
		defer close(done)
		ReadWriteSeekerS := []io.ReadWriteSeeker{}
		for i := range inp {
			ReadWriteSeekerS = append(ReadWriteSeekerS, i)
		}
		done <- ReadWriteSeekerS
	}(inp, cha)
	return cha
}

// DoneReadWriteSeekerFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneReadWriteSeekerFunc(inp <-chan io.ReadWriteSeeker, act func(a io.ReadWriteSeeker)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a io.ReadWriteSeeker) { return }
	}
	go func(done chan<- struct{}, inp <-chan io.ReadWriteSeeker, act func(a io.ReadWriteSeeker)) {
		defer close(done)
		for i := range inp {
			act(i) // Apply action
		}
		done <- struct{}{}
	}(cha, inp, act)
	return cha
}

// PipeReadWriteSeekerBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeReadWriteSeekerBuffer(inp <-chan io.ReadWriteSeeker, cap int) (out <-chan io.ReadWriteSeeker) {
	cha := make(chan io.ReadWriteSeeker, cap)
	go func(out chan<- io.ReadWriteSeeker, inp <-chan io.ReadWriteSeeker) {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}(cha, inp)
	return cha
}

// PipeReadWriteSeekerFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeReadWriteSeekerMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeReadWriteSeekerFunc(inp <-chan io.ReadWriteSeeker, act func(a io.ReadWriteSeeker) io.ReadWriteSeeker) (out <-chan io.ReadWriteSeeker) {
	cha := make(chan io.ReadWriteSeeker)
	if act == nil {
		act = func(a io.ReadWriteSeeker) io.ReadWriteSeeker { return a }
	}
	go func(out chan<- io.ReadWriteSeeker, inp <-chan io.ReadWriteSeeker, act func(a io.ReadWriteSeeker) io.ReadWriteSeeker) {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}(cha, inp, act)
	return cha
}

// PipeReadWriteSeekerFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeReadWriteSeekerFork(inp <-chan io.ReadWriteSeeker) (out1, out2 <-chan io.ReadWriteSeeker) {
	cha1 := make(chan io.ReadWriteSeeker)
	cha2 := make(chan io.ReadWriteSeeker)
	go func(out1, out2 chan<- io.ReadWriteSeeker, inp <-chan io.ReadWriteSeeker) {
		defer close(out1)
		defer close(out2)
		for i := range inp {
			out1 <- i
			out2 <- i
		}
	}(cha1, cha2, inp)
	return cha1, cha2
}

// ReadWriteSeekerTube is the signature for a pipe function.
type ReadWriteSeekerTube func(inp <-chan io.ReadWriteSeeker, out <-chan io.ReadWriteSeeker)

// ReadWriteSeekerDaisy returns a channel to receive all inp after having passed thru tube.
func ReadWriteSeekerDaisy(inp <-chan io.ReadWriteSeeker, tube ReadWriteSeekerTube) (out <-chan io.ReadWriteSeeker) {
	cha := make(chan io.ReadWriteSeeker)
	go tube(inp, cha)
	return cha
}

// ReadWriteSeekerDaisyChain returns a channel to receive all inp after having passed thru all tubes.
func ReadWriteSeekerDaisyChain(inp <-chan io.ReadWriteSeeker, tubes ...ReadWriteSeekerTube) (out <-chan io.ReadWriteSeeker) {
	cha := inp
	for i := range tubes {
		cha = ReadWriteSeekerDaisy(cha, tubes[i])
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
