// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsByte

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"io"
)

// MakeByteWriterChan returns a new open channel
// (simply a 'chan io.ByteWriter' that is).
//
// Note: No 'ByteWriter-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myByteWriterPipelineStartsHere := MakeByteWriterChan()
//	// ... lot's of code to design and build Your favourite "myByteWriterWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myByteWriterPipelineStartsHere <- drop
//	}
//	close(myByteWriterPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeByteWriterBuffer) the channel is unbuffered.
//
func MakeByteWriterChan() (out chan io.ByteWriter) {
	return make(chan io.ByteWriter)
}

// ChanByteWriter returns a channel to receive all inputs before close.
func ChanByteWriter(inp ...io.ByteWriter) (out <-chan io.ByteWriter) {
	cha := make(chan io.ByteWriter)
	go func(out chan<- io.ByteWriter, inp ...io.ByteWriter) {
		defer close(out)
		for i := range inp {
			out <- inp[i]
		}
	}(cha, inp...)
	return cha
}

// ChanByteWriterSlice returns a channel to receive all inputs before close.
func ChanByteWriterSlice(inp ...[]io.ByteWriter) (out <-chan io.ByteWriter) {
	cha := make(chan io.ByteWriter)
	go func(out chan<- io.ByteWriter, inp ...[]io.ByteWriter) {
		defer close(out)
		for i := range inp {
			for j := range inp[i] {
				out <- inp[i][j]
			}
		}
	}(cha, inp...)
	return cha
}

// ChanByteWriterFuncNok returns a channel to receive all results of act until nok before close.
func ChanByteWriterFuncNok(act func() (io.ByteWriter, bool)) (out <-chan io.ByteWriter) {
	cha := make(chan io.ByteWriter)
	go func(out chan<- io.ByteWriter, act func() (io.ByteWriter, bool)) {
		defer close(out)
		for {
			res, ok := act() // Apply action
			if !ok {
				return
			}
			out <- res
		}
	}(cha, act)
	return cha
}

// ChanByteWriterFuncErr returns a channel to receive all results of act until err != nil before close.
func ChanByteWriterFuncErr(act func() (io.ByteWriter, error)) (out <-chan io.ByteWriter) {
	cha := make(chan io.ByteWriter)
	go func(out chan<- io.ByteWriter, act func() (io.ByteWriter, error)) {
		defer close(out)
		for {
			res, err := act() // Apply action
			if err != nil {
				return
			}
			out <- res
		}
	}(cha, act)
	return cha
}

// JoinByteWriter sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinByteWriter(out chan<- io.ByteWriter, inp ...io.ByteWriter) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- io.ByteWriter, inp ...io.ByteWriter) {
		defer close(done)
		for i := range inp {
			out <- inp[i]
		}
		done <- struct{}{}
	}(cha, out, inp...)
	return cha
}

// JoinByteWriterSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinByteWriterSlice(out chan<- io.ByteWriter, inp ...[]io.ByteWriter) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- io.ByteWriter, inp ...[]io.ByteWriter) {
		defer close(done)
		for i := range inp {
			for j := range inp[i] {
				out <- inp[i][j]
			}
		}
		done <- struct{}{}
	}(cha, out, inp...)
	return cha
}

// JoinByteWriterChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinByteWriterChan(out chan<- io.ByteWriter, inp <-chan io.ByteWriter) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- io.ByteWriter, inp <-chan io.ByteWriter) {
		defer close(done)
		for i := range inp {
			out <- i
		}
		done <- struct{}{}
	}(cha, out, inp)
	return cha
}

// DoneByteWriter returns a channel to receive one signal before close after inp has been drained.
func DoneByteWriter(inp <-chan io.ByteWriter) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, inp <-chan io.ByteWriter) {
		defer close(done)
		for i := range inp {
			_ = i // Drain inp
		}
		done <- struct{}{}
	}(cha, inp)
	return cha
}

// DoneByteWriterSlice returns a channel which will receive a slice
// of all the ByteWriters received on inp channel before close.
// Unlike DoneByteWriter, a full slice is sent once, not just an event.
func DoneByteWriterSlice(inp <-chan io.ByteWriter) (done <-chan []io.ByteWriter) {
	cha := make(chan []io.ByteWriter)
	go func(inp <-chan io.ByteWriter, done chan<- []io.ByteWriter) {
		defer close(done)
		ByteWriterS := []io.ByteWriter{}
		for i := range inp {
			ByteWriterS = append(ByteWriterS, i)
		}
		done <- ByteWriterS
	}(inp, cha)
	return cha
}

// DoneByteWriterFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneByteWriterFunc(inp <-chan io.ByteWriter, act func(a io.ByteWriter)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a io.ByteWriter) { return }
	}
	go func(done chan<- struct{}, inp <-chan io.ByteWriter, act func(a io.ByteWriter)) {
		defer close(done)
		for i := range inp {
			act(i) // Apply action
		}
		done <- struct{}{}
	}(cha, inp, act)
	return cha
}

// PipeByteWriterBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeByteWriterBuffer(inp <-chan io.ByteWriter, cap int) (out <-chan io.ByteWriter) {
	cha := make(chan io.ByteWriter, cap)
	go func(out chan<- io.ByteWriter, inp <-chan io.ByteWriter) {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}(cha, inp)
	return cha
}

// PipeByteWriterFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeByteWriterMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeByteWriterFunc(inp <-chan io.ByteWriter, act func(a io.ByteWriter) io.ByteWriter) (out <-chan io.ByteWriter) {
	cha := make(chan io.ByteWriter)
	if act == nil {
		act = func(a io.ByteWriter) io.ByteWriter { return a }
	}
	go func(out chan<- io.ByteWriter, inp <-chan io.ByteWriter, act func(a io.ByteWriter) io.ByteWriter) {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}(cha, inp, act)
	return cha
}

// PipeByteWriterFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeByteWriterFork(inp <-chan io.ByteWriter) (out1, out2 <-chan io.ByteWriter) {
	cha1 := make(chan io.ByteWriter)
	cha2 := make(chan io.ByteWriter)
	go func(out1, out2 chan<- io.ByteWriter, inp <-chan io.ByteWriter) {
		defer close(out1)
		defer close(out2)
		for i := range inp {
			out1 <- i
			out2 <- i
		}
	}(cha1, cha2, inp)
	return cha1, cha2
}

// ByteWriterTube is the signature for a pipe function.
type ByteWriterTube func(inp <-chan io.ByteWriter, out <-chan io.ByteWriter)

// ByteWriterDaisy returns a channel to receive all inp after having passed thru tube.
func ByteWriterDaisy(inp <-chan io.ByteWriter, tube ByteWriterTube) (out <-chan io.ByteWriter) {
	cha := make(chan io.ByteWriter)
	go tube(inp, cha)
	return cha
}

// ByteWriterDaisyChain returns a channel to receive all inp after having passed thru all tubes.
func ByteWriterDaisyChain(inp <-chan io.ByteWriter, tubes ...ByteWriterTube) (out <-chan io.ByteWriter) {
	cha := inp
	for i := range tubes {
		cha = ByteWriterDaisy(cha, tubes[i])
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
