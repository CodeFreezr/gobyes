// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"os"
)

// MakeSignalChan returns a new open channel
// (simply a 'chan os.Signal' that is).
//
// Note: No 'Signal-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var mySignalPipelineStartsHere := MakeSignalChan()
//	// ... lot's of code to design and build Your favourite "mySignalWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		mySignalPipelineStartsHere <- drop
//	}
//	close(mySignalPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeSignalBuffer) the channel is unbuffered.
//
func MakeSignalChan() (out chan os.Signal) {
	return make(chan os.Signal)
}

// ChanSignal returns a channel to receive all inputs before close.
func ChanSignal(inp ...os.Signal) (out <-chan os.Signal) {
	cha := make(chan os.Signal)
	go func(out chan<- os.Signal, inp ...os.Signal) {
		defer close(out)
		for i := range inp {
			out <- inp[i]
		}
	}(cha, inp...)
	return cha
}

// ChanSignalSlice returns a channel to receive all inputs before close.
func ChanSignalSlice(inp ...[]os.Signal) (out <-chan os.Signal) {
	cha := make(chan os.Signal)
	go func(out chan<- os.Signal, inp ...[]os.Signal) {
		defer close(out)
		for i := range inp {
			for j := range inp[i] {
				out <- inp[i][j]
			}
		}
	}(cha, inp...)
	return cha
}

// ChanSignalFuncNok returns a channel to receive all results of act until nok before close.
func ChanSignalFuncNok(act func() (os.Signal, bool)) (out <-chan os.Signal) {
	cha := make(chan os.Signal)
	go func(out chan<- os.Signal, act func() (os.Signal, bool)) {
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

// ChanSignalFuncErr returns a channel to receive all results of act until err != nil before close.
func ChanSignalFuncErr(act func() (os.Signal, error)) (out <-chan os.Signal) {
	cha := make(chan os.Signal)
	go func(out chan<- os.Signal, act func() (os.Signal, error)) {
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

// JoinSignal sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinSignal(out chan<- os.Signal, inp ...os.Signal) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- os.Signal, inp ...os.Signal) {
		defer close(done)
		for i := range inp {
			out <- inp[i]
		}
		done <- struct{}{}
	}(cha, out, inp...)
	return cha
}

// JoinSignalSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinSignalSlice(out chan<- os.Signal, inp ...[]os.Signal) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- os.Signal, inp ...[]os.Signal) {
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

// JoinSignalChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinSignalChan(out chan<- os.Signal, inp <-chan os.Signal) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- os.Signal, inp <-chan os.Signal) {
		defer close(done)
		for i := range inp {
			out <- i
		}
		done <- struct{}{}
	}(cha, out, inp)
	return cha
}

// DoneSignal returns a channel to receive one signal before close after inp has been drained.
func DoneSignal(inp <-chan os.Signal) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, inp <-chan os.Signal) {
		defer close(done)
		for i := range inp {
			_ = i // Drain inp
		}
		done <- struct{}{}
	}(cha, inp)
	return cha
}

// DoneSignalSlice returns a channel which will receive a slice
// of all the Signals received on inp channel before close.
// Unlike DoneSignal, a full slice is sent once, not just an event.
func DoneSignalSlice(inp <-chan os.Signal) (done <-chan []os.Signal) {
	cha := make(chan []os.Signal)
	go func(inp <-chan os.Signal, done chan<- []os.Signal) {
		defer close(done)
		SignalS := []os.Signal{}
		for i := range inp {
			SignalS = append(SignalS, i)
		}
		done <- SignalS
	}(inp, cha)
	return cha
}

// DoneSignalFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneSignalFunc(inp <-chan os.Signal, act func(a os.Signal)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a os.Signal) { return }
	}
	go func(done chan<- struct{}, inp <-chan os.Signal, act func(a os.Signal)) {
		defer close(done)
		for i := range inp {
			act(i) // Apply action
		}
		done <- struct{}{}
	}(cha, inp, act)
	return cha
}

// PipeSignalBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeSignalBuffer(inp <-chan os.Signal, cap int) (out <-chan os.Signal) {
	cha := make(chan os.Signal, cap)
	go func(out chan<- os.Signal, inp <-chan os.Signal) {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}(cha, inp)
	return cha
}

// PipeSignalFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeSignalMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeSignalFunc(inp <-chan os.Signal, act func(a os.Signal) os.Signal) (out <-chan os.Signal) {
	cha := make(chan os.Signal)
	if act == nil {
		act = func(a os.Signal) os.Signal { return a }
	}
	go func(out chan<- os.Signal, inp <-chan os.Signal, act func(a os.Signal) os.Signal) {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}(cha, inp, act)
	return cha
}

// PipeSignalFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeSignalFork(inp <-chan os.Signal) (out1, out2 <-chan os.Signal) {
	cha1 := make(chan os.Signal)
	cha2 := make(chan os.Signal)
	go func(out1, out2 chan<- os.Signal, inp <-chan os.Signal) {
		defer close(out1)
		defer close(out2)
		for i := range inp {
			out1 <- i
			out2 <- i
		}
	}(cha1, cha2, inp)
	return cha1, cha2
}

// SignalTube is the signature for a pipe function.
type SignalTube func(inp <-chan os.Signal, out <-chan os.Signal)

// SignalDaisy returns a channel to receive all inp after having passed thru tube.
func SignalDaisy(inp <-chan os.Signal, tube SignalTube) (out <-chan os.Signal) {
	cha := make(chan os.Signal)
	go tube(inp, cha)
	return cha
}

// SignalDaisyChain returns a channel to receive all inp after having passed thru all tubes.
func SignalDaisyChain(inp <-chan os.Signal, tubes ...SignalTube) (out <-chan os.Signal) {
	cha := inp
	for i := range tubes {
		cha = SignalDaisy(cha, tubes[i])
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
