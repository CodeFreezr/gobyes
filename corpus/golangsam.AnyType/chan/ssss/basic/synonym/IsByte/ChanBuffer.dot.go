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
func MakeBufferChan() chan bytes.Buffer {
	return make(chan bytes.Buffer)
}

// ChanBuffer returns a channel to receive all inputs before close.
func ChanBuffer(inp ...bytes.Buffer) chan bytes.Buffer {
	out := make(chan bytes.Buffer)
	go func() {
		defer close(out)
		for i := range inp {
			out <- inp[i]
		}
	}()
	return out
}

// ChanBufferSlice returns a channel to receive all inputs before close.
func ChanBufferSlice(inp ...[]bytes.Buffer) chan bytes.Buffer {
	out := make(chan bytes.Buffer)
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

// ChanBufferFuncNok returns a channel to receive all results of act until nok before close.
func ChanBufferFuncNok(act func() (bytes.Buffer, bool)) <-chan bytes.Buffer {
	out := make(chan bytes.Buffer)
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

// ChanBufferFuncErr returns a channel to receive all results of act until err != nil before close.
func ChanBufferFuncErr(act func() (bytes.Buffer, error)) <-chan bytes.Buffer {
	out := make(chan bytes.Buffer)
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

// JoinBuffer sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinBuffer(out chan<- bytes.Buffer, inp ...bytes.Buffer) chan struct{} {
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

// JoinBufferSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinBufferSlice(out chan<- bytes.Buffer, inp ...[]bytes.Buffer) chan struct{} {
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

// JoinBufferChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinBufferChan(out chan<- bytes.Buffer, inp <-chan bytes.Buffer) chan struct{} {
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

// DoneBuffer returns a channel to receive one signal before close after inp has been drained.
func DoneBuffer(inp <-chan bytes.Buffer) chan struct{} {
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

// DoneBufferSlice returns a channel which will receive a slice
// of all the Buffers received on inp channel before close.
// Unlike DoneBuffer, a full slice is sent once, not just an event.
func DoneBufferSlice(inp <-chan bytes.Buffer) chan []bytes.Buffer {
	done := make(chan []bytes.Buffer)
	go func() {
		defer close(done)
		BufferS := []bytes.Buffer{}
		for i := range inp {
			BufferS = append(BufferS, i)
		}
		done <- BufferS
	}()
	return done
}

// DoneBufferFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneBufferFunc(inp <-chan bytes.Buffer, act func(a bytes.Buffer)) chan struct{} {
	done := make(chan struct{})
	if act == nil {
		act = func(a bytes.Buffer) { return }
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

// PipeBufferBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeBufferBuffer(inp <-chan bytes.Buffer, cap int) chan bytes.Buffer {
	out := make(chan bytes.Buffer, cap)
	go func() {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}()
	return out
}

// PipeBufferFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeBufferMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeBufferFunc(inp <-chan bytes.Buffer, act func(a bytes.Buffer) bytes.Buffer) chan bytes.Buffer {
	out := make(chan bytes.Buffer)
	if act == nil {
		act = func(a bytes.Buffer) bytes.Buffer { return a }
	}
	go func() {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}()
	return out
}

// PipeBufferFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeBufferFork(inp <-chan bytes.Buffer) (chan bytes.Buffer, chan bytes.Buffer) {
	out1 := make(chan bytes.Buffer)
	out2 := make(chan bytes.Buffer)
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
