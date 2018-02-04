// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package IsByte

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	"bytes"
)

// MakeBufferChan returns a new open channel
// (simply a 'chan bytes.Buffer' that is).
//
// Note: No 'Buffer-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myBufferPipelineStartsHere := MakeBufferChan()
//	// ... lot's of code to design and build Your favourite "myBufferWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myBufferPipelineStartsHere <- drop
//	}
//	close(myBufferPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeBufferBuffer) the channel is unbuffered.
//
func MakeBufferChan() (out chan bytes.Buffer) {
	return make(chan bytes.Buffer)
}

// ChanBuffer returns a channel to receive all inputs before close.
func ChanBuffer(inp ...bytes.Buffer) (out <-chan bytes.Buffer) {
	cha := make(chan bytes.Buffer)
	go func(out chan<- bytes.Buffer, inp ...bytes.Buffer) {
		defer close(out)
		for i := range inp {
			out <- inp[i]
		}
	}(cha, inp...)
	return cha
}

// ChanBufferSlice returns a channel to receive all inputs before close.
func ChanBufferSlice(inp ...[]bytes.Buffer) (out <-chan bytes.Buffer) {
	cha := make(chan bytes.Buffer)
	go func(out chan<- bytes.Buffer, inp ...[]bytes.Buffer) {
		defer close(out)
		for i := range inp {
			for j := range inp[i] {
				out <- inp[i][j]
			}
		}
	}(cha, inp...)
	return cha
}

// ChanBufferFuncNok returns a channel to receive all results of act until nok before close.
func ChanBufferFuncNok(act func() (bytes.Buffer, bool)) (out <-chan bytes.Buffer) {
	cha := make(chan bytes.Buffer)
	go func(out chan<- bytes.Buffer, act func() (bytes.Buffer, bool)) {
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

// ChanBufferFuncErr returns a channel to receive all results of act until err != nil before close.
func ChanBufferFuncErr(act func() (bytes.Buffer, error)) (out <-chan bytes.Buffer) {
	cha := make(chan bytes.Buffer)
	go func(out chan<- bytes.Buffer, act func() (bytes.Buffer, error)) {
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

// JoinBuffer sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinBuffer(out chan<- bytes.Buffer, inp ...bytes.Buffer) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- bytes.Buffer, inp ...bytes.Buffer) {
		defer close(done)
		for i := range inp {
			out <- inp[i]
		}
		done <- struct{}{}
	}(cha, out, inp...)
	return cha
}

// JoinBufferSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinBufferSlice(out chan<- bytes.Buffer, inp ...[]bytes.Buffer) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- bytes.Buffer, inp ...[]bytes.Buffer) {
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

// JoinBufferChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinBufferChan(out chan<- bytes.Buffer, inp <-chan bytes.Buffer) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- bytes.Buffer, inp <-chan bytes.Buffer) {
		defer close(done)
		for i := range inp {
			out <- i
		}
		done <- struct{}{}
	}(cha, out, inp)
	return cha
}

// DoneBuffer returns a channel to receive one signal before close after inp has been drained.
func DoneBuffer(inp <-chan bytes.Buffer) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, inp <-chan bytes.Buffer) {
		defer close(done)
		for i := range inp {
			_ = i // Drain inp
		}
		done <- struct{}{}
	}(cha, inp)
	return cha
}

// DoneBufferSlice returns a channel which will receive a slice
// of all the Buffers received on inp channel before close.
// Unlike DoneBuffer, a full slice is sent once, not just an event.
func DoneBufferSlice(inp <-chan bytes.Buffer) (done <-chan []bytes.Buffer) {
	cha := make(chan []bytes.Buffer)
	go func(inp <-chan bytes.Buffer, done chan<- []bytes.Buffer) {
		defer close(done)
		BufferS := []bytes.Buffer{}
		for i := range inp {
			BufferS = append(BufferS, i)
		}
		done <- BufferS
	}(inp, cha)
	return cha
}

// DoneBufferFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneBufferFunc(inp <-chan bytes.Buffer, act func(a bytes.Buffer)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a bytes.Buffer) { return }
	}
	go func(done chan<- struct{}, inp <-chan bytes.Buffer, act func(a bytes.Buffer)) {
		defer close(done)
		for i := range inp {
			act(i) // Apply action
		}
		done <- struct{}{}
	}(cha, inp, act)
	return cha
}

// PipeBufferBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeBufferBuffer(inp <-chan bytes.Buffer, cap int) (out <-chan bytes.Buffer) {
	cha := make(chan bytes.Buffer, cap)
	go func(out chan<- bytes.Buffer, inp <-chan bytes.Buffer) {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}(cha, inp)
	return cha
}

// PipeBufferFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeBufferMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeBufferFunc(inp <-chan bytes.Buffer, act func(a bytes.Buffer) bytes.Buffer) (out <-chan bytes.Buffer) {
	cha := make(chan bytes.Buffer)
	if act == nil {
		act = func(a bytes.Buffer) bytes.Buffer { return a }
	}
	go func(out chan<- bytes.Buffer, inp <-chan bytes.Buffer, act func(a bytes.Buffer) bytes.Buffer) {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}(cha, inp, act)
	return cha
}

// PipeBufferFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeBufferFork(inp <-chan bytes.Buffer) (out1, out2 <-chan bytes.Buffer) {
	cha1 := make(chan bytes.Buffer)
	cha2 := make(chan bytes.Buffer)
	go func(out1, out2 chan<- bytes.Buffer, inp <-chan bytes.Buffer) {
		defer close(out1)
		defer close(out2)
		for i := range inp {
			out1 <- i
			out2 <- i
		}
	}(cha1, cha2, inp)
	return cha1, cha2
}

// BufferTube is the signature for a pipe function.
type BufferTube func(inp <-chan bytes.Buffer, out <-chan bytes.Buffer)

// BufferDaisy returns a channel to receive all inp after having passed thru tube.
func BufferDaisy(inp <-chan bytes.Buffer, tube BufferTube) (out <-chan bytes.Buffer) {
	cha := make(chan bytes.Buffer)
	go tube(inp, cha)
	return cha
}

// BufferDaisyChain returns a channel to receive all inp after having passed thru all tubes.
func BufferDaisyChain(inp <-chan bytes.Buffer, tubes ...BufferTube) (out <-chan bytes.Buffer) {
	cha := inp
	for i := range tubes {
		cha = BufferDaisy(cha, tubes[i])
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
