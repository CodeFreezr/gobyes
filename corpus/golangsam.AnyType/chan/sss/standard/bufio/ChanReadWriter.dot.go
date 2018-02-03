// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bufio

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	bufio "bufio"
)

// MakeReadWriterChan returns a new open channel
// (simply a 'chan *bufio.ReadWriter' that is).
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
func MakeReadWriterChan() (out chan *bufio.ReadWriter) {
	return make(chan *bufio.ReadWriter)
}

// ChanReadWriter returns a channel to receive all inputs before close.
func ChanReadWriter(inp ...*bufio.ReadWriter) (out <-chan *bufio.ReadWriter) {
	cha := make(chan *bufio.ReadWriter)
	go func(out chan<- *bufio.ReadWriter, inp ...*bufio.ReadWriter) {
		defer close(out)
		for i := range inp {
			out <- inp[i]
		}
	}(cha, inp...)
	return cha
}

// ChanReadWriterSlice returns a channel to receive all inputs before close.
func ChanReadWriterSlice(inp ...[]*bufio.ReadWriter) (out <-chan *bufio.ReadWriter) {
	cha := make(chan *bufio.ReadWriter)
	go func(out chan<- *bufio.ReadWriter, inp ...[]*bufio.ReadWriter) {
		defer close(out)
		for i := range inp {
			for j := range inp[i] {
				out <- inp[i][j]
			}
		}
	}(cha, inp...)
	return cha
}

// ChanReadWriterFuncNok returns a channel to receive all results of act until nok before close.
func ChanReadWriterFuncNok(act func() (*bufio.ReadWriter, bool)) (out <-chan *bufio.ReadWriter) {
	cha := make(chan *bufio.ReadWriter)
	go func(out chan<- *bufio.ReadWriter, act func() (*bufio.ReadWriter, bool)) {
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

// ChanReadWriterFuncErr returns a channel to receive all results of act until err != nil before close.
func ChanReadWriterFuncErr(act func() (*bufio.ReadWriter, error)) (out <-chan *bufio.ReadWriter) {
	cha := make(chan *bufio.ReadWriter)
	go func(out chan<- *bufio.ReadWriter, act func() (*bufio.ReadWriter, error)) {
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

// JoinReadWriter sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinReadWriter(out chan<- *bufio.ReadWriter, inp ...*bufio.ReadWriter) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- *bufio.ReadWriter, inp ...*bufio.ReadWriter) {
		defer close(done)
		for i := range inp {
			out <- inp[i]
		}
		done <- struct{}{}
	}(cha, out, inp...)
	return cha
}

// JoinReadWriterSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinReadWriterSlice(out chan<- *bufio.ReadWriter, inp ...[]*bufio.ReadWriter) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- *bufio.ReadWriter, inp ...[]*bufio.ReadWriter) {
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

// JoinReadWriterChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinReadWriterChan(out chan<- *bufio.ReadWriter, inp <-chan *bufio.ReadWriter) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- *bufio.ReadWriter, inp <-chan *bufio.ReadWriter) {
		defer close(done)
		for i := range inp {
			out <- i
		}
		done <- struct{}{}
	}(cha, out, inp)
	return cha
}

// DoneReadWriter returns a channel to receive one signal before close after inp has been drained.
func DoneReadWriter(inp <-chan *bufio.ReadWriter) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, inp <-chan *bufio.ReadWriter) {
		defer close(done)
		for i := range inp {
			_ = i // Drain inp
		}
		done <- struct{}{}
	}(cha, inp)
	return cha
}

// DoneReadWriterSlice returns a channel which will receive a slice
// of all the ReadWriters received on inp channel before close.
// Unlike DoneReadWriter, a full slice is sent once, not just an event.
func DoneReadWriterSlice(inp <-chan *bufio.ReadWriter) (done <-chan []*bufio.ReadWriter) {
	cha := make(chan []*bufio.ReadWriter)
	go func(inp <-chan *bufio.ReadWriter, done chan<- []*bufio.ReadWriter) {
		defer close(done)
		ReadWriterS := []*bufio.ReadWriter{}
		for i := range inp {
			ReadWriterS = append(ReadWriterS, i)
		}
		done <- ReadWriterS
	}(inp, cha)
	return cha
}

// DoneReadWriterFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneReadWriterFunc(inp <-chan *bufio.ReadWriter, act func(a *bufio.ReadWriter)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a *bufio.ReadWriter) { return }
	}
	go func(done chan<- struct{}, inp <-chan *bufio.ReadWriter, act func(a *bufio.ReadWriter)) {
		defer close(done)
		for i := range inp {
			act(i) // Apply action
		}
		done <- struct{}{}
	}(cha, inp, act)
	return cha
}

// PipeReadWriterBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeReadWriterBuffer(inp <-chan *bufio.ReadWriter, cap int) (out <-chan *bufio.ReadWriter) {
	cha := make(chan *bufio.ReadWriter, cap)
	go func(out chan<- *bufio.ReadWriter, inp <-chan *bufio.ReadWriter) {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}(cha, inp)
	return cha
}

// PipeReadWriterFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeReadWriterMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeReadWriterFunc(inp <-chan *bufio.ReadWriter, act func(a *bufio.ReadWriter) *bufio.ReadWriter) (out <-chan *bufio.ReadWriter) {
	cha := make(chan *bufio.ReadWriter)
	if act == nil {
		act = func(a *bufio.ReadWriter) *bufio.ReadWriter { return a }
	}
	go func(out chan<- *bufio.ReadWriter, inp <-chan *bufio.ReadWriter, act func(a *bufio.ReadWriter) *bufio.ReadWriter) {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}(cha, inp, act)
	return cha
}

// PipeReadWriterFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeReadWriterFork(inp <-chan *bufio.ReadWriter) (out1, out2 <-chan *bufio.ReadWriter) {
	cha1 := make(chan *bufio.ReadWriter)
	cha2 := make(chan *bufio.ReadWriter)
	go func(out1, out2 chan<- *bufio.ReadWriter, inp <-chan *bufio.ReadWriter) {
		defer close(out1)
		defer close(out2)
		for i := range inp {
			out1 <- i
			out2 <- i
		}
	}(cha1, cha2, inp)
	return cha1, cha2
}

// ReadWriterTube is the signature for a pipe function.
type ReadWriterTube func(inp <-chan *bufio.ReadWriter, out <-chan *bufio.ReadWriter)

// ReadWriterDaisy returns a channel to receive all inp after having passed thru tube.
func ReadWriterDaisy(inp <-chan *bufio.ReadWriter, tube ReadWriterTube) (out <-chan *bufio.ReadWriter) {
	cha := make(chan *bufio.ReadWriter)
	go tube(inp, cha)
	return cha
}

// ReadWriterDaisyChain returns a channel to receive all inp after having passed thru all tubes.
func ReadWriterDaisyChain(inp <-chan *bufio.ReadWriter, tubes ...ReadWriterTube) (out <-chan *bufio.ReadWriter) {
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
