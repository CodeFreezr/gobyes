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
func MakeSignalChan() chan os.Signal {
	return make(chan os.Signal)
}

// ChanSignal returns a channel to receive all inputs before close.
func ChanSignal(inp ...os.Signal) chan os.Signal {
	out := make(chan os.Signal)
	go func() {
		defer close(out)
		for i := range inp {
			out <- inp[i]
		}
	}()
	return out
}

// ChanSignalSlice returns a channel to receive all inputs before close.
func ChanSignalSlice(inp ...[]os.Signal) chan os.Signal {
	out := make(chan os.Signal)
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

// ChanSignalFuncNok returns a channel to receive all results of act until nok before close.
func ChanSignalFuncNok(act func() (os.Signal, bool)) <-chan os.Signal {
	out := make(chan os.Signal)
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

// ChanSignalFuncErr returns a channel to receive all results of act until err != nil before close.
func ChanSignalFuncErr(act func() (os.Signal, error)) <-chan os.Signal {
	out := make(chan os.Signal)
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

// JoinSignal sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinSignal(out chan<- os.Signal, inp ...os.Signal) chan struct{} {
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

// JoinSignalSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinSignalSlice(out chan<- os.Signal, inp ...[]os.Signal) chan struct{} {
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

// JoinSignalChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinSignalChan(out chan<- os.Signal, inp <-chan os.Signal) chan struct{} {
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

// DoneSignal returns a channel to receive one signal before close after inp has been drained.
func DoneSignal(inp <-chan os.Signal) chan struct{} {
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

// DoneSignalSlice returns a channel which will receive a slice
// of all the Signals received on inp channel before close.
// Unlike DoneSignal, a full slice is sent once, not just an event.
func DoneSignalSlice(inp <-chan os.Signal) chan []os.Signal {
	done := make(chan []os.Signal)
	go func() {
		defer close(done)
		SignalS := []os.Signal{}
		for i := range inp {
			SignalS = append(SignalS, i)
		}
		done <- SignalS
	}()
	return done
}

// DoneSignalFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneSignalFunc(inp <-chan os.Signal, act func(a os.Signal)) chan struct{} {
	done := make(chan struct{})
	if act == nil {
		act = func(a os.Signal) { return }
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

// PipeSignalBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeSignalBuffer(inp <-chan os.Signal, cap int) chan os.Signal {
	out := make(chan os.Signal, cap)
	go func() {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}()
	return out
}

// PipeSignalFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeSignalMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeSignalFunc(inp <-chan os.Signal, act func(a os.Signal) os.Signal) chan os.Signal {
	out := make(chan os.Signal)
	if act == nil {
		act = func(a os.Signal) os.Signal { return a }
	}
	go func() {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}()
	return out
}

// PipeSignalFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeSignalFork(inp <-chan os.Signal) (chan os.Signal, chan os.Signal) {
	out1 := make(chan os.Signal)
	out2 := make(chan os.Signal)
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
