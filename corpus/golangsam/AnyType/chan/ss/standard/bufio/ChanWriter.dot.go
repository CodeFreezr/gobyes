// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bufio

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	bufio "bufio"
)

// MakeWriterChan returns a new open channel
// (simply a 'chan *bufio.Writer' that is).
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
func MakeWriterChan() (out chan *bufio.Writer) {
	return make(chan *bufio.Writer)
}

func sendWriter(out chan<- *bufio.Writer, inp ...*bufio.Writer) {
	defer close(out)
	for i := range inp {
		out <- inp[i]
	}
}

// ChanWriter returns a channel to receive all inputs before close.
func ChanWriter(inp ...*bufio.Writer) (out <-chan *bufio.Writer) {
	cha := make(chan *bufio.Writer)
	go sendWriter(cha, inp...)
	return cha
}

func sendWriterSlice(out chan<- *bufio.Writer, inp ...[]*bufio.Writer) {
	defer close(out)
	for i := range inp {
		for j := range inp[i] {
			out <- inp[i][j]
		}
	}
}

// ChanWriterSlice returns a channel to receive all inputs before close.
func ChanWriterSlice(inp ...[]*bufio.Writer) (out <-chan *bufio.Writer) {
	cha := make(chan *bufio.Writer)
	go sendWriterSlice(cha, inp...)
	return cha
}

func chanWriterFuncNok(out chan<- *bufio.Writer, act func() (*bufio.Writer, bool)) {
	defer close(out)
	for {
		res, ok := act() // Apply action
		if !ok {
			return
		}
		out <- res
	}
}

// ChanWriterFuncNok returns a channel to receive all results of act until nok before close.
func ChanWriterFuncNok(act func() (*bufio.Writer, bool)) (out <-chan *bufio.Writer) {
	cha := make(chan *bufio.Writer)
	go chanWriterFuncNok(cha, act)
	return cha
}

func chanWriterFuncErr(out chan<- *bufio.Writer, act func() (*bufio.Writer, error)) {
	defer close(out)
	for {
		res, err := act() // Apply action
		if err != nil {
			return
		}
		out <- res
	}
}

// ChanWriterFuncErr returns a channel to receive all results of act until err != nil before close.
func ChanWriterFuncErr(act func() (*bufio.Writer, error)) (out <-chan *bufio.Writer) {
	cha := make(chan *bufio.Writer)
	go chanWriterFuncErr(cha, act)
	return cha
}

func joinWriter(done chan<- struct{}, out chan<- *bufio.Writer, inp ...*bufio.Writer) {
	defer close(done)
	for i := range inp {
		out <- inp[i]
	}
	done <- struct{}{}
}

// JoinWriter sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinWriter(out chan<- *bufio.Writer, inp ...*bufio.Writer) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinWriter(cha, out, inp...)
	return cha
}

func joinWriterSlice(done chan<- struct{}, out chan<- *bufio.Writer, inp ...[]*bufio.Writer) {
	defer close(done)
	for i := range inp {
		for j := range inp[i] {
			out <- inp[i][j]
		}
	}
	done <- struct{}{}
}

// JoinWriterSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinWriterSlice(out chan<- *bufio.Writer, inp ...[]*bufio.Writer) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinWriterSlice(cha, out, inp...)
	return cha
}

func joinWriterChan(done chan<- struct{}, out chan<- *bufio.Writer, inp <-chan *bufio.Writer) {
	defer close(done)
	for i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinWriterChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinWriterChan(out chan<- *bufio.Writer, inp <-chan *bufio.Writer) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinWriterChan(cha, out, inp)
	return cha
}

func doitWriter(done chan<- struct{}, inp <-chan *bufio.Writer) {
	defer close(done)
	for i := range inp {
		_ = i // Drain inp
	}
	done <- struct{}{}
}

// DoneWriter returns a channel to receive one signal before close after inp has been drained.
func DoneWriter(inp <-chan *bufio.Writer) (done <-chan struct{}) {
	cha := make(chan struct{})
	go doitWriter(cha, inp)
	return cha
}

func doitWriterSlice(done chan<- ([]*bufio.Writer), inp <-chan *bufio.Writer) {
	defer close(done)
	WriterS := []*bufio.Writer{}
	for i := range inp {
		WriterS = append(WriterS, i)
	}
	done <- WriterS
}

// DoneWriterSlice returns a channel which will receive a slice
// of all the Writers received on inp channel before close.
// Unlike DoneWriter, a full slice is sent once, not just an event.
func DoneWriterSlice(inp <-chan *bufio.Writer) (done <-chan ([]*bufio.Writer)) {
	cha := make(chan ([]*bufio.Writer))
	go doitWriterSlice(cha, inp)
	return cha
}

func doitWriterFunc(done chan<- struct{}, inp <-chan *bufio.Writer, act func(a *bufio.Writer)) {
	defer close(done)
	for i := range inp {
		act(i) // Apply action
	}
	done <- struct{}{}
}

// DoneWriterFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneWriterFunc(inp <-chan *bufio.Writer, act func(a *bufio.Writer)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a *bufio.Writer) { return }
	}
	go doitWriterFunc(cha, inp, act)
	return cha
}

func pipeWriterBuffer(out chan<- *bufio.Writer, inp <-chan *bufio.Writer) {
	defer close(out)
	for i := range inp {
		out <- i
	}
}

// PipeWriterBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeWriterBuffer(inp <-chan *bufio.Writer, cap int) (out <-chan *bufio.Writer) {
	cha := make(chan *bufio.Writer, cap)
	go pipeWriterBuffer(cha, inp)
	return cha
}

func pipeWriterFunc(out chan<- *bufio.Writer, inp <-chan *bufio.Writer, act func(a *bufio.Writer) *bufio.Writer) {
	defer close(out)
	for i := range inp {
		out <- act(i)
	}
}

// PipeWriterFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeWriterMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeWriterFunc(inp <-chan *bufio.Writer, act func(a *bufio.Writer) *bufio.Writer) (out <-chan *bufio.Writer) {
	cha := make(chan *bufio.Writer)
	if act == nil {
		act = func(a *bufio.Writer) *bufio.Writer { return a }
	}
	go pipeWriterFunc(cha, inp, act)
	return cha
}

func pipeWriterFork(out1, out2 chan<- *bufio.Writer, inp <-chan *bufio.Writer) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		out1 <- i
		out2 <- i
	}
}

// PipeWriterFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeWriterFork(inp <-chan *bufio.Writer) (out1, out2 <-chan *bufio.Writer) {
	cha1 := make(chan *bufio.Writer)
	cha2 := make(chan *bufio.Writer)
	go pipeWriterFork(cha1, cha2, inp)
	return cha1, cha2
}

// WriterTube is the signature for a pipe function.
type WriterTube func(inp <-chan *bufio.Writer, out <-chan *bufio.Writer)

// WriterDaisy returns a channel to receive all inp after having passed thru tube.
func WriterDaisy(inp <-chan *bufio.Writer, tube WriterTube) (out <-chan *bufio.Writer) {
	cha := make(chan *bufio.Writer)
	go tube(inp, cha)
	return cha
}

// WriterDaisyChain returns a channel to receive all inp after having passed thru all tubes.
func WriterDaisyChain(inp <-chan *bufio.Writer, tubes ...WriterTube) (out <-chan *bufio.Writer) {
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
