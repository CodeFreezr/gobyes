// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package list

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	list "container/list"
)

// MakeElementChan returns a new open channel
// (simply a 'chan *list.Element' that is).
//
// Note: No 'Element-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myElementPipelineStartsHere := MakeElementChan()
//	// ... lot's of code to design and build Your favourite "myElementWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myElementPipelineStartsHere <- drop
//	}
//	close(myElementPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeElementBuffer) the channel is unbuffered.
//
func MakeElementChan() (out chan *list.Element) {
	return make(chan *list.Element)
}

func sendElement(out chan<- *list.Element, inp ...*list.Element) {
	defer close(out)
	for i := range inp {
		out <- inp[i]
	}
}

// ChanElement returns a channel to receive all inputs before close.
func ChanElement(inp ...*list.Element) (out <-chan *list.Element) {
	cha := make(chan *list.Element)
	go sendElement(cha, inp...)
	return cha
}

func sendElementSlice(out chan<- *list.Element, inp ...[]*list.Element) {
	defer close(out)
	for i := range inp {
		for j := range inp[i] {
			out <- inp[i][j]
		}
	}
}

// ChanElementSlice returns a channel to receive all inputs before close.
func ChanElementSlice(inp ...[]*list.Element) (out <-chan *list.Element) {
	cha := make(chan *list.Element)
	go sendElementSlice(cha, inp...)
	return cha
}

func chanElementFuncNil(out chan<- *list.Element, act func() *list.Element) {
	defer close(out)
	for {
		res := act() // Apply action
		if res == nil {
			return
		}
		out <- res
	}
}

// ChanElementFuncNil returns a channel to receive all results of act until nil before close.
func ChanElementFuncNil(act func() *list.Element) (out <-chan *list.Element) {
	cha := make(chan *list.Element)
	go chanElementFuncNil(cha, act)
	return cha
}

func chanElementFuncNok(out chan<- *list.Element, act func() (*list.Element, bool)) {
	defer close(out)
	for {
		res, ok := act() // Apply action
		if !ok {
			return
		}
		out <- res
	}
}

// ChanElementFuncNok returns a channel to receive all results of act until nok before close.
func ChanElementFuncNok(act func() (*list.Element, bool)) (out <-chan *list.Element) {
	cha := make(chan *list.Element)
	go chanElementFuncNok(cha, act)
	return cha
}

func chanElementFuncErr(out chan<- *list.Element, act func() (*list.Element, error)) {
	defer close(out)
	for {
		res, err := act() // Apply action
		if err != nil {
			return
		}
		out <- res
	}
}

// ChanElementFuncErr returns a channel to receive all results of act until err != nil before close.
func ChanElementFuncErr(act func() (*list.Element, error)) (out <-chan *list.Element) {
	cha := make(chan *list.Element)
	go chanElementFuncErr(cha, act)
	return cha
}

func joinElement(done chan<- struct{}, out chan<- *list.Element, inp ...*list.Element) {
	defer close(done)
	for i := range inp {
		out <- inp[i]
	}
	done <- struct{}{}
}

// JoinElement sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinElement(out chan<- *list.Element, inp ...*list.Element) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinElement(cha, out, inp...)
	return cha
}

func joinElementSlice(done chan<- struct{}, out chan<- *list.Element, inp ...[]*list.Element) {
	defer close(done)
	for i := range inp {
		for j := range inp[i] {
			out <- inp[i][j]
		}
	}
	done <- struct{}{}
}

// JoinElementSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinElementSlice(out chan<- *list.Element, inp ...[]*list.Element) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinElementSlice(cha, out, inp...)
	return cha
}

func joinElementChan(done chan<- struct{}, out chan<- *list.Element, inp <-chan *list.Element) {
	defer close(done)
	for i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinElementChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinElementChan(out chan<- *list.Element, inp <-chan *list.Element) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinElementChan(cha, out, inp)
	return cha
}

func doitElement(done chan<- struct{}, inp <-chan *list.Element) {
	defer close(done)
	for i := range inp {
		_ = i // Drain inp
	}
	done <- struct{}{}
}

// DoneElement returns a channel to receive one signal before close after inp has been drained.
func DoneElement(inp <-chan *list.Element) (done <-chan struct{}) {
	cha := make(chan struct{})
	go doitElement(cha, inp)
	return cha
}

func doitElementSlice(done chan<- ([]*list.Element), inp <-chan *list.Element) {
	defer close(done)
	ElementS := []*list.Element{}
	for i := range inp {
		ElementS = append(ElementS, i)
	}
	done <- ElementS
}

// DoneElementSlice returns a channel which will receive a slice
// of all the Elements received on inp channel before close.
// Unlike DoneElement, a full slice is sent once, not just an event.
func DoneElementSlice(inp <-chan *list.Element) (done <-chan ([]*list.Element)) {
	cha := make(chan ([]*list.Element))
	go doitElementSlice(cha, inp)
	return cha
}

func doitElementFunc(done chan<- struct{}, inp <-chan *list.Element, act func(a *list.Element)) {
	defer close(done)
	for i := range inp {
		act(i) // Apply action
	}
	done <- struct{}{}
}

// DoneElementFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneElementFunc(inp <-chan *list.Element, act func(a *list.Element)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a *list.Element) { return }
	}
	go doitElementFunc(cha, inp, act)
	return cha
}

func pipeElementBuffer(out chan<- *list.Element, inp <-chan *list.Element) {
	defer close(out)
	for i := range inp {
		out <- i
	}
}

// PipeElementBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeElementBuffer(inp <-chan *list.Element, cap int) (out <-chan *list.Element) {
	cha := make(chan *list.Element, cap)
	go pipeElementBuffer(cha, inp)
	return cha
}

func pipeElementFunc(out chan<- *list.Element, inp <-chan *list.Element, act func(a *list.Element) *list.Element) {
	defer close(out)
	for i := range inp {
		out <- act(i)
	}
}

// PipeElementFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeElementMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeElementFunc(inp <-chan *list.Element, act func(a *list.Element) *list.Element) (out <-chan *list.Element) {
	cha := make(chan *list.Element)
	if act == nil {
		act = func(a *list.Element) *list.Element { return a }
	}
	go pipeElementFunc(cha, inp, act)
	return cha
}

func pipeElementFork(out1, out2 chan<- *list.Element, inp <-chan *list.Element) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		out1 <- i
		out2 <- i
	}
}

// PipeElementFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeElementFork(inp <-chan *list.Element) (out1, out2 <-chan *list.Element) {
	cha1 := make(chan *list.Element)
	cha2 := make(chan *list.Element)
	go pipeElementFork(cha1, cha2, inp)
	return cha1, cha2
}

// ElementTube is the signature for a pipe function.
type ElementTube func(inp <-chan *list.Element, out <-chan *list.Element)

// ElementDaisy returns a channel to receive all inp after having passed thru tube.
func ElementDaisy(inp <-chan *list.Element, tube ElementTube) (out <-chan *list.Element) {
	cha := make(chan *list.Element)
	go tube(inp, cha)
	return cha
}

// ElementDaisyChain returns a channel to receive all inp after having passed thru all tubes.
func ElementDaisyChain(inp <-chan *list.Element, tubes ...ElementTube) (out <-chan *list.Element) {
	cha := inp
	for i := range tubes {
		cha = ElementDaisy(cha, tubes[i])
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
