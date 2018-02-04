// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tar

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	tar "archive/tar"
)

// MakeWriterChan returns a new open channel
// (simply a 'chan *tar.Writer' that is).
//
// Note: No 'Writer-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myWriterPipelineStartsHere := MakeWriterChan()
//	// ... lot's of code to design and build Your favourite "myWriterWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myWriterPipelineStartsHere <- drop
//	}
//	close(myWriterPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeWriterBuffer) the channel is unbuffered.
//
func MakeWriterChan() (out chan *tar.Writer) {
	return make(chan *tar.Writer)
}

// ChanWriter returns a channel to receive all inputs before close.
func ChanWriter(inp ...*tar.Writer) (out <-chan *tar.Writer) {
	cha := make(chan *tar.Writer)
	go func(out chan<- *tar.Writer, inp ...*tar.Writer) {
		defer close(out)
		for i := range inp {
			out <- inp[i]
		}
	}(cha, inp...)
	return cha
}

// ChanWriterSlice returns a channel to receive all inputs before close.
func ChanWriterSlice(inp ...[]*tar.Writer) (out <-chan *tar.Writer) {
	cha := make(chan *tar.Writer)
	go func(out chan<- *tar.Writer, inp ...[]*tar.Writer) {
		defer close(out)
		for i := range inp {
			for j := range inp[i] {
				out <- inp[i][j]
			}
		}
	}(cha, inp...)
	return cha
}

// ChanWriterFuncNil returns a channel to receive all results of act until nil before close.
func ChanWriterFuncNil(act func() *tar.Writer) (out <-chan *tar.Writer) {
	cha := make(chan *tar.Writer)
	go func(out chan<- *tar.Writer, act func() *tar.Writer) {
		defer close(out)
		for {
			res := act() // Apply action
			if res == nil {
				return
			}
			out <- res
		}
	}(cha, act)
	return cha
}

// ChanWriterFuncNok returns a channel to receive all results of act until nok before close.
func ChanWriterFuncNok(act func() (*tar.Writer, bool)) (out <-chan *tar.Writer) {
	cha := make(chan *tar.Writer)
	go func(out chan<- *tar.Writer, act func() (*tar.Writer, bool)) {
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

// ChanWriterFuncErr returns a channel to receive all results of act until err != nil before close.
func ChanWriterFuncErr(act func() (*tar.Writer, error)) (out <-chan *tar.Writer) {
	cha := make(chan *tar.Writer)
	go func(out chan<- *tar.Writer, act func() (*tar.Writer, error)) {
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

// JoinWriter sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinWriter(out chan<- *tar.Writer, inp ...*tar.Writer) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- *tar.Writer, inp ...*tar.Writer) {
		defer close(done)
		for i := range inp {
			out <- inp[i]
		}
		done <- struct{}{}
	}(cha, out, inp...)
	return cha
}

// JoinWriterSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinWriterSlice(out chan<- *tar.Writer, inp ...[]*tar.Writer) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- *tar.Writer, inp ...[]*tar.Writer) {
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

// JoinWriterChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinWriterChan(out chan<- *tar.Writer, inp <-chan *tar.Writer) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, out chan<- *tar.Writer, inp <-chan *tar.Writer) {
		defer close(done)
		for i := range inp {
			out <- i
		}
		done <- struct{}{}
	}(cha, out, inp)
	return cha
}

// DoneWriter returns a channel to receive one signal before close after inp has been drained.
func DoneWriter(inp <-chan *tar.Writer) (done <-chan struct{}) {
	cha := make(chan struct{})
	go func(done chan<- struct{}, inp <-chan *tar.Writer) {
		defer close(done)
		for i := range inp {
			_ = i // Drain inp
		}
		done <- struct{}{}
	}(cha, inp)
	return cha
}

// DoneWriterSlice returns a channel which will receive a slice
// of all the Writers received on inp channel before close.
// Unlike DoneWriter, a full slice is sent once, not just an event.
func DoneWriterSlice(inp <-chan *tar.Writer) (done <-chan []*tar.Writer) {
	cha := make(chan []*tar.Writer)
	go func(inp <-chan *tar.Writer, done chan<- []*tar.Writer) {
		defer close(done)
		WriterS := []*tar.Writer{}
		for i := range inp {
			WriterS = append(WriterS, i)
		}
		done <- WriterS
	}(inp, cha)
	return cha
}

// DoneWriterFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneWriterFunc(inp <-chan *tar.Writer, act func(a *tar.Writer)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a *tar.Writer) { return }
	}
	go func(done chan<- struct{}, inp <-chan *tar.Writer, act func(a *tar.Writer)) {
		defer close(done)
		for i := range inp {
			act(i) // Apply action
		}
		done <- struct{}{}
	}(cha, inp, act)
	return cha
}

// PipeWriterBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeWriterBuffer(inp <-chan *tar.Writer, cap int) (out <-chan *tar.Writer) {
	cha := make(chan *tar.Writer, cap)
	go func(out chan<- *tar.Writer, inp <-chan *tar.Writer) {
		defer close(out)
		for i := range inp {
			out <- i
		}
	}(cha, inp)
	return cha
}

// PipeWriterFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeWriterMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeWriterFunc(inp <-chan *tar.Writer, act func(a *tar.Writer) *tar.Writer) (out <-chan *tar.Writer) {
	cha := make(chan *tar.Writer)
	if act == nil {
		act = func(a *tar.Writer) *tar.Writer { return a }
	}
	go func(out chan<- *tar.Writer, inp <-chan *tar.Writer, act func(a *tar.Writer) *tar.Writer) {
		defer close(out)
		for i := range inp {
			out <- act(i)
		}
	}(cha, inp, act)
	return cha
}

// PipeWriterFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeWriterFork(inp <-chan *tar.Writer) (out1, out2 <-chan *tar.Writer) {
	cha1 := make(chan *tar.Writer)
	cha2 := make(chan *tar.Writer)
	go func(out1, out2 chan<- *tar.Writer, inp <-chan *tar.Writer) {
		defer close(out1)
		defer close(out2)
		for i := range inp {
			out1 <- i
			out2 <- i
		}
	}(cha1, cha2, inp)
	return cha1, cha2
}

// WriterTube is the signature for a pipe function.
type WriterTube func(inp <-chan *tar.Writer, out <-chan *tar.Writer)

// WriterDaisy returns a channel to receive all inp after having passed thru tube.
func WriterDaisy(inp <-chan *tar.Writer, tube WriterTube) (out <-chan *tar.Writer) {
	cha := make(chan *tar.Writer)
	go tube(inp, cha)
	return cha
}

// WriterDaisyChain returns a channel to receive all inp after having passed thru all tubes.
func WriterDaisyChain(inp <-chan *tar.Writer, tubes ...WriterTube) (out <-chan *tar.Writer) {
	cha := inp
	for i := range tubes {
		cha = WriterDaisy(cha, tubes[i])
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
