// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// MakePipeWriterChan returns a new open channel
// (simply a 'chan *io.PipeWriter' that is).
//
// Note: No 'PipeWriter-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myPipeWriterPipelineStartsHere := MakePipeWriterChan()
//	// ... lot's of code to design and build Your favourite "myPipeWriterWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myPipeWriterPipelineStartsHere <- drop
//	}
//	close(myPipeWriterPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipePipeWriterBuffer) the channel is unbuffered.
//
func MakePipeWriterChan() (out chan *io.PipeWriter) {
	return make(chan *io.PipeWriter)
}

func sendPipeWriter(out chan<- *io.PipeWriter, inp ...*io.PipeWriter) {
	defer close(out)
	for i := range inp {
		out <- inp[i]
	}
}

// ChanPipeWriter returns a channel to receive all inputs before close.
func ChanPipeWriter(inp ...*io.PipeWriter) (out <-chan *io.PipeWriter) {
	cha := make(chan *io.PipeWriter)
	go sendPipeWriter(cha, inp...)
	return cha
}

func sendPipeWriterSlice(out chan<- *io.PipeWriter, inp ...[]*io.PipeWriter) {
	defer close(out)
	for i := range inp {
		for j := range inp[i] {
			out <- inp[i][j]
		}
	}
}

// ChanPipeWriterSlice returns a channel to receive all inputs before close.
func ChanPipeWriterSlice(inp ...[]*io.PipeWriter) (out <-chan *io.PipeWriter) {
	cha := make(chan *io.PipeWriter)
	go sendPipeWriterSlice(cha, inp...)
	return cha
}

func chanPipeWriterFuncNok(out chan<- *io.PipeWriter, act func() (*io.PipeWriter, bool)) {
	defer close(out)
	for {
		res, ok := act() // Apply action
		if !ok {
			return
		}
		out <- res
	}
}

// ChanPipeWriterFuncNok returns a channel to receive all results of act until nok before close.
func ChanPipeWriterFuncNok(act func() (*io.PipeWriter, bool)) (out <-chan *io.PipeWriter) {
	cha := make(chan *io.PipeWriter)
	go chanPipeWriterFuncNok(cha, act)
	return cha
}

func chanPipeWriterFuncErr(out chan<- *io.PipeWriter, act func() (*io.PipeWriter, error)) {
	defer close(out)
	for {
		res, err := act() // Apply action
		if err != nil {
			return
		}
		out <- res
	}
}

// ChanPipeWriterFuncErr returns a channel to receive all results of act until err != nil before close.
func ChanPipeWriterFuncErr(act func() (*io.PipeWriter, error)) (out <-chan *io.PipeWriter) {
	cha := make(chan *io.PipeWriter)
	go chanPipeWriterFuncErr(cha, act)
	return cha
}

func joinPipeWriter(done chan<- struct{}, out chan<- *io.PipeWriter, inp ...*io.PipeWriter) {
	defer close(done)
	for i := range inp {
		out <- inp[i]
	}
	done <- struct{}{}
}

// JoinPipeWriter sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinPipeWriter(out chan<- *io.PipeWriter, inp ...*io.PipeWriter) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinPipeWriter(cha, out, inp...)
	return cha
}

func joinPipeWriterSlice(done chan<- struct{}, out chan<- *io.PipeWriter, inp ...[]*io.PipeWriter) {
	defer close(done)
	for i := range inp {
		for j := range inp[i] {
			out <- inp[i][j]
		}
	}
	done <- struct{}{}
}

// JoinPipeWriterSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinPipeWriterSlice(out chan<- *io.PipeWriter, inp ...[]*io.PipeWriter) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinPipeWriterSlice(cha, out, inp...)
	return cha
}

func joinPipeWriterChan(done chan<- struct{}, out chan<- *io.PipeWriter, inp <-chan *io.PipeWriter) {
	defer close(done)
	for i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinPipeWriterChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinPipeWriterChan(out chan<- *io.PipeWriter, inp <-chan *io.PipeWriter) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinPipeWriterChan(cha, out, inp)
	return cha
}

func doitPipeWriter(done chan<- struct{}, inp <-chan *io.PipeWriter) {
	defer close(done)
	for i := range inp {
		_ = i // Drain inp
	}
	done <- struct{}{}
}

// DonePipeWriter returns a channel to receive one signal before close after inp has been drained.
func DonePipeWriter(inp <-chan *io.PipeWriter) (done <-chan struct{}) {
	cha := make(chan struct{})
	go doitPipeWriter(cha, inp)
	return cha
}

func doitPipeWriterSlice(done chan<- ([]*io.PipeWriter), inp <-chan *io.PipeWriter) {
	defer close(done)
	PipeWriterS := []*io.PipeWriter{}
	for i := range inp {
		PipeWriterS = append(PipeWriterS, i)
	}
	done <- PipeWriterS
}

// DonePipeWriterSlice returns a channel which will receive a slice
// of all the PipeWriters received on inp channel before close.
// Unlike DonePipeWriter, a full slice is sent once, not just an event.
func DonePipeWriterSlice(inp <-chan *io.PipeWriter) (done <-chan ([]*io.PipeWriter)) {
	cha := make(chan ([]*io.PipeWriter))
	go doitPipeWriterSlice(cha, inp)
	return cha
}

func doitPipeWriterFunc(done chan<- struct{}, inp <-chan *io.PipeWriter, act func(a *io.PipeWriter)) {
	defer close(done)
	for i := range inp {
		act(i) // Apply action
	}
	done <- struct{}{}
}

// DonePipeWriterFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DonePipeWriterFunc(inp <-chan *io.PipeWriter, act func(a *io.PipeWriter)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a *io.PipeWriter) { return }
	}
	go doitPipeWriterFunc(cha, inp, act)
	return cha
}

func pipePipeWriterBuffer(out chan<- *io.PipeWriter, inp <-chan *io.PipeWriter) {
	defer close(out)
	for i := range inp {
		out <- i
	}
}

// PipePipeWriterBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipePipeWriterBuffer(inp <-chan *io.PipeWriter, cap int) (out <-chan *io.PipeWriter) {
	cha := make(chan *io.PipeWriter, cap)
	go pipePipeWriterBuffer(cha, inp)
	return cha
}

func pipePipeWriterFunc(out chan<- *io.PipeWriter, inp <-chan *io.PipeWriter, act func(a *io.PipeWriter) *io.PipeWriter) {
	defer close(out)
	for i := range inp {
		out <- act(i)
	}
}

// PipePipeWriterFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipePipeWriterMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipePipeWriterFunc(inp <-chan *io.PipeWriter, act func(a *io.PipeWriter) *io.PipeWriter) (out <-chan *io.PipeWriter) {
	cha := make(chan *io.PipeWriter)
	if act == nil {
		act = func(a *io.PipeWriter) *io.PipeWriter { return a }
	}
	go pipePipeWriterFunc(cha, inp, act)
	return cha
}

func pipePipeWriterFork(out1, out2 chan<- *io.PipeWriter, inp <-chan *io.PipeWriter) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		out1 <- i
		out2 <- i
	}
}

// PipePipeWriterFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipePipeWriterFork(inp <-chan *io.PipeWriter) (out1, out2 <-chan *io.PipeWriter) {
	cha1 := make(chan *io.PipeWriter)
	cha2 := make(chan *io.PipeWriter)
	go pipePipeWriterFork(cha1, cha2, inp)
	return cha1, cha2
}

// PipeWriterTube is the signature for a pipe function.
type PipeWriterTube func(inp <-chan *io.PipeWriter, out <-chan *io.PipeWriter)

// PipeWriterDaisy returns a channel to receive all inp after having passed thru tube.
func PipeWriterDaisy(inp <-chan *io.PipeWriter, tube PipeWriterTube) (out <-chan *io.PipeWriter) {
	cha := make(chan *io.PipeWriter)
	go tube(inp, cha)
	return cha
}

// PipeWriterDaisyChain returns a channel to receive all inp after having passed thru all tubes.
func PipeWriterDaisyChain(inp <-chan *io.PipeWriter, tubes ...PipeWriterTube) (out <-chan *io.PipeWriter) {
	cha := inp
	for i := range tubes {
		cha = PipeWriterDaisy(cha, tubes[i])
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
