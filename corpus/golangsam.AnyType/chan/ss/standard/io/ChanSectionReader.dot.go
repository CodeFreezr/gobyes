// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// MakeSectionReaderChan returns a new open channel
// (simply a 'chan *io.SectionReader' that is).
//
// Note: No 'SectionReader-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var mySectionReaderPipelineStartsHere := MakeSectionReaderChan()
//	// ... lot's of code to design and build Your favourite "mySectionReaderWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		mySectionReaderPipelineStartsHere <- drop
//	}
//	close(mySectionReaderPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeSectionReaderBuffer) the channel is unbuffered.
//
func MakeSectionReaderChan() (out chan *io.SectionReader) {
	return make(chan *io.SectionReader)
}

func sendSectionReader(out chan<- *io.SectionReader, inp ...*io.SectionReader) {
	defer close(out)
	for i := range inp {
		out <- inp[i]
	}
}

// ChanSectionReader returns a channel to receive all inputs before close.
func ChanSectionReader(inp ...*io.SectionReader) (out <-chan *io.SectionReader) {
	cha := make(chan *io.SectionReader)
	go sendSectionReader(cha, inp...)
	return cha
}

func sendSectionReaderSlice(out chan<- *io.SectionReader, inp ...[]*io.SectionReader) {
	defer close(out)
	for i := range inp {
		for j := range inp[i] {
			out <- inp[i][j]
		}
	}
}

// ChanSectionReaderSlice returns a channel to receive all inputs before close.
func ChanSectionReaderSlice(inp ...[]*io.SectionReader) (out <-chan *io.SectionReader) {
	cha := make(chan *io.SectionReader)
	go sendSectionReaderSlice(cha, inp...)
	return cha
}

func chanSectionReaderFuncNok(out chan<- *io.SectionReader, act func() (*io.SectionReader, bool)) {
	defer close(out)
	for {
		res, ok := act() // Apply action
		if !ok {
			return
		}
		out <- res
	}
}

// ChanSectionReaderFuncNok returns a channel to receive all results of act until nok before close.
func ChanSectionReaderFuncNok(act func() (*io.SectionReader, bool)) (out <-chan *io.SectionReader) {
	cha := make(chan *io.SectionReader)
	go chanSectionReaderFuncNok(cha, act)
	return cha
}

func chanSectionReaderFuncErr(out chan<- *io.SectionReader, act func() (*io.SectionReader, error)) {
	defer close(out)
	for {
		res, err := act() // Apply action
		if err != nil {
			return
		}
		out <- res
	}
}

// ChanSectionReaderFuncErr returns a channel to receive all results of act until err != nil before close.
func ChanSectionReaderFuncErr(act func() (*io.SectionReader, error)) (out <-chan *io.SectionReader) {
	cha := make(chan *io.SectionReader)
	go chanSectionReaderFuncErr(cha, act)
	return cha
}

func joinSectionReader(done chan<- struct{}, out chan<- *io.SectionReader, inp ...*io.SectionReader) {
	defer close(done)
	for i := range inp {
		out <- inp[i]
	}
	done <- struct{}{}
}

// JoinSectionReader sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinSectionReader(out chan<- *io.SectionReader, inp ...*io.SectionReader) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinSectionReader(cha, out, inp...)
	return cha
}

func joinSectionReaderSlice(done chan<- struct{}, out chan<- *io.SectionReader, inp ...[]*io.SectionReader) {
	defer close(done)
	for i := range inp {
		for j := range inp[i] {
			out <- inp[i][j]
		}
	}
	done <- struct{}{}
}

// JoinSectionReaderSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinSectionReaderSlice(out chan<- *io.SectionReader, inp ...[]*io.SectionReader) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinSectionReaderSlice(cha, out, inp...)
	return cha
}

func joinSectionReaderChan(done chan<- struct{}, out chan<- *io.SectionReader, inp <-chan *io.SectionReader) {
	defer close(done)
	for i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinSectionReaderChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinSectionReaderChan(out chan<- *io.SectionReader, inp <-chan *io.SectionReader) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinSectionReaderChan(cha, out, inp)
	return cha
}

func doitSectionReader(done chan<- struct{}, inp <-chan *io.SectionReader) {
	defer close(done)
	for i := range inp {
		_ = i // Drain inp
	}
	done <- struct{}{}
}

// DoneSectionReader returns a channel to receive one signal before close after inp has been drained.
func DoneSectionReader(inp <-chan *io.SectionReader) (done <-chan struct{}) {
	cha := make(chan struct{})
	go doitSectionReader(cha, inp)
	return cha
}

func doitSectionReaderSlice(done chan<- ([]*io.SectionReader), inp <-chan *io.SectionReader) {
	defer close(done)
	SectionReaderS := []*io.SectionReader{}
	for i := range inp {
		SectionReaderS = append(SectionReaderS, i)
	}
	done <- SectionReaderS
}

// DoneSectionReaderSlice returns a channel which will receive a slice
// of all the SectionReaders received on inp channel before close.
// Unlike DoneSectionReader, a full slice is sent once, not just an event.
func DoneSectionReaderSlice(inp <-chan *io.SectionReader) (done <-chan ([]*io.SectionReader)) {
	cha := make(chan ([]*io.SectionReader))
	go doitSectionReaderSlice(cha, inp)
	return cha
}

func doitSectionReaderFunc(done chan<- struct{}, inp <-chan *io.SectionReader, act func(a *io.SectionReader)) {
	defer close(done)
	for i := range inp {
		act(i) // Apply action
	}
	done <- struct{}{}
}

// DoneSectionReaderFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneSectionReaderFunc(inp <-chan *io.SectionReader, act func(a *io.SectionReader)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a *io.SectionReader) { return }
	}
	go doitSectionReaderFunc(cha, inp, act)
	return cha
}

func pipeSectionReaderBuffer(out chan<- *io.SectionReader, inp <-chan *io.SectionReader) {
	defer close(out)
	for i := range inp {
		out <- i
	}
}

// PipeSectionReaderBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeSectionReaderBuffer(inp <-chan *io.SectionReader, cap int) (out <-chan *io.SectionReader) {
	cha := make(chan *io.SectionReader, cap)
	go pipeSectionReaderBuffer(cha, inp)
	return cha
}

func pipeSectionReaderFunc(out chan<- *io.SectionReader, inp <-chan *io.SectionReader, act func(a *io.SectionReader) *io.SectionReader) {
	defer close(out)
	for i := range inp {
		out <- act(i)
	}
}

// PipeSectionReaderFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeSectionReaderMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeSectionReaderFunc(inp <-chan *io.SectionReader, act func(a *io.SectionReader) *io.SectionReader) (out <-chan *io.SectionReader) {
	cha := make(chan *io.SectionReader)
	if act == nil {
		act = func(a *io.SectionReader) *io.SectionReader { return a }
	}
	go pipeSectionReaderFunc(cha, inp, act)
	return cha
}

func pipeSectionReaderFork(out1, out2 chan<- *io.SectionReader, inp <-chan *io.SectionReader) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		out1 <- i
		out2 <- i
	}
}

// PipeSectionReaderFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeSectionReaderFork(inp <-chan *io.SectionReader) (out1, out2 <-chan *io.SectionReader) {
	cha1 := make(chan *io.SectionReader)
	cha2 := make(chan *io.SectionReader)
	go pipeSectionReaderFork(cha1, cha2, inp)
	return cha1, cha2
}

// SectionReaderTube is the signature for a pipe function.
type SectionReaderTube func(inp <-chan *io.SectionReader, out <-chan *io.SectionReader)

// SectionReaderDaisy returns a channel to receive all inp after having passed thru tube.
func SectionReaderDaisy(inp <-chan *io.SectionReader, tube SectionReaderTube) (out <-chan *io.SectionReader) {
	cha := make(chan *io.SectionReader)
	go tube(inp, cha)
	return cha
}

// SectionReaderDaisyChain returns a channel to receive all inp after having passed thru all tubes.
func SectionReaderDaisyChain(inp <-chan *io.SectionReader, tubes ...SectionReaderTube) (out <-chan *io.SectionReader) {
	cha := inp
	for i := range tubes {
		cha = SectionReaderDaisy(cha, tubes[i])
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
