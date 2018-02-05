// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lsm

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"github.com/GoLangsam/container/ccsafe/lsm"
)

// MakeLSMChan returns a new open channel
// (simply a 'chan *lsm.LazyStringerMap' that is).
//
// Note: No 'LSM-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myLSMPipelineStartsHere := MakeLSMChan()
//	// ... lot's of code to design and build Your favourite "myLSMWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myLSMPipelineStartsHere <- drop
//	}
//	close(myLSMPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeLSMBuffer) the channel is unbuffered.
//
func MakeLSMChan() (out chan *lsm.LazyStringerMap) {
	return make(chan *lsm.LazyStringerMap)
}

// ChanLSM returns a channel to receive all inputs before close.
func ChanLSM(inp ...*lsm.LazyStringerMap) (out <-chan *lsm.LazyStringerMap) {
	cha := make(chan *lsm.LazyStringerMap)
	go func(out chan<- *lsm.LazyStringerMap, inp ...*lsm.LazyStringerMap) {
		defer close(out)
		for i := range inp {
			out <- inp[i]
		}
	}(cha, inp...)
	return cha
}

// ChanLSMSlice returns a channel to receive all inputs before close.
func ChanLSMSlice(inp ...[]*lsm.LazyStringerMap) (out <-chan *lsm.LazyStringerMap) {
	cha := make(chan *lsm.LazyStringerMap)
	go func(out chan<- *lsm.LazyStringerMap, inp ...[]*lsm.LazyStringerMap) {
		defer close(out)
		for i := range inp {
			for j := range inp[i] {
				out <- inp[i][j]
			}
		}
	}(cha, inp...)
	return cha
}

// ChanLSMFuncNil returns a channel to receive all results of act until nil before close.
func ChanLSMFuncNil(act func() *lsm.LazyStringerMap) (out <-chan *lsm.LazyStringerMap) {
	cha := make(chan *lsm.LazyStringerMap)
	go func(out chan<- *lsm.LazyStringerMap, act func() *lsm.LazyStringerMap) {
		defer close(out)
		for {
			res := act() // Apply action
			if res == nil {
				return
			}
			out <- res
		}
	}(cha, act)
	return cha
}

// ChanLSMFuncNok returns a channel to receive all results of act until nok before close.
func ChanLSMFuncNok(act func() (*lsm.LazyStringerMap, bool)) (out <-chan *lsm.LazyStringerMap) {
	cha := make(chan *lsm.LazyStringerMap)
	go func(out chan<- *lsm.LazyStringerMap, act func() (*lsm.LazyStringerMap, bool)) {
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

// ChanLSMFuncErr returns a channel to receive all results of act until err != nil before close.
func ChanLSMFuncErr(act func() (*lsm.LazyStringerMap, error)) (out <-chan *lsm.LazyStringerMap) {
	cha := make(chan *lsm.LazyStringerMap)
	go func(out chan<- *lsm.LazyStringerMap, act func() (*lsm.LazyStringerMap, error)) {
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

// JoinLSM sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinLSM(out chan<- *lsm.LazyStringerMap, inp ...*lsm.LazyStringerMap) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- *lsm.LazyStringerMap, inp ...*lsm.LazyStringerMap) {
		defer close(done)
		for i := range inp {
			out <- inp[i]
		}
		done <- struct{}{}
	}(cha, out, inp...)
	return cha
}

// JoinLSMSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinLSMSlice(out chan<- *lsm.LazyStringerMap, inp ...[]*lsm.LazyStringerMap) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- *lsm.LazyStringerMap, inp ...[]*lsm.LazyStringerMap) {
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

// JoinLSMChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinLSMChan(out chan<- *lsm.LazyStringerMap, inp <-chan *lsm.LazyStringerMap) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- *lsm.LazyStringerMap, inp <-chan *lsm.LazyStringerMap) {
		defer close(done)
		for i := range inp {
			out <- i
		}
		done <- struct{}{}
	}(cha, out, inp)
	return cha
}

// DoneLSM returns a channel to receive one signal before close after inp has been drained.
func DoneLSM(inp <-chan *lsm.LazyStringerMap) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, inp <-chan *lsm.LazyStringerMap) {
		defer close(done)
		for i := range inp {
			_ = i // Drain inp
		}
		done <- struct{}{}
	}(cha, inp)
	return cha
}

// DoneLSMSlice returns a channel which will receive a slice
// of all the LSMs received on inp channel before close.
// Unlike DoneLSM, a full slice is sent once, not just an event.
func DoneLSMSlice(inp <-chan *lsm.LazyStringerMap) (done <-chan []*lsm.LazyStringerMap) {
	cha := make(chan []*lsm.LazyStringerMap)
	go func(inp <-chan *lsm.LazyStringerMap, done chan<- []*lsm.LazyStringerMap) {
		defer close(done)
		LSMS := []*lsm.LazyStringerMap{}
		for i := range inp {
			LSMS = append(LSMS, i)
		}
		done <- LSMS
	}(inp, cha)
	return cha
}

// DoneLSMFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneLSMFunc(inp <-chan *lsm.LazyStringerMap, act func(a *lsm.LazyStringerMap)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a *lsm.LazyStringerMap) { return }
	}
	go func(done chan<- struct{}, inp <-chan *lsm.LazyStringerMap, act func(a *lsm.LazyStringerMap)) {
		defer close(done)
		for i := range inp {
			act(i) // Apply action
		}
		done <- struct{}{}
	}(cha, inp, act)
	return cha
}

// PipeLSMBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeLSMBuffer(inp <-chan *lsm.LazyStringerMap, cap int) (out <-chan *lsm.LazyStringerMap) {
	cha := make(chan *lsm.LazyStringerMap, cap)
	go func(out chan<- *lsm.LazyStringerMap, inp <-chan *lsm.LazyStringerMap) {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}(cha, inp)
	return cha
}

// PipeLSMFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeLSMMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeLSMFunc(inp <-chan *lsm.LazyStringerMap, act func(a *lsm.LazyStringerMap) *lsm.LazyStringerMap) (out <-chan *lsm.LazyStringerMap) {
	cha := make(chan *lsm.LazyStringerMap)
	if act == nil {
		act = func(a *lsm.LazyStringerMap) *lsm.LazyStringerMap { return a }
	}
	go func(out chan<- *lsm.LazyStringerMap, inp <-chan *lsm.LazyStringerMap, act func(a *lsm.LazyStringerMap) *lsm.LazyStringerMap) {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}(cha, inp, act)
	return cha
}

// PipeLSMFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeLSMFork(inp <-chan *lsm.LazyStringerMap) (out1, out2 <-chan *lsm.LazyStringerMap) {
	cha1 := make(chan *lsm.LazyStringerMap)
	cha2 := make(chan *lsm.LazyStringerMap)
	go func(out1, out2 chan<- *lsm.LazyStringerMap, inp <-chan *lsm.LazyStringerMap) {
		defer close(out1)
		defer close(out2)
		for i := range inp {
			out1 <- i
			out2 <- i
		}
	}(cha1, cha2, inp)
	return cha1, cha2
}

// LSMTube is the signature for a pipe function.
type LSMTube func(inp <-chan *lsm.LazyStringerMap, out <-chan *lsm.LazyStringerMap)

// LSMDaisy returns a channel to receive all inp after having passed thru tube.
func LSMDaisy(inp <-chan *lsm.LazyStringerMap, tube LSMTube) (out <-chan *lsm.LazyStringerMap) {
	cha := make(chan *lsm.LazyStringerMap)
	go tube(inp, cha)
	return cha
}

// LSMDaisyChain returns a channel to receive all inp after having passed thru all tubes.
func LSMDaisyChain(inp <-chan *lsm.LazyStringerMap, tubes ...LSMTube) (out <-chan *lsm.LazyStringerMap) {
	cha := inp
	for i := range tubes {
		cha = LSMDaisy(cha, tubes[i])
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
