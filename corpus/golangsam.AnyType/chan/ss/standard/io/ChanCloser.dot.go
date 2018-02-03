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
func MakeCloserChan() (out chan io.Closer) {
	return make(chan io.Closer)
}

func sendCloser(out chan<- io.Closer, inp ...io.Closer) {
	defer close(out)
	for i := range inp {
		out <- inp[i]
	}
}

// ChanCloser returns a channel to receive all inputs before close.
func ChanCloser(inp ...io.Closer) (out <-chan io.Closer) {
	cha := make(chan io.Closer)
	go sendCloser(cha, inp...)
	return cha
}

func sendCloserSlice(out chan<- io.Closer, inp ...[]io.Closer) {
	defer close(out)
	for i := range inp {
		for j := range inp[i] {
			out <- inp[i][j]
		}
	}
}

// ChanCloserSlice returns a channel to receive all inputs before close.
func ChanCloserSlice(inp ...[]io.Closer) (out <-chan io.Closer) {
	cha := make(chan io.Closer)
	go sendCloserSlice(cha, inp...)
	return cha
}

func chanCloserFuncNok(out chan<- io.Closer, act func() (io.Closer, bool)) {
	defer close(out)
	for {
		res, ok := act() // Apply action
		if !ok {
			return
		}
		out <- res
	}
}

// ChanCloserFuncNok returns a channel to receive all results of act until nok before close.
func ChanCloserFuncNok(act func() (io.Closer, bool)) (out <-chan io.Closer) {
	cha := make(chan io.Closer)
	go chanCloserFuncNok(cha, act)
	return cha
}

func chanCloserFuncErr(out chan<- io.Closer, act func() (io.Closer, error)) {
	defer close(out)
	for {
		res, err := act() // Apply action
		if err != nil {
			return
		}
		out <- res
	}
}

// ChanCloserFuncErr returns a channel to receive all results of act until err != nil before close.
func ChanCloserFuncErr(act func() (io.Closer, error)) (out <-chan io.Closer) {
	cha := make(chan io.Closer)
	go chanCloserFuncErr(cha, act)
	return cha
}

func joinCloser(done chan<- struct{}, out chan<- io.Closer, inp ...io.Closer) {
	defer close(done)
	for i := range inp {
		out <- inp[i]
	}
	done <- struct{}{}
}

// JoinCloser sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinCloser(out chan<- io.Closer, inp ...io.Closer) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinCloser(cha, out, inp...)
	return cha
}

func joinCloserSlice(done chan<- struct{}, out chan<- io.Closer, inp ...[]io.Closer) {
	defer close(done)
	for i := range inp {
		for j := range inp[i] {
			out <- inp[i][j]
		}
	}
	done <- struct{}{}
}

// JoinCloserSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinCloserSlice(out chan<- io.Closer, inp ...[]io.Closer) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinCloserSlice(cha, out, inp...)
	return cha
}

func joinCloserChan(done chan<- struct{}, out chan<- io.Closer, inp <-chan io.Closer) {
	defer close(done)
	for i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinCloserChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinCloserChan(out chan<- io.Closer, inp <-chan io.Closer) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinCloserChan(cha, out, inp)
	return cha
}

func doitCloser(done chan<- struct{}, inp <-chan io.Closer) {
	defer close(done)
	for i := range inp {
		_ = i // Drain inp
	}
	done <- struct{}{}
}

// DoneCloser returns a channel to receive one signal before close after inp has been drained.
func DoneCloser(inp <-chan io.Closer) (done <-chan struct{}) {
	cha := make(chan struct{})
	go doitCloser(cha, inp)
	return cha
}

func doitCloserSlice(done chan<- ([]io.Closer), inp <-chan io.Closer) {
	defer close(done)
	CloserS := []io.Closer{}
	for i := range inp {
		CloserS = append(CloserS, i)
	}
	done <- CloserS
}

// DoneCloserSlice returns a channel which will receive a slice
// of all the Closers received on inp channel before close.
// Unlike DoneCloser, a full slice is sent once, not just an event.
func DoneCloserSlice(inp <-chan io.Closer) (done <-chan ([]io.Closer)) {
	cha := make(chan ([]io.Closer))
	go doitCloserSlice(cha, inp)
	return cha
}

func doitCloserFunc(done chan<- struct{}, inp <-chan io.Closer, act func(a io.Closer)) {
	defer close(done)
	for i := range inp {
		act(i) // Apply action
	}
	done <- struct{}{}
}

// DoneCloserFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneCloserFunc(inp <-chan io.Closer, act func(a io.Closer)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a io.Closer) { return }
	}
	go doitCloserFunc(cha, inp, act)
	return cha
}

func pipeCloserBuffer(out chan<- io.Closer, inp <-chan io.Closer) {
	defer close(out)
	for i := range inp {
		out <- i
	}
}

// PipeCloserBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeCloserBuffer(inp <-chan io.Closer, cap int) (out <-chan io.Closer) {
	cha := make(chan io.Closer, cap)
	go pipeCloserBuffer(cha, inp)
	return cha
}

func pipeCloserFunc(out chan<- io.Closer, inp <-chan io.Closer, act func(a io.Closer) io.Closer) {
	defer close(out)
	for i := range inp {
		out <- act(i)
	}
}

// PipeCloserFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeCloserMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeCloserFunc(inp <-chan io.Closer, act func(a io.Closer) io.Closer) (out <-chan io.Closer) {
	cha := make(chan io.Closer)
	if act == nil {
		act = func(a io.Closer) io.Closer { return a }
	}
	go pipeCloserFunc(cha, inp, act)
	return cha
}

func pipeCloserFork(out1, out2 chan<- io.Closer, inp <-chan io.Closer) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		out1 <- i
		out2 <- i
	}
}

// PipeCloserFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeCloserFork(inp <-chan io.Closer) (out1, out2 <-chan io.Closer) {
	cha1 := make(chan io.Closer)
	cha2 := make(chan io.Closer)
	go pipeCloserFork(cha1, cha2, inp)
	return cha1, cha2
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
