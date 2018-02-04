// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package list

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	list "container/list"
)

// MakeChan returns a new open channel
// (simply a 'chan *list.List' that is).
//
// Note: No '-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myPipelineStartsHere := MakeChan()
//	// ... lot's of code to design and build Your favourite "myWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myPipelineStartsHere <- drop
//	}
//	close(myPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeBuffer) the channel is unbuffered.
//
func MakeChan() (out chan *list.List) {
	return make(chan *list.List)
}

// Chan returns a channel to receive all inputs before close.
func Chan(inp ...*list.List) (out <-chan *list.List) {
	cha := make(chan *list.List)
	go func(out chan<- *list.List, inp ...*list.List) {
		defer close(out)
		for i := range inp {
			out <- inp[i]
		}
	}(cha, inp...)
	return cha
}

// ChanSlice returns a channel to receive all inputs before close.
func ChanSlice(inp ...[]*list.List) (out <-chan *list.List) {
	cha := make(chan *list.List)
	go func(out chan<- *list.List, inp ...[]*list.List) {
		defer close(out)
		for i := range inp {
			for j := range inp[i] {
				out <- inp[i][j]
			}
		}
	}(cha, inp...)
	return cha
}

// ChanFuncNil returns a channel to receive all results of act until nil before close.
func ChanFuncNil(act func() *list.List) (out <-chan *list.List) {
	cha := make(chan *list.List)
	go func(out chan<- *list.List, act func() *list.List) {
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

// ChanFuncNok returns a channel to receive all results of act until nok before close.
func ChanFuncNok(act func() (*list.List, bool)) (out <-chan *list.List) {
	cha := make(chan *list.List)
	go func(out chan<- *list.List, act func() (*list.List, bool)) {
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

// ChanFuncErr returns a channel to receive all results of act until err != nil before close.
func ChanFuncErr(act func() (*list.List, error)) (out <-chan *list.List) {
	cha := make(chan *list.List)
	go func(out chan<- *list.List, act func() (*list.List, error)) {
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

// Join sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func Join(out chan<- *list.List, inp ...*list.List) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- *list.List, inp ...*list.List) {
		defer close(done)
		for i := range inp {
			out <- inp[i]
		}
		done <- struct{}{}
	}(cha, out, inp...)
	return cha
}

// JoinSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinSlice(out chan<- *list.List, inp ...[]*list.List) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- *list.List, inp ...[]*list.List) {
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

// JoinChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinChan(out chan<- *list.List, inp <-chan *list.List) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- *list.List, inp <-chan *list.List) {
		defer close(done)
		for i := range inp {
			out <- i
		}
		done <- struct{}{}
	}(cha, out, inp)
	return cha
}

// Done returns a channel to receive one signal before close after inp has been drained.
func Done(inp <-chan *list.List) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, inp <-chan *list.List) {
		defer close(done)
		for i := range inp {
			_ = i // Drain inp
		}
		done <- struct{}{}
	}(cha, inp)
	return cha
}

// DoneSlice returns a channel which will receive a slice
// of all the s received on inp channel before close.
// Unlike Done, a full slice is sent once, not just an event.
func DoneSlice(inp <-chan *list.List) (done <-chan []*list.List) {
	cha := make(chan []*list.List)
	go func(inp <-chan *list.List, done chan<- []*list.List) {
		defer close(done)
		S := []*list.List{}
		for i := range inp {
			S = append(S, i)
		}
		done <- S
	}(inp, cha)
	return cha
}

// DoneFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneFunc(inp <-chan *list.List, act func(a *list.List)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a *list.List) { return }
	}
	go func(done chan<- struct{}, inp <-chan *list.List, act func(a *list.List)) {
		defer close(done)
		for i := range inp {
			act(i) // Apply action
		}
		done <- struct{}{}
	}(cha, inp, act)
	return cha
}

// PipeBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeBuffer(inp <-chan *list.List, cap int) (out <-chan *list.List) {
	cha := make(chan *list.List, cap)
	go func(out chan<- *list.List, inp <-chan *list.List) {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}(cha, inp)
	return cha
}

// PipeFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeFunc(inp <-chan *list.List, act func(a *list.List) *list.List) (out <-chan *list.List) {
	cha := make(chan *list.List)
	if act == nil {
		act = func(a *list.List) *list.List { return a }
	}
	go func(out chan<- *list.List, inp <-chan *list.List, act func(a *list.List) *list.List) {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}(cha, inp, act)
	return cha
}

// PipeFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeFork(inp <-chan *list.List) (out1, out2 <-chan *list.List) {
	cha1 := make(chan *list.List)
	cha2 := make(chan *list.List)
	go func(out1, out2 chan<- *list.List, inp <-chan *list.List) {
		defer close(out1)
		defer close(out2)
		for i := range inp {
			out1 <- i
			out2 <- i
		}
	}(cha1, cha2, inp)
	return cha1, cha2
}

// Tube is the signature for a pipe function.
type Tube func(inp <-chan *list.List, out <-chan *list.List)

// Daisy returns a channel to receive all inp after having passed thru tube.
func Daisy(inp <-chan *list.List, tube Tube) (out <-chan *list.List) {
	cha := make(chan *list.List)
	go tube(inp, cha)
	return cha
}

// DaisyChain returns a channel to receive all inp after having passed thru all tubes.
func DaisyChain(inp <-chan *list.List, tubes ...Tube) (out <-chan *list.List) {
	cha := inp
	for i := range tubes {
		cha = Daisy(cha, tubes[i])
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
