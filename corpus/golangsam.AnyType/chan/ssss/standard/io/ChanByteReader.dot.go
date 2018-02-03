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
func MakeByteReaderChan() chan io.ByteReader {
	return make(chan io.ByteReader)
}

// ChanByteReader returns a channel to receive all inputs before close.
func ChanByteReader(inp ...io.ByteReader) chan io.ByteReader {
	out := make(chan io.ByteReader)
	go func() {
		defer close(out)
		for i := range inp {
			out <- inp[i]
		}
	}()
	return out
}

// ChanByteReaderSlice returns a channel to receive all inputs before close.
func ChanByteReaderSlice(inp ...[]io.ByteReader) chan io.ByteReader {
	out := make(chan io.ByteReader)
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

// ChanByteReaderFuncNok returns a channel to receive all results of act until nok before close.
func ChanByteReaderFuncNok(act func() (io.ByteReader, bool)) <-chan io.ByteReader {
	out := make(chan io.ByteReader)
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

// ChanByteReaderFuncErr returns a channel to receive all results of act until err != nil before close.
func ChanByteReaderFuncErr(act func() (io.ByteReader, error)) <-chan io.ByteReader {
	out := make(chan io.ByteReader)
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

// JoinByteReader sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinByteReader(out chan<- io.ByteReader, inp ...io.ByteReader) chan struct{} {
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

// JoinByteReaderSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinByteReaderSlice(out chan<- io.ByteReader, inp ...[]io.ByteReader) chan struct{} {
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

// JoinByteReaderChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinByteReaderChan(out chan<- io.ByteReader, inp <-chan io.ByteReader) chan struct{} {
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

// DoneByteReader returns a channel to receive one signal before close after inp has been drained.
func DoneByteReader(inp <-chan io.ByteReader) chan struct{} {
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

// DoneByteReaderSlice returns a channel which will receive a slice
// of all the ByteReaders received on inp channel before close.
// Unlike DoneByteReader, a full slice is sent once, not just an event.
func DoneByteReaderSlice(inp <-chan io.ByteReader) chan []io.ByteReader {
	done := make(chan []io.ByteReader)
	go func() {
		defer close(done)
		ByteReaderS := []io.ByteReader{}
		for i := range inp {
			ByteReaderS = append(ByteReaderS, i)
		}
		done <- ByteReaderS
	}()
	return done
}

// DoneByteReaderFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneByteReaderFunc(inp <-chan io.ByteReader, act func(a io.ByteReader)) chan struct{} {
	done := make(chan struct{})
	if act == nil {
		act = func(a io.ByteReader) { return }
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

// PipeByteReaderBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeByteReaderBuffer(inp <-chan io.ByteReader, cap int) chan io.ByteReader {
	out := make(chan io.ByteReader, cap)
	go func() {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}()
	return out
}

// PipeByteReaderFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeByteReaderMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeByteReaderFunc(inp <-chan io.ByteReader, act func(a io.ByteReader) io.ByteReader) chan io.ByteReader {
	out := make(chan io.ByteReader)
	if act == nil {
		act = func(a io.ByteReader) io.ByteReader { return a }
	}
	go func() {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}()
	return out
}

// PipeByteReaderFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeByteReaderFork(inp <-chan io.ByteReader) (chan io.ByteReader, chan io.ByteReader) {
	out1 := make(chan io.ByteReader)
	out2 := make(chan io.ByteReader)
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
