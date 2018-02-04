// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// MakeReaderChan returns a new open channel
// (simply a 'chan io.Reader' that is).
//
// Note: No 'Reader-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myReaderPipelineStartsHere := MakeReaderChan()
//	// ... lot's of code to design and build Your favourite "myReaderWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myReaderPipelineStartsHere <- drop
//	}
//	close(myReaderPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeReaderBuffer) the channel is unbuffered.
//
func MakeReaderChan() (out chan io.Reader) {
	return make(chan io.Reader)
}

func sendReader(out chan<- io.Reader, inp ...io.Reader) {
	defer close(out)
	for i := range inp {
		out <- inp[i]
	}
}

// ChanReader returns a channel to receive all inputs before close.
func ChanReader(inp ...io.Reader) (out <-chan io.Reader) {
	cha := make(chan io.Reader)
	go sendReader(cha, inp...)
	return cha
}

func sendReaderSlice(out chan<- io.Reader, inp ...[]io.Reader) {
	defer close(out)
	for i := range inp {
		for j := range inp[i] {
			out <- inp[i][j]
		}
	}
}

// ChanReaderSlice returns a channel to receive all inputs before close.
func ChanReaderSlice(inp ...[]io.Reader) (out <-chan io.Reader) {
	cha := make(chan io.Reader)
	go sendReaderSlice(cha, inp...)
	return cha
}

func chanReaderFuncNok(out chan<- io.Reader, act func() (io.Reader, bool)) {
	defer close(out)
	for {
		res, ok := act() // Apply action
		if !ok {
			return
		}
		out <- res
	}
}

// ChanReaderFuncNok returns a channel to receive all results of act until nok before close.
func ChanReaderFuncNok(act func() (io.Reader, bool)) (out <-chan io.Reader) {
	cha := make(chan io.Reader)
	go chanReaderFuncNok(cha, act)
	return cha
}

func chanReaderFuncErr(out chan<- io.Reader, act func() (io.Reader, error)) {
	defer close(out)
	for {
		res, err := act() // Apply action
		if err != nil {
			return
		}
		out <- res
	}
}

// ChanReaderFuncErr returns a channel to receive all results of act until err != nil before close.
func ChanReaderFuncErr(act func() (io.Reader, error)) (out <-chan io.Reader) {
	cha := make(chan io.Reader)
	go chanReaderFuncErr(cha, act)
	return cha
}

func joinReader(done chan<- struct{}, out chan<- io.Reader, inp ...io.Reader) {
	defer close(done)
	for i := range inp {
		out <- inp[i]
	}
	done <- struct{}{}
}

// JoinReader sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinReader(out chan<- io.Reader, inp ...io.Reader) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinReader(cha, out, inp...)
	return cha
}

func joinReaderSlice(done chan<- struct{}, out chan<- io.Reader, inp ...[]io.Reader) {
	defer close(done)
	for i := range inp {
		for j := range inp[i] {
			out <- inp[i][j]
		}
	}
	done <- struct{}{}
}

// JoinReaderSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinReaderSlice(out chan<- io.Reader, inp ...[]io.Reader) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinReaderSlice(cha, out, inp...)
	return cha
}

func joinReaderChan(done chan<- struct{}, out chan<- io.Reader, inp <-chan io.Reader) {
	defer close(done)
	for i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinReaderChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinReaderChan(out chan<- io.Reader, inp <-chan io.Reader) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinReaderChan(cha, out, inp)
	return cha
}

func doitReader(done chan<- struct{}, inp <-chan io.Reader) {
	defer close(done)
	for i := range inp {
		_ = i // Drain inp
	}
	done <- struct{}{}
}

// DoneReader returns a channel to receive one signal before close after inp has been drained.
func DoneReader(inp <-chan io.Reader) (done <-chan struct{}) {
	cha := make(chan struct{})
	go doitReader(cha, inp)
	return cha
}

func doitReaderSlice(done chan<- ([]io.Reader), inp <-chan io.Reader) {
	defer close(done)
	ReaderS := []io.Reader{}
	for i := range inp {
		ReaderS = append(ReaderS, i)
	}
	done <- ReaderS
}

// DoneReaderSlice returns a channel which will receive a slice
// of all the Readers received on inp channel before close.
// Unlike DoneReader, a full slice is sent once, not just an event.
func DoneReaderSlice(inp <-chan io.Reader) (done <-chan ([]io.Reader)) {
	cha := make(chan ([]io.Reader))
	go doitReaderSlice(cha, inp)
	return cha
}

func doitReaderFunc(done chan<- struct{}, inp <-chan io.Reader, act func(a io.Reader)) {
	defer close(done)
	for i := range inp {
		act(i) // Apply action
	}
	done <- struct{}{}
}

// DoneReaderFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneReaderFunc(inp <-chan io.Reader, act func(a io.Reader)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a io.Reader) { return }
	}
	go doitReaderFunc(cha, inp, act)
	return cha
}

func pipeReaderBuffer(out chan<- io.Reader, inp <-chan io.Reader) {
	defer close(out)
	for i := range inp {
		out <- i
	}
}

// PipeReaderBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeReaderBuffer(inp <-chan io.Reader, cap int) (out <-chan io.Reader) {
	cha := make(chan io.Reader, cap)
	go pipeReaderBuffer(cha, inp)
	return cha
}

func pipeReaderFunc(out chan<- io.Reader, inp <-chan io.Reader, act func(a io.Reader) io.Reader) {
	defer close(out)
	for i := range inp {
		out <- act(i)
	}
}

// PipeReaderFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeReaderMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeReaderFunc(inp <-chan io.Reader, act func(a io.Reader) io.Reader) (out <-chan io.Reader) {
	cha := make(chan io.Reader)
	if act == nil {
		act = func(a io.Reader) io.Reader { return a }
	}
	go pipeReaderFunc(cha, inp, act)
	return cha
}

func pipeReaderFork(out1, out2 chan<- io.Reader, inp <-chan io.Reader) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		out1 <- i
		out2 <- i
	}
}

// PipeReaderFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeReaderFork(inp <-chan io.Reader) (out1, out2 <-chan io.Reader) {
	cha1 := make(chan io.Reader)
	cha2 := make(chan io.Reader)
	go pipeReaderFork(cha1, cha2, inp)
	return cha1, cha2
}

// ReaderTube is the signature for a pipe function.
type ReaderTube func(inp <-chan io.Reader, out <-chan io.Reader)

// ReaderDaisy returns a channel to receive all inp after having passed thru tube.
func ReaderDaisy(inp <-chan io.Reader, tube ReaderTube) (out <-chan io.Reader) {
	cha := make(chan io.Reader)
	go tube(inp, cha)
	return cha
}

// ReaderDaisyChain returns a channel to receive all inp after having passed thru all tubes.
func ReaderDaisyChain(inp <-chan io.Reader, tubes ...ReaderTube) (out <-chan io.Reader) {
	cha := inp
	for i := range tubes {
		cha = ReaderDaisy(cha, tubes[i])
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
