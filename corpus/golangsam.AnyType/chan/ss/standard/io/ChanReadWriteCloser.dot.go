// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// MakeReadWriteCloserChan returns a new open channel
// (simply a 'chan io.ReadWriteCloser' that is).
//
// Note: No 'ReadWriteCloser-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myReadWriteCloserPipelineStartsHere := MakeReadWriteCloserChan()
//	// ... lot's of code to design and build Your favourite "myReadWriteCloserWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myReadWriteCloserPipelineStartsHere <- drop
//	}
//	close(myReadWriteCloserPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeReadWriteCloserBuffer) the channel is unbuffered.
//
func MakeReadWriteCloserChan() (out chan io.ReadWriteCloser) {
	return make(chan io.ReadWriteCloser)
}

func sendReadWriteCloser(out chan<- io.ReadWriteCloser, inp ...io.ReadWriteCloser) {
	defer close(out)
	for i := range inp {
		out <- inp[i]
	}
}

// ChanReadWriteCloser returns a channel to receive all inputs before close.
func ChanReadWriteCloser(inp ...io.ReadWriteCloser) (out <-chan io.ReadWriteCloser) {
	cha := make(chan io.ReadWriteCloser)
	go sendReadWriteCloser(cha, inp...)
	return cha
}

func sendReadWriteCloserSlice(out chan<- io.ReadWriteCloser, inp ...[]io.ReadWriteCloser) {
	defer close(out)
	for i := range inp {
		for j := range inp[i] {
			out <- inp[i][j]
		}
	}
}

// ChanReadWriteCloserSlice returns a channel to receive all inputs before close.
func ChanReadWriteCloserSlice(inp ...[]io.ReadWriteCloser) (out <-chan io.ReadWriteCloser) {
	cha := make(chan io.ReadWriteCloser)
	go sendReadWriteCloserSlice(cha, inp...)
	return cha
}

func chanReadWriteCloserFuncNok(out chan<- io.ReadWriteCloser, act func() (io.ReadWriteCloser, bool)) {
	defer close(out)
	for {
		res, ok := act() // Apply action
		if !ok {
			return
		}
		out <- res
	}
}

// ChanReadWriteCloserFuncNok returns a channel to receive all results of act until nok before close.
func ChanReadWriteCloserFuncNok(act func() (io.ReadWriteCloser, bool)) (out <-chan io.ReadWriteCloser) {
	cha := make(chan io.ReadWriteCloser)
	go chanReadWriteCloserFuncNok(cha, act)
	return cha
}

func chanReadWriteCloserFuncErr(out chan<- io.ReadWriteCloser, act func() (io.ReadWriteCloser, error)) {
	defer close(out)
	for {
		res, err := act() // Apply action
		if err != nil {
			return
		}
		out <- res
	}
}

// ChanReadWriteCloserFuncErr returns a channel to receive all results of act until err != nil before close.
func ChanReadWriteCloserFuncErr(act func() (io.ReadWriteCloser, error)) (out <-chan io.ReadWriteCloser) {
	cha := make(chan io.ReadWriteCloser)
	go chanReadWriteCloserFuncErr(cha, act)
	return cha
}

func joinReadWriteCloser(done chan<- struct{}, out chan<- io.ReadWriteCloser, inp ...io.ReadWriteCloser) {
	defer close(done)
	for i := range inp {
		out <- inp[i]
	}
	done <- struct{}{}
}

// JoinReadWriteCloser sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinReadWriteCloser(out chan<- io.ReadWriteCloser, inp ...io.ReadWriteCloser) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinReadWriteCloser(cha, out, inp...)
	return cha
}

func joinReadWriteCloserSlice(done chan<- struct{}, out chan<- io.ReadWriteCloser, inp ...[]io.ReadWriteCloser) {
	defer close(done)
	for i := range inp {
		for j := range inp[i] {
			out <- inp[i][j]
		}
	}
	done <- struct{}{}
}

// JoinReadWriteCloserSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinReadWriteCloserSlice(out chan<- io.ReadWriteCloser, inp ...[]io.ReadWriteCloser) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinReadWriteCloserSlice(cha, out, inp...)
	return cha
}

func joinReadWriteCloserChan(done chan<- struct{}, out chan<- io.ReadWriteCloser, inp <-chan io.ReadWriteCloser) {
	defer close(done)
	for i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinReadWriteCloserChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinReadWriteCloserChan(out chan<- io.ReadWriteCloser, inp <-chan io.ReadWriteCloser) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinReadWriteCloserChan(cha, out, inp)
	return cha
}

func doitReadWriteCloser(done chan<- struct{}, inp <-chan io.ReadWriteCloser) {
	defer close(done)
	for i := range inp {
		_ = i // Drain inp
	}
	done <- struct{}{}
}

// DoneReadWriteCloser returns a channel to receive one signal before close after inp has been drained.
func DoneReadWriteCloser(inp <-chan io.ReadWriteCloser) (done <-chan struct{}) {
	cha := make(chan struct{})
	go doitReadWriteCloser(cha, inp)
	return cha
}

func doitReadWriteCloserSlice(done chan<- ([]io.ReadWriteCloser), inp <-chan io.ReadWriteCloser) {
	defer close(done)
	ReadWriteCloserS := []io.ReadWriteCloser{}
	for i := range inp {
		ReadWriteCloserS = append(ReadWriteCloserS, i)
	}
	done <- ReadWriteCloserS
}

// DoneReadWriteCloserSlice returns a channel which will receive a slice
// of all the ReadWriteClosers received on inp channel before close.
// Unlike DoneReadWriteCloser, a full slice is sent once, not just an event.
func DoneReadWriteCloserSlice(inp <-chan io.ReadWriteCloser) (done <-chan ([]io.ReadWriteCloser)) {
	cha := make(chan ([]io.ReadWriteCloser))
	go doitReadWriteCloserSlice(cha, inp)
	return cha
}

func doitReadWriteCloserFunc(done chan<- struct{}, inp <-chan io.ReadWriteCloser, act func(a io.ReadWriteCloser)) {
	defer close(done)
	for i := range inp {
		act(i) // Apply action
	}
	done <- struct{}{}
}

// DoneReadWriteCloserFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneReadWriteCloserFunc(inp <-chan io.ReadWriteCloser, act func(a io.ReadWriteCloser)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a io.ReadWriteCloser) { return }
	}
	go doitReadWriteCloserFunc(cha, inp, act)
	return cha
}

func pipeReadWriteCloserBuffer(out chan<- io.ReadWriteCloser, inp <-chan io.ReadWriteCloser) {
	defer close(out)
	for i := range inp {
		out <- i
	}
}

// PipeReadWriteCloserBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeReadWriteCloserBuffer(inp <-chan io.ReadWriteCloser, cap int) (out <-chan io.ReadWriteCloser) {
	cha := make(chan io.ReadWriteCloser, cap)
	go pipeReadWriteCloserBuffer(cha, inp)
	return cha
}

func pipeReadWriteCloserFunc(out chan<- io.ReadWriteCloser, inp <-chan io.ReadWriteCloser, act func(a io.ReadWriteCloser) io.ReadWriteCloser) {
	defer close(out)
	for i := range inp {
		out <- act(i)
	}
}

// PipeReadWriteCloserFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeReadWriteCloserMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeReadWriteCloserFunc(inp <-chan io.ReadWriteCloser, act func(a io.ReadWriteCloser) io.ReadWriteCloser) (out <-chan io.ReadWriteCloser) {
	cha := make(chan io.ReadWriteCloser)
	if act == nil {
		act = func(a io.ReadWriteCloser) io.ReadWriteCloser { return a }
	}
	go pipeReadWriteCloserFunc(cha, inp, act)
	return cha
}

func pipeReadWriteCloserFork(out1, out2 chan<- io.ReadWriteCloser, inp <-chan io.ReadWriteCloser) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		out1 <- i
		out2 <- i
	}
}

// PipeReadWriteCloserFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeReadWriteCloserFork(inp <-chan io.ReadWriteCloser) (out1, out2 <-chan io.ReadWriteCloser) {
	cha1 := make(chan io.ReadWriteCloser)
	cha2 := make(chan io.ReadWriteCloser)
	go pipeReadWriteCloserFork(cha1, cha2, inp)
	return cha1, cha2
}

// ReadWriteCloserTube is the signature for a pipe function.
type ReadWriteCloserTube func(inp <-chan io.ReadWriteCloser, out <-chan io.ReadWriteCloser)

// ReadWriteCloserDaisy returns a channel to receive all inp after having passed thru tube.
func ReadWriteCloserDaisy(inp <-chan io.ReadWriteCloser, tube ReadWriteCloserTube) (out <-chan io.ReadWriteCloser) {
	cha := make(chan io.ReadWriteCloser)
	go tube(inp, cha)
	return cha
}

// ReadWriteCloserDaisyChain returns a channel to receive all inp after having passed thru all tubes.
func ReadWriteCloserDaisyChain(inp <-chan io.ReadWriteCloser, tubes ...ReadWriteCloserTube) (out <-chan io.ReadWriteCloser) {
	cha := inp
	for i := range tubes {
		cha = ReadWriteCloserDaisy(cha, tubes[i])
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
