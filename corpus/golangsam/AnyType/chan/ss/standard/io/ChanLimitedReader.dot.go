// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// MakeLimitedReaderChan returns a new open channel
// (simply a 'chan *io.LimitedReader' that is).
//
// Note: No 'LimitedReader-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myLimitedReaderPipelineStartsHere := MakeLimitedReaderChan()
//	// ... lot's of code to design and build Your favourite "myLimitedReaderWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myLimitedReaderPipelineStartsHere <- drop
//	}
//	close(myLimitedReaderPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeLimitedReaderBuffer) the channel is unbuffered.
//
func MakeLimitedReaderChan() (out chan *io.LimitedReader) {
	return make(chan *io.LimitedReader)
}

func sendLimitedReader(out chan<- *io.LimitedReader, inp ...*io.LimitedReader) {
	defer close(out)
	for i := range inp {
		out <- inp[i]
	}
}

// ChanLimitedReader returns a channel to receive all inputs before close.
func ChanLimitedReader(inp ...*io.LimitedReader) (out <-chan *io.LimitedReader) {
	cha := make(chan *io.LimitedReader)
	go sendLimitedReader(cha, inp...)
	return cha
}

func sendLimitedReaderSlice(out chan<- *io.LimitedReader, inp ...[]*io.LimitedReader) {
	defer close(out)
	for i := range inp {
		for j := range inp[i] {
			out <- inp[i][j]
		}
	}
}

// ChanLimitedReaderSlice returns a channel to receive all inputs before close.
func ChanLimitedReaderSlice(inp ...[]*io.LimitedReader) (out <-chan *io.LimitedReader) {
	cha := make(chan *io.LimitedReader)
	go sendLimitedReaderSlice(cha, inp...)
	return cha
}

func chanLimitedReaderFuncNok(out chan<- *io.LimitedReader, act func() (*io.LimitedReader, bool)) {
	defer close(out)
	for {
		res, ok := act() // Apply action
		if !ok {
			return
		}
		out <- res
	}
}

// ChanLimitedReaderFuncNok returns a channel to receive all results of act until nok before close.
func ChanLimitedReaderFuncNok(act func() (*io.LimitedReader, bool)) (out <-chan *io.LimitedReader) {
	cha := make(chan *io.LimitedReader)
	go chanLimitedReaderFuncNok(cha, act)
	return cha
}

func chanLimitedReaderFuncErr(out chan<- *io.LimitedReader, act func() (*io.LimitedReader, error)) {
	defer close(out)
	for {
		res, err := act() // Apply action
		if err != nil {
			return
		}
		out <- res
	}
}

// ChanLimitedReaderFuncErr returns a channel to receive all results of act until err != nil before close.
func ChanLimitedReaderFuncErr(act func() (*io.LimitedReader, error)) (out <-chan *io.LimitedReader) {
	cha := make(chan *io.LimitedReader)
	go chanLimitedReaderFuncErr(cha, act)
	return cha
}

func joinLimitedReader(done chan<- struct{}, out chan<- *io.LimitedReader, inp ...*io.LimitedReader) {
	defer close(done)
	for i := range inp {
		out <- inp[i]
	}
	done <- struct{}{}
}

// JoinLimitedReader sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinLimitedReader(out chan<- *io.LimitedReader, inp ...*io.LimitedReader) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinLimitedReader(cha, out, inp...)
	return cha
}

func joinLimitedReaderSlice(done chan<- struct{}, out chan<- *io.LimitedReader, inp ...[]*io.LimitedReader) {
	defer close(done)
	for i := range inp {
		for j := range inp[i] {
			out <- inp[i][j]
		}
	}
	done <- struct{}{}
}

// JoinLimitedReaderSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinLimitedReaderSlice(out chan<- *io.LimitedReader, inp ...[]*io.LimitedReader) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinLimitedReaderSlice(cha, out, inp...)
	return cha
}

func joinLimitedReaderChan(done chan<- struct{}, out chan<- *io.LimitedReader, inp <-chan *io.LimitedReader) {
	defer close(done)
	for i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinLimitedReaderChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinLimitedReaderChan(out chan<- *io.LimitedReader, inp <-chan *io.LimitedReader) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinLimitedReaderChan(cha, out, inp)
	return cha
}

func doitLimitedReader(done chan<- struct{}, inp <-chan *io.LimitedReader) {
	defer close(done)
	for i := range inp {
		_ = i // Drain inp
	}
	done <- struct{}{}
}

// DoneLimitedReader returns a channel to receive one signal before close after inp has been drained.
func DoneLimitedReader(inp <-chan *io.LimitedReader) (done <-chan struct{}) {
	cha := make(chan struct{})
	go doitLimitedReader(cha, inp)
	return cha
}

func doitLimitedReaderSlice(done chan<- ([]*io.LimitedReader), inp <-chan *io.LimitedReader) {
	defer close(done)
	LimitedReaderS := []*io.LimitedReader{}
	for i := range inp {
		LimitedReaderS = append(LimitedReaderS, i)
	}
	done <- LimitedReaderS
}

// DoneLimitedReaderSlice returns a channel which will receive a slice
// of all the LimitedReaders received on inp channel before close.
// Unlike DoneLimitedReader, a full slice is sent once, not just an event.
func DoneLimitedReaderSlice(inp <-chan *io.LimitedReader) (done <-chan ([]*io.LimitedReader)) {
	cha := make(chan ([]*io.LimitedReader))
	go doitLimitedReaderSlice(cha, inp)
	return cha
}

func doitLimitedReaderFunc(done chan<- struct{}, inp <-chan *io.LimitedReader, act func(a *io.LimitedReader)) {
	defer close(done)
	for i := range inp {
		act(i) // Apply action
	}
	done <- struct{}{}
}

// DoneLimitedReaderFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneLimitedReaderFunc(inp <-chan *io.LimitedReader, act func(a *io.LimitedReader)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a *io.LimitedReader) { return }
	}
	go doitLimitedReaderFunc(cha, inp, act)
	return cha
}

func pipeLimitedReaderBuffer(out chan<- *io.LimitedReader, inp <-chan *io.LimitedReader) {
	defer close(out)
	for i := range inp {
		out <- i
	}
}

// PipeLimitedReaderBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeLimitedReaderBuffer(inp <-chan *io.LimitedReader, cap int) (out <-chan *io.LimitedReader) {
	cha := make(chan *io.LimitedReader, cap)
	go pipeLimitedReaderBuffer(cha, inp)
	return cha
}

func pipeLimitedReaderFunc(out chan<- *io.LimitedReader, inp <-chan *io.LimitedReader, act func(a *io.LimitedReader) *io.LimitedReader) {
	defer close(out)
	for i := range inp {
		out <- act(i)
	}
}

// PipeLimitedReaderFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeLimitedReaderMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeLimitedReaderFunc(inp <-chan *io.LimitedReader, act func(a *io.LimitedReader) *io.LimitedReader) (out <-chan *io.LimitedReader) {
	cha := make(chan *io.LimitedReader)
	if act == nil {
		act = func(a *io.LimitedReader) *io.LimitedReader { return a }
	}
	go pipeLimitedReaderFunc(cha, inp, act)
	return cha
}

func pipeLimitedReaderFork(out1, out2 chan<- *io.LimitedReader, inp <-chan *io.LimitedReader) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		out1 <- i
		out2 <- i
	}
}

// PipeLimitedReaderFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeLimitedReaderFork(inp <-chan *io.LimitedReader) (out1, out2 <-chan *io.LimitedReader) {
	cha1 := make(chan *io.LimitedReader)
	cha2 := make(chan *io.LimitedReader)
	go pipeLimitedReaderFork(cha1, cha2, inp)
	return cha1, cha2
}

// LimitedReaderTube is the signature for a pipe function.
type LimitedReaderTube func(inp <-chan *io.LimitedReader, out <-chan *io.LimitedReader)

// LimitedReaderDaisy returns a channel to receive all inp after having passed thru tube.
func LimitedReaderDaisy(inp <-chan *io.LimitedReader, tube LimitedReaderTube) (out <-chan *io.LimitedReader) {
	cha := make(chan *io.LimitedReader)
	go tube(inp, cha)
	return cha
}

// LimitedReaderDaisyChain returns a channel to receive all inp after having passed thru all tubes.
func LimitedReaderDaisyChain(inp <-chan *io.LimitedReader, tubes ...LimitedReaderTube) (out <-chan *io.LimitedReader) {
	cha := inp
	for i := range tubes {
		cha = LimitedReaderDaisy(cha, tubes[i])
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
