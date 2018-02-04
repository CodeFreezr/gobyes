// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// MakeReadWriterChan returns a new open channel
// (simply a 'chan io.ReadWriter' that is).
//
// Note: No 'ReadWriter-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myReadWriterPipelineStartsHere := MakeReadWriterChan()
//	// ... lot's of code to design and build Your favourite "myReadWriterWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myReadWriterPipelineStartsHere <- drop
//	}
//	close(myReadWriterPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeReadWriterBuffer) the channel is unbuffered.
//
func MakeReadWriterChan() (out chan io.ReadWriter) {
	return make(chan io.ReadWriter)
}

func sendReadWriter(out chan<- io.ReadWriter, inp ...io.ReadWriter) {
	defer close(out)
	for i := range inp {
		out <- inp[i]
	}
}

// ChanReadWriter returns a channel to receive all inputs before close.
func ChanReadWriter(inp ...io.ReadWriter) (out <-chan io.ReadWriter) {
	cha := make(chan io.ReadWriter)
	go sendReadWriter(cha, inp...)
	return cha
}

func sendReadWriterSlice(out chan<- io.ReadWriter, inp ...[]io.ReadWriter) {
	defer close(out)
	for i := range inp {
		for j := range inp[i] {
			out <- inp[i][j]
		}
	}
}

// ChanReadWriterSlice returns a channel to receive all inputs before close.
func ChanReadWriterSlice(inp ...[]io.ReadWriter) (out <-chan io.ReadWriter) {
	cha := make(chan io.ReadWriter)
	go sendReadWriterSlice(cha, inp...)
	return cha
}

func chanReadWriterFuncNok(out chan<- io.ReadWriter, act func() (io.ReadWriter, bool)) {
	defer close(out)
	for {
		res, ok := act() // Apply action
		if !ok {
			return
		}
		out <- res
	}
}

// ChanReadWriterFuncNok returns a channel to receive all results of act until nok before close.
func ChanReadWriterFuncNok(act func() (io.ReadWriter, bool)) (out <-chan io.ReadWriter) {
	cha := make(chan io.ReadWriter)
	go chanReadWriterFuncNok(cha, act)
	return cha
}

func chanReadWriterFuncErr(out chan<- io.ReadWriter, act func() (io.ReadWriter, error)) {
	defer close(out)
	for {
		res, err := act() // Apply action
		if err != nil {
			return
		}
		out <- res
	}
}

// ChanReadWriterFuncErr returns a channel to receive all results of act until err != nil before close.
func ChanReadWriterFuncErr(act func() (io.ReadWriter, error)) (out <-chan io.ReadWriter) {
	cha := make(chan io.ReadWriter)
	go chanReadWriterFuncErr(cha, act)
	return cha
}

func joinReadWriter(done chan<- struct{}, out chan<- io.ReadWriter, inp ...io.ReadWriter) {
	defer close(done)
	for i := range inp {
		out <- inp[i]
	}
	done <- struct{}{}
}

// JoinReadWriter sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinReadWriter(out chan<- io.ReadWriter, inp ...io.ReadWriter) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinReadWriter(cha, out, inp...)
	return cha
}

func joinReadWriterSlice(done chan<- struct{}, out chan<- io.ReadWriter, inp ...[]io.ReadWriter) {
	defer close(done)
	for i := range inp {
		for j := range inp[i] {
			out <- inp[i][j]
		}
	}
	done <- struct{}{}
}

// JoinReadWriterSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinReadWriterSlice(out chan<- io.ReadWriter, inp ...[]io.ReadWriter) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinReadWriterSlice(cha, out, inp...)
	return cha
}

func joinReadWriterChan(done chan<- struct{}, out chan<- io.ReadWriter, inp <-chan io.ReadWriter) {
	defer close(done)
	for i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinReadWriterChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinReadWriterChan(out chan<- io.ReadWriter, inp <-chan io.ReadWriter) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinReadWriterChan(cha, out, inp)
	return cha
}

func doitReadWriter(done chan<- struct{}, inp <-chan io.ReadWriter) {
	defer close(done)
	for i := range inp {
		_ = i // Drain inp
	}
	done <- struct{}{}
}

// DoneReadWriter returns a channel to receive one signal before close after inp has been drained.
func DoneReadWriter(inp <-chan io.ReadWriter) (done <-chan struct{}) {
	cha := make(chan struct{})
	go doitReadWriter(cha, inp)
	return cha
}

func doitReadWriterSlice(done chan<- ([]io.ReadWriter), inp <-chan io.ReadWriter) {
	defer close(done)
	ReadWriterS := []io.ReadWriter{}
	for i := range inp {
		ReadWriterS = append(ReadWriterS, i)
	}
	done <- ReadWriterS
}

// DoneReadWriterSlice returns a channel which will receive a slice
// of all the ReadWriters received on inp channel before close.
// Unlike DoneReadWriter, a full slice is sent once, not just an event.
func DoneReadWriterSlice(inp <-chan io.ReadWriter) (done <-chan ([]io.ReadWriter)) {
	cha := make(chan ([]io.ReadWriter))
	go doitReadWriterSlice(cha, inp)
	return cha
}

func doitReadWriterFunc(done chan<- struct{}, inp <-chan io.ReadWriter, act func(a io.ReadWriter)) {
	defer close(done)
	for i := range inp {
		act(i) // Apply action
	}
	done <- struct{}{}
}

// DoneReadWriterFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneReadWriterFunc(inp <-chan io.ReadWriter, act func(a io.ReadWriter)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a io.ReadWriter) { return }
	}
	go doitReadWriterFunc(cha, inp, act)
	return cha
}

func pipeReadWriterBuffer(out chan<- io.ReadWriter, inp <-chan io.ReadWriter) {
	defer close(out)
	for i := range inp {
		out <- i
	}
}

// PipeReadWriterBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeReadWriterBuffer(inp <-chan io.ReadWriter, cap int) (out <-chan io.ReadWriter) {
	cha := make(chan io.ReadWriter, cap)
	go pipeReadWriterBuffer(cha, inp)
	return cha
}

func pipeReadWriterFunc(out chan<- io.ReadWriter, inp <-chan io.ReadWriter, act func(a io.ReadWriter) io.ReadWriter) {
	defer close(out)
	for i := range inp {
		out <- act(i)
	}
}

// PipeReadWriterFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeReadWriterMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeReadWriterFunc(inp <-chan io.ReadWriter, act func(a io.ReadWriter) io.ReadWriter) (out <-chan io.ReadWriter) {
	cha := make(chan io.ReadWriter)
	if act == nil {
		act = func(a io.ReadWriter) io.ReadWriter { return a }
	}
	go pipeReadWriterFunc(cha, inp, act)
	return cha
}

func pipeReadWriterFork(out1, out2 chan<- io.ReadWriter, inp <-chan io.ReadWriter) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		out1 <- i
		out2 <- i
	}
}

// PipeReadWriterFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeReadWriterFork(inp <-chan io.ReadWriter) (out1, out2 <-chan io.ReadWriter) {
	cha1 := make(chan io.ReadWriter)
	cha2 := make(chan io.ReadWriter)
	go pipeReadWriterFork(cha1, cha2, inp)
	return cha1, cha2
}

// ReadWriterTube is the signature for a pipe function.
type ReadWriterTube func(inp <-chan io.ReadWriter, out <-chan io.ReadWriter)

// ReadWriterDaisy returns a channel to receive all inp after having passed thru tube.
func ReadWriterDaisy(inp <-chan io.ReadWriter, tube ReadWriterTube) (out <-chan io.ReadWriter) {
	cha := make(chan io.ReadWriter)
	go tube(inp, cha)
	return cha
}

// ReadWriterDaisyChain returns a channel to receive all inp after having passed thru all tubes.
func ReadWriterDaisyChain(inp <-chan io.ReadWriter, tubes ...ReadWriterTube) (out <-chan io.ReadWriter) {
	cha := inp
	for i := range tubes {
		cha = ReadWriterDaisy(cha, tubes[i])
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
