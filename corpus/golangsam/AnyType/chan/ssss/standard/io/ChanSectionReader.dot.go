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
func MakeSectionReaderChan() chan *io.SectionReader {
	return make(chan *io.SectionReader)
}

// ChanSectionReader returns a channel to receive all inputs before close.
func ChanSectionReader(inp ...*io.SectionReader) chan *io.SectionReader {
	out := make(chan *io.SectionReader)
	go func() {
		defer close(out)
		for i := range inp {
			out <- inp[i]
		}
	}()
	return out
}

// ChanSectionReaderSlice returns a channel to receive all inputs before close.
func ChanSectionReaderSlice(inp ...[]*io.SectionReader) chan *io.SectionReader {
	out := make(chan *io.SectionReader)
	go func() {
		defer close(out)
		for i := range inp {
			for j := range inp[i] {
				out <- inp[i][j]
			}
		}
	}()
	return out
}

// ChanSectionReaderFuncNok returns a channel to receive all results of act until nok before close.
func ChanSectionReaderFuncNok(act func() (*io.SectionReader, bool)) <-chan *io.SectionReader {
	out := make(chan *io.SectionReader)
	go func() {
		defer close(out)
		for {
			res, ok := act() // Apply action
			if !ok {
				return
			}
			out <- res
		}
	}()
	return out
}

// ChanSectionReaderFuncErr returns a channel to receive all results of act until err != nil before close.
func ChanSectionReaderFuncErr(act func() (*io.SectionReader, error)) <-chan *io.SectionReader {
	out := make(chan *io.SectionReader)
	go func() {
		defer close(out)
		for {
			res, err := act() // Apply action
			if err != nil {
				return
			}
			out <- res
		}
	}()
	return out
}

// JoinSectionReader sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinSectionReader(out chan<- *io.SectionReader, inp ...*io.SectionReader) chan struct{} {
	done := make(chan struct{})
	go func() {
		defer close(done)
		for i := range inp {
			out <- inp[i]
		}
		done <- struct{}{}
	}()
	return done
}

// JoinSectionReaderSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinSectionReaderSlice(out chan<- *io.SectionReader, inp ...[]*io.SectionReader) chan struct{} {
	done := make(chan struct{})
	go func() {
		defer close(done)
		for i := range inp {
			for j := range inp[i] {
				out <- inp[i][j]
			}
		}
		done <- struct{}{}
	}()
	return done
}

// JoinSectionReaderChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinSectionReaderChan(out chan<- *io.SectionReader, inp <-chan *io.SectionReader) chan struct{} {
	done := make(chan struct{})
	go func() {
		defer close(done)
		for i := range inp {
			out <- i
		}
		done <- struct{}{}
	}()
	return done
}

// DoneSectionReader returns a channel to receive one signal before close after inp has been drained.
func DoneSectionReader(inp <-chan *io.SectionReader) chan struct{} {
	done := make(chan struct{})
	go func() {
		defer close(done)
		for i := range inp {
			_ = i // Drain inp
		}
		done <- struct{}{}
	}()
	return done
}

// DoneSectionReaderSlice returns a channel which will receive a slice
// of all the SectionReaders received on inp channel before close.
// Unlike DoneSectionReader, a full slice is sent once, not just an event.
func DoneSectionReaderSlice(inp <-chan *io.SectionReader) chan []*io.SectionReader {
	done := make(chan []*io.SectionReader)
	go func() {
		defer close(done)
		SectionReaderS := []*io.SectionReader{}
		for i := range inp {
			SectionReaderS = append(SectionReaderS, i)
		}
		done <- SectionReaderS
	}()
	return done
}

// DoneSectionReaderFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneSectionReaderFunc(inp <-chan *io.SectionReader, act func(a *io.SectionReader)) chan struct{} {
	done := make(chan struct{})
	if act == nil {
		act = func(a *io.SectionReader) { return }
	}
	go func() {
		defer close(done)
		for i := range inp {
			act(i) // Apply action
		}
		done <- struct{}{}
	}()
	return done
}

// PipeSectionReaderBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeSectionReaderBuffer(inp <-chan *io.SectionReader, cap int) chan *io.SectionReader {
	out := make(chan *io.SectionReader, cap)
	go func() {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}()
	return out
}

// PipeSectionReaderFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeSectionReaderMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeSectionReaderFunc(inp <-chan *io.SectionReader, act func(a *io.SectionReader) *io.SectionReader) chan *io.SectionReader {
	out := make(chan *io.SectionReader)
	if act == nil {
		act = func(a *io.SectionReader) *io.SectionReader { return a }
	}
	go func() {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}()
	return out
}

// PipeSectionReaderFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeSectionReaderFork(inp <-chan *io.SectionReader) (chan *io.SectionReader, chan *io.SectionReader) {
	out1 := make(chan *io.SectionReader)
	out2 := make(chan *io.SectionReader)
	go func() {
		defer close(out1)
		defer close(out2)
		for i := range inp {
			out1 <- i
			out2 <- i
		}
	}()
	return out1, out2
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
