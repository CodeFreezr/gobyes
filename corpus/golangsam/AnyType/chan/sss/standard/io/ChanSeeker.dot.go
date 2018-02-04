// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// MakeSeekerChan returns a new open channel
// (simply a 'chan io.Seeker' that is).
//
// Note: No 'Seeker-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var mySeekerPipelineStartsHere := MakeSeekerChan()
//	// ... lot's of code to design and build Your favourite "mySeekerWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		mySeekerPipelineStartsHere <- drop
//	}
//	close(mySeekerPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeSeekerBuffer) the channel is unbuffered.
//
func MakeSeekerChan() (out chan io.Seeker) {
	return make(chan io.Seeker)
}

// ChanSeeker returns a channel to receive all inputs before close.
func ChanSeeker(inp ...io.Seeker) (out <-chan io.Seeker) {
	cha := make(chan io.Seeker)
	go func(out chan<- io.Seeker, inp ...io.Seeker) {
		defer close(out)
		for i := range inp {
			out <- inp[i]
		}
	}(cha, inp...)
	return cha
}

// ChanSeekerSlice returns a channel to receive all inputs before close.
func ChanSeekerSlice(inp ...[]io.Seeker) (out <-chan io.Seeker) {
	cha := make(chan io.Seeker)
	go func(out chan<- io.Seeker, inp ...[]io.Seeker) {
		defer close(out)
		for i := range inp {
			for j := range inp[i] {
				out <- inp[i][j]
			}
		}
	}(cha, inp...)
	return cha
}

// ChanSeekerFuncNok returns a channel to receive all results of act until nok before close.
func ChanSeekerFuncNok(act func() (io.Seeker, bool)) (out <-chan io.Seeker) {
	cha := make(chan io.Seeker)
	go func(out chan<- io.Seeker, act func() (io.Seeker, bool)) {
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

// ChanSeekerFuncErr returns a channel to receive all results of act until err != nil before close.
func ChanSeekerFuncErr(act func() (io.Seeker, error)) (out <-chan io.Seeker) {
	cha := make(chan io.Seeker)
	go func(out chan<- io.Seeker, act func() (io.Seeker, error)) {
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

// JoinSeeker sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinSeeker(out chan<- io.Seeker, inp ...io.Seeker) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- io.Seeker, inp ...io.Seeker) {
		defer close(done)
		for i := range inp {
			out <- inp[i]
		}
		done <- struct{}{}
	}(cha, out, inp...)
	return cha
}

// JoinSeekerSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinSeekerSlice(out chan<- io.Seeker, inp ...[]io.Seeker) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- io.Seeker, inp ...[]io.Seeker) {
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

// JoinSeekerChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinSeekerChan(out chan<- io.Seeker, inp <-chan io.Seeker) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- io.Seeker, inp <-chan io.Seeker) {
		defer close(done)
		for i := range inp {
			out <- i
		}
		done <- struct{}{}
	}(cha, out, inp)
	return cha
}

// DoneSeeker returns a channel to receive one signal before close after inp has been drained.
func DoneSeeker(inp <-chan io.Seeker) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, inp <-chan io.Seeker) {
		defer close(done)
		for i := range inp {
			_ = i // Drain inp
		}
		done <- struct{}{}
	}(cha, inp)
	return cha
}

// DoneSeekerSlice returns a channel which will receive a slice
// of all the Seekers received on inp channel before close.
// Unlike DoneSeeker, a full slice is sent once, not just an event.
func DoneSeekerSlice(inp <-chan io.Seeker) (done <-chan []io.Seeker) {
	cha := make(chan []io.Seeker)
	go func(inp <-chan io.Seeker, done chan<- []io.Seeker) {
		defer close(done)
		SeekerS := []io.Seeker{}
		for i := range inp {
			SeekerS = append(SeekerS, i)
		}
		done <- SeekerS
	}(inp, cha)
	return cha
}

// DoneSeekerFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneSeekerFunc(inp <-chan io.Seeker, act func(a io.Seeker)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a io.Seeker) { return }
	}
	go func(done chan<- struct{}, inp <-chan io.Seeker, act func(a io.Seeker)) {
		defer close(done)
		for i := range inp {
			act(i) // Apply action
		}
		done <- struct{}{}
	}(cha, inp, act)
	return cha
}

// PipeSeekerBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeSeekerBuffer(inp <-chan io.Seeker, cap int) (out <-chan io.Seeker) {
	cha := make(chan io.Seeker, cap)
	go func(out chan<- io.Seeker, inp <-chan io.Seeker) {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}(cha, inp)
	return cha
}

// PipeSeekerFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeSeekerMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeSeekerFunc(inp <-chan io.Seeker, act func(a io.Seeker) io.Seeker) (out <-chan io.Seeker) {
	cha := make(chan io.Seeker)
	if act == nil {
		act = func(a io.Seeker) io.Seeker { return a }
	}
	go func(out chan<- io.Seeker, inp <-chan io.Seeker, act func(a io.Seeker) io.Seeker) {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}(cha, inp, act)
	return cha
}

// PipeSeekerFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeSeekerFork(inp <-chan io.Seeker) (out1, out2 <-chan io.Seeker) {
	cha1 := make(chan io.Seeker)
	cha2 := make(chan io.Seeker)
	go func(out1, out2 chan<- io.Seeker, inp <-chan io.Seeker) {
		defer close(out1)
		defer close(out2)
		for i := range inp {
			out1 <- i
			out2 <- i
		}
	}(cha1, cha2, inp)
	return cha1, cha2
}

// SeekerTube is the signature for a pipe function.
type SeekerTube func(inp <-chan io.Seeker, out <-chan io.Seeker)

// SeekerDaisy returns a channel to receive all inp after having passed thru tube.
func SeekerDaisy(inp <-chan io.Seeker, tube SeekerTube) (out <-chan io.Seeker) {
	cha := make(chan io.Seeker)
	go tube(inp, cha)
	return cha
}

// SeekerDaisyChain returns a channel to receive all inp after having passed thru all tubes.
func SeekerDaisyChain(inp <-chan io.Seeker, tubes ...SeekerTube) (out <-chan io.Seeker) {
	cha := inp
	for i := range tubes {
		cha = SeekerDaisy(cha, tubes[i])
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
