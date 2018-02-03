// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// MakeByteReaderChan returns a new open channel
// (simply a 'chan io.ByteReader' that is).
//
// Note: No 'ByteReader-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myByteReaderPipelineStartsHere := MakeByteReaderChan()
//	// ... lot's of code to design and build Your favourite "myByteReaderWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myByteReaderPipelineStartsHere <- drop
//	}
//	close(myByteReaderPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeByteReaderBuffer) the channel is unbuffered.
//
func MakeByteReaderChan() (out chan io.ByteReader) {
	return make(chan io.ByteReader)
}

func sendByteReader(out chan<- io.ByteReader, inp ...io.ByteReader) {
	defer close(out)
	for i := range inp {
		out <- inp[i]
	}
}

// ChanByteReader returns a channel to receive all inputs before close.
func ChanByteReader(inp ...io.ByteReader) (out <-chan io.ByteReader) {
	cha := make(chan io.ByteReader)
	go sendByteReader(cha, inp...)
	return cha
}

func sendByteReaderSlice(out chan<- io.ByteReader, inp ...[]io.ByteReader) {
	defer close(out)
	for i := range inp {
		for j := range inp[i] {
			out <- inp[i][j]
		}
	}
}

// ChanByteReaderSlice returns a channel to receive all inputs before close.
func ChanByteReaderSlice(inp ...[]io.ByteReader) (out <-chan io.ByteReader) {
	cha := make(chan io.ByteReader)
	go sendByteReaderSlice(cha, inp...)
	return cha
}

func chanByteReaderFuncNok(out chan<- io.ByteReader, act func() (io.ByteReader, bool)) {
	defer close(out)
	for {
		res, ok := act() // Apply action
		if !ok {
			return
		}
		out <- res
	}
}

// ChanByteReaderFuncNok returns a channel to receive all results of act until nok before close.
func ChanByteReaderFuncNok(act func() (io.ByteReader, bool)) (out <-chan io.ByteReader) {
	cha := make(chan io.ByteReader)
	go chanByteReaderFuncNok(cha, act)
	return cha
}

func chanByteReaderFuncErr(out chan<- io.ByteReader, act func() (io.ByteReader, error)) {
	defer close(out)
	for {
		res, err := act() // Apply action
		if err != nil {
			return
		}
		out <- res
	}
}

// ChanByteReaderFuncErr returns a channel to receive all results of act until err != nil before close.
func ChanByteReaderFuncErr(act func() (io.ByteReader, error)) (out <-chan io.ByteReader) {
	cha := make(chan io.ByteReader)
	go chanByteReaderFuncErr(cha, act)
	return cha
}

func joinByteReader(done chan<- struct{}, out chan<- io.ByteReader, inp ...io.ByteReader) {
	defer close(done)
	for i := range inp {
		out <- inp[i]
	}
	done <- struct{}{}
}

// JoinByteReader sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinByteReader(out chan<- io.ByteReader, inp ...io.ByteReader) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinByteReader(cha, out, inp...)
	return cha
}

func joinByteReaderSlice(done chan<- struct{}, out chan<- io.ByteReader, inp ...[]io.ByteReader) {
	defer close(done)
	for i := range inp {
		for j := range inp[i] {
			out <- inp[i][j]
		}
	}
	done <- struct{}{}
}

// JoinByteReaderSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinByteReaderSlice(out chan<- io.ByteReader, inp ...[]io.ByteReader) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinByteReaderSlice(cha, out, inp...)
	return cha
}

func joinByteReaderChan(done chan<- struct{}, out chan<- io.ByteReader, inp <-chan io.ByteReader) {
	defer close(done)
	for i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinByteReaderChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinByteReaderChan(out chan<- io.ByteReader, inp <-chan io.ByteReader) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinByteReaderChan(cha, out, inp)
	return cha
}

func doitByteReader(done chan<- struct{}, inp <-chan io.ByteReader) {
	defer close(done)
	for i := range inp {
		_ = i // Drain inp
	}
	done <- struct{}{}
}

// DoneByteReader returns a channel to receive one signal before close after inp has been drained.
func DoneByteReader(inp <-chan io.ByteReader) (done <-chan struct{}) {
	cha := make(chan struct{})
	go doitByteReader(cha, inp)
	return cha
}

func doitByteReaderSlice(done chan<- ([]io.ByteReader), inp <-chan io.ByteReader) {
	defer close(done)
	ByteReaderS := []io.ByteReader{}
	for i := range inp {
		ByteReaderS = append(ByteReaderS, i)
	}
	done <- ByteReaderS
}

// DoneByteReaderSlice returns a channel which will receive a slice
// of all the ByteReaders received on inp channel before close.
// Unlike DoneByteReader, a full slice is sent once, not just an event.
func DoneByteReaderSlice(inp <-chan io.ByteReader) (done <-chan ([]io.ByteReader)) {
	cha := make(chan ([]io.ByteReader))
	go doitByteReaderSlice(cha, inp)
	return cha
}

func doitByteReaderFunc(done chan<- struct{}, inp <-chan io.ByteReader, act func(a io.ByteReader)) {
	defer close(done)
	for i := range inp {
		act(i) // Apply action
	}
	done <- struct{}{}
}

// DoneByteReaderFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneByteReaderFunc(inp <-chan io.ByteReader, act func(a io.ByteReader)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a io.ByteReader) { return }
	}
	go doitByteReaderFunc(cha, inp, act)
	return cha
}

func pipeByteReaderBuffer(out chan<- io.ByteReader, inp <-chan io.ByteReader) {
	defer close(out)
	for i := range inp {
		out <- i
	}
}

// PipeByteReaderBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeByteReaderBuffer(inp <-chan io.ByteReader, cap int) (out <-chan io.ByteReader) {
	cha := make(chan io.ByteReader, cap)
	go pipeByteReaderBuffer(cha, inp)
	return cha
}

func pipeByteReaderFunc(out chan<- io.ByteReader, inp <-chan io.ByteReader, act func(a io.ByteReader) io.ByteReader) {
	defer close(out)
	for i := range inp {
		out <- act(i)
	}
}

// PipeByteReaderFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeByteReaderMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeByteReaderFunc(inp <-chan io.ByteReader, act func(a io.ByteReader) io.ByteReader) (out <-chan io.ByteReader) {
	cha := make(chan io.ByteReader)
	if act == nil {
		act = func(a io.ByteReader) io.ByteReader { return a }
	}
	go pipeByteReaderFunc(cha, inp, act)
	return cha
}

func pipeByteReaderFork(out1, out2 chan<- io.ByteReader, inp <-chan io.ByteReader) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		out1 <- i
		out2 <- i
	}
}

// PipeByteReaderFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeByteReaderFork(inp <-chan io.ByteReader) (out1, out2 <-chan io.ByteReader) {
	cha1 := make(chan io.ByteReader)
	cha2 := make(chan io.ByteReader)
	go pipeByteReaderFork(cha1, cha2, inp)
	return cha1, cha2
}

// ByteReaderTube is the signature for a pipe function.
type ByteReaderTube func(inp <-chan io.ByteReader, out <-chan io.ByteReader)

// ByteReaderDaisy returns a channel to receive all inp after having passed thru tube.
func ByteReaderDaisy(inp <-chan io.ByteReader, tube ByteReaderTube) (out <-chan io.ByteReader) {
	cha := make(chan io.ByteReader)
	go tube(inp, cha)
	return cha
}

// ByteReaderDaisyChain returns a channel to receive all inp after having passed thru all tubes.
func ByteReaderDaisyChain(inp <-chan io.ByteReader, tubes ...ByteReaderTube) (out <-chan io.ByteReader) {
	cha := inp
	for i := range tubes {
		cha = ByteReaderDaisy(cha, tubes[i])
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
