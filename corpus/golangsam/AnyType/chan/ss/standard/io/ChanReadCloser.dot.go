// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// MakeReadCloserChan returns a new open channel
// (simply a 'chan io.ReadCloser' that is).
//
// Note: No 'ReadCloser-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myReadCloserPipelineStartsHere := MakeReadCloserChan()
//	// ... lot's of code to design and build Your favourite "myReadCloserWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myReadCloserPipelineStartsHere <- drop
//	}
//	close(myReadCloserPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeReadCloserBuffer) the channel is unbuffered.
//
func MakeReadCloserChan() (out chan io.ReadCloser) {
	return make(chan io.ReadCloser)
}

func sendReadCloser(out chan<- io.ReadCloser, inp ...io.ReadCloser) {
	defer close(out)
	for i := range inp {
		out <- inp[i]
	}
}

// ChanReadCloser returns a channel to receive all inputs before close.
func ChanReadCloser(inp ...io.ReadCloser) (out <-chan io.ReadCloser) {
	cha := make(chan io.ReadCloser)
	go sendReadCloser(cha, inp...)
	return cha
}

func sendReadCloserSlice(out chan<- io.ReadCloser, inp ...[]io.ReadCloser) {
	defer close(out)
	for i := range inp {
		for j := range inp[i] {
			out <- inp[i][j]
		}
	}
}

// ChanReadCloserSlice returns a channel to receive all inputs before close.
func ChanReadCloserSlice(inp ...[]io.ReadCloser) (out <-chan io.ReadCloser) {
	cha := make(chan io.ReadCloser)
	go sendReadCloserSlice(cha, inp...)
	return cha
}

func chanReadCloserFuncNok(out chan<- io.ReadCloser, act func() (io.ReadCloser, bool)) {
	defer close(out)
	for {
		res, ok := act() // Apply action
		if !ok {
			return
		}
		out <- res
	}
}

// ChanReadCloserFuncNok returns a channel to receive all results of act until nok before close.
func ChanReadCloserFuncNok(act func() (io.ReadCloser, bool)) (out <-chan io.ReadCloser) {
	cha := make(chan io.ReadCloser)
	go chanReadCloserFuncNok(cha, act)
	return cha
}

func chanReadCloserFuncErr(out chan<- io.ReadCloser, act func() (io.ReadCloser, error)) {
	defer close(out)
	for {
		res, err := act() // Apply action
		if err != nil {
			return
		}
		out <- res
	}
}

// ChanReadCloserFuncErr returns a channel to receive all results of act until err != nil before close.
func ChanReadCloserFuncErr(act func() (io.ReadCloser, error)) (out <-chan io.ReadCloser) {
	cha := make(chan io.ReadCloser)
	go chanReadCloserFuncErr(cha, act)
	return cha
}

func joinReadCloser(done chan<- struct{}, out chan<- io.ReadCloser, inp ...io.ReadCloser) {
	defer close(done)
	for i := range inp {
		out <- inp[i]
	}
	done <- struct{}{}
}

// JoinReadCloser sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinReadCloser(out chan<- io.ReadCloser, inp ...io.ReadCloser) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinReadCloser(cha, out, inp...)
	return cha
}

func joinReadCloserSlice(done chan<- struct{}, out chan<- io.ReadCloser, inp ...[]io.ReadCloser) {
	defer close(done)
	for i := range inp {
		for j := range inp[i] {
			out <- inp[i][j]
		}
	}
	done <- struct{}{}
}

// JoinReadCloserSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinReadCloserSlice(out chan<- io.ReadCloser, inp ...[]io.ReadCloser) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinReadCloserSlice(cha, out, inp...)
	return cha
}

func joinReadCloserChan(done chan<- struct{}, out chan<- io.ReadCloser, inp <-chan io.ReadCloser) {
	defer close(done)
	for i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinReadCloserChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinReadCloserChan(out chan<- io.ReadCloser, inp <-chan io.ReadCloser) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinReadCloserChan(cha, out, inp)
	return cha
}

func doitReadCloser(done chan<- struct{}, inp <-chan io.ReadCloser) {
	defer close(done)
	for i := range inp {
		_ = i // Drain inp
	}
	done <- struct{}{}
}

// DoneReadCloser returns a channel to receive one signal before close after inp has been drained.
func DoneReadCloser(inp <-chan io.ReadCloser) (done <-chan struct{}) {
	cha := make(chan struct{})
	go doitReadCloser(cha, inp)
	return cha
}

func doitReadCloserSlice(done chan<- ([]io.ReadCloser), inp <-chan io.ReadCloser) {
	defer close(done)
	ReadCloserS := []io.ReadCloser{}
	for i := range inp {
		ReadCloserS = append(ReadCloserS, i)
	}
	done <- ReadCloserS
}

// DoneReadCloserSlice returns a channel which will receive a slice
// of all the ReadClosers received on inp channel before close.
// Unlike DoneReadCloser, a full slice is sent once, not just an event.
func DoneReadCloserSlice(inp <-chan io.ReadCloser) (done <-chan ([]io.ReadCloser)) {
	cha := make(chan ([]io.ReadCloser))
	go doitReadCloserSlice(cha, inp)
	return cha
}

func doitReadCloserFunc(done chan<- struct{}, inp <-chan io.ReadCloser, act func(a io.ReadCloser)) {
	defer close(done)
	for i := range inp {
		act(i) // Apply action
	}
	done <- struct{}{}
}

// DoneReadCloserFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneReadCloserFunc(inp <-chan io.ReadCloser, act func(a io.ReadCloser)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a io.ReadCloser) { return }
	}
	go doitReadCloserFunc(cha, inp, act)
	return cha
}

func pipeReadCloserBuffer(out chan<- io.ReadCloser, inp <-chan io.ReadCloser) {
	defer close(out)
	for i := range inp {
		out <- i
	}
}

// PipeReadCloserBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeReadCloserBuffer(inp <-chan io.ReadCloser, cap int) (out <-chan io.ReadCloser) {
	cha := make(chan io.ReadCloser, cap)
	go pipeReadCloserBuffer(cha, inp)
	return cha
}

func pipeReadCloserFunc(out chan<- io.ReadCloser, inp <-chan io.ReadCloser, act func(a io.ReadCloser) io.ReadCloser) {
	defer close(out)
	for i := range inp {
		out <- act(i)
	}
}

// PipeReadCloserFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeReadCloserMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeReadCloserFunc(inp <-chan io.ReadCloser, act func(a io.ReadCloser) io.ReadCloser) (out <-chan io.ReadCloser) {
	cha := make(chan io.ReadCloser)
	if act == nil {
		act = func(a io.ReadCloser) io.ReadCloser { return a }
	}
	go pipeReadCloserFunc(cha, inp, act)
	return cha
}

func pipeReadCloserFork(out1, out2 chan<- io.ReadCloser, inp <-chan io.ReadCloser) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		out1 <- i
		out2 <- i
	}
}

// PipeReadCloserFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeReadCloserFork(inp <-chan io.ReadCloser) (out1, out2 <-chan io.ReadCloser) {
	cha1 := make(chan io.ReadCloser)
	cha2 := make(chan io.ReadCloser)
	go pipeReadCloserFork(cha1, cha2, inp)
	return cha1, cha2
}

// ReadCloserTube is the signature for a pipe function.
type ReadCloserTube func(inp <-chan io.ReadCloser, out <-chan io.ReadCloser)

// ReadCloserDaisy returns a channel to receive all inp after having passed thru tube.
func ReadCloserDaisy(inp <-chan io.ReadCloser, tube ReadCloserTube) (out <-chan io.ReadCloser) {
	cha := make(chan io.ReadCloser)
	go tube(inp, cha)
	return cha
}

// ReadCloserDaisyChain returns a channel to receive all inp after having passed thru all tubes.
func ReadCloserDaisyChain(inp <-chan io.ReadCloser, tubes ...ReadCloserTube) (out <-chan io.ReadCloser) {
	cha := inp
	for i := range tubes {
		cha = ReadCloserDaisy(cha, tubes[i])
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
