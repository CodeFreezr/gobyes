// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package list

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	list "container/list"
)

// MakeListSChan returns a new open channel
// (simply a 'chan []*list.List' that is).
//
// Note: No 'ListS-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var myListSPipelineStartsHere := MakeListSChan()
//	// ... lot's of code to design and build Your favourite "myListSWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		myListSPipelineStartsHere <- drop
//	}
//	close(myListSPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeListSBuffer) the channel is unbuffered.
//
func MakeListSChan() (out chan []*list.List) {
	return make(chan []*list.List)
}

func sendListS(out chan<- []*list.List, inp ...[]*list.List) {
	defer close(out)
	for i := range inp {
		out <- inp[i]
	}
}

// ChanListS returns a channel to receive all inputs before close.
func ChanListS(inp ...[]*list.List) (out <-chan []*list.List) {
	cha := make(chan []*list.List)
	go sendListS(cha, inp...)
	return cha
}

func sendListSSlice(out chan<- []*list.List, inp ...[][]*list.List) {
	defer close(out)
	for i := range inp {
		for j := range inp[i] {
			out <- inp[i][j]
		}
	}
}

// ChanListSSlice returns a channel to receive all inputs before close.
func ChanListSSlice(inp ...[][]*list.List) (out <-chan []*list.List) {
	cha := make(chan []*list.List)
	go sendListSSlice(cha, inp...)
	return cha
}

func chanListSFuncNil(out chan<- []*list.List, act func() []*list.List) {
	defer close(out)
	for {
		res := act() // Apply action
		if res == nil {
			return
		}
		out <- res
	}
}

// ChanListSFuncNil returns a channel to receive all results of act until nil before close.
func ChanListSFuncNil(act func() []*list.List) (out <-chan []*list.List) {
	cha := make(chan []*list.List)
	go chanListSFuncNil(cha, act)
	return cha
}

func chanListSFuncNok(out chan<- []*list.List, act func() ([]*list.List, bool)) {
	defer close(out)
	for {
		res, ok := act() // Apply action
		if !ok {
			return
		}
		out <- res
	}
}

// ChanListSFuncNok returns a channel to receive all results of act until nok before close.
func ChanListSFuncNok(act func() ([]*list.List, bool)) (out <-chan []*list.List) {
	cha := make(chan []*list.List)
	go chanListSFuncNok(cha, act)
	return cha
}

func chanListSFuncErr(out chan<- []*list.List, act func() ([]*list.List, error)) {
	defer close(out)
	for {
		res, err := act() // Apply action
		if err != nil {
			return
		}
		out <- res
	}
}

// ChanListSFuncErr returns a channel to receive all results of act until err != nil before close.
func ChanListSFuncErr(act func() ([]*list.List, error)) (out <-chan []*list.List) {
	cha := make(chan []*list.List)
	go chanListSFuncErr(cha, act)
	return cha
}

func joinListS(done chan<- struct{}, out chan<- []*list.List, inp ...[]*list.List) {
	defer close(done)
	for i := range inp {
		out <- inp[i]
	}
	done <- struct{}{}
}

// JoinListS sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinListS(out chan<- []*list.List, inp ...[]*list.List) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinListS(cha, out, inp...)
	return cha
}

func joinListSSlice(done chan<- struct{}, out chan<- []*list.List, inp ...[][]*list.List) {
	defer close(done)
	for i := range inp {
		for j := range inp[i] {
			out <- inp[i][j]
		}
	}
	done <- struct{}{}
}

// JoinListSSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinListSSlice(out chan<- []*list.List, inp ...[][]*list.List) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinListSSlice(cha, out, inp...)
	return cha
}

func joinListSChan(done chan<- struct{}, out chan<- []*list.List, inp <-chan []*list.List) {
	defer close(done)
	for i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinListSChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinListSChan(out chan<- []*list.List, inp <-chan []*list.List) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinListSChan(cha, out, inp)
	return cha
}

func doitListS(done chan<- struct{}, inp <-chan []*list.List) {
	defer close(done)
	for i := range inp {
		_ = i // Drain inp
	}
	done <- struct{}{}
}

// DoneListS returns a channel to receive one signal before close after inp has been drained.
func DoneListS(inp <-chan []*list.List) (done <-chan struct{}) {
	cha := make(chan struct{})
	go doitListS(cha, inp)
	return cha
}

func doitListSSlice(done chan<- ([][]*list.List), inp <-chan []*list.List) {
	defer close(done)
	ListSS := [][]*list.List{}
	for i := range inp {
		ListSS = append(ListSS, i)
	}
	done <- ListSS
}

// DoneListSSlice returns a channel which will receive a slice
// of all the ListSs received on inp channel before close.
// Unlike DoneListS, a full slice is sent once, not just an event.
func DoneListSSlice(inp <-chan []*list.List) (done <-chan ([][]*list.List)) {
	cha := make(chan ([][]*list.List))
	go doitListSSlice(cha, inp)
	return cha
}

func doitListSFunc(done chan<- struct{}, inp <-chan []*list.List, act func(a []*list.List)) {
	defer close(done)
	for i := range inp {
		act(i) // Apply action
	}
	done <- struct{}{}
}

// DoneListSFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneListSFunc(inp <-chan []*list.List, act func(a []*list.List)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a []*list.List) { return }
	}
	go doitListSFunc(cha, inp, act)
	return cha
}

func pipeListSBuffer(out chan<- []*list.List, inp <-chan []*list.List) {
	defer close(out)
	for i := range inp {
		out <- i
	}
}

// PipeListSBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeListSBuffer(inp <-chan []*list.List, cap int) (out <-chan []*list.List) {
	cha := make(chan []*list.List, cap)
	go pipeListSBuffer(cha, inp)
	return cha
}

func pipeListSFunc(out chan<- []*list.List, inp <-chan []*list.List, act func(a []*list.List) []*list.List) {
	defer close(out)
	for i := range inp {
		out <- act(i)
	}
}

// PipeListSFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeListSMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeListSFunc(inp <-chan []*list.List, act func(a []*list.List) []*list.List) (out <-chan []*list.List) {
	cha := make(chan []*list.List)
	if act == nil {
		act = func(a []*list.List) []*list.List { return a }
	}
	go pipeListSFunc(cha, inp, act)
	return cha
}

func pipeListSFork(out1, out2 chan<- []*list.List, inp <-chan []*list.List) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		out1 <- i
		out2 <- i
	}
}

// PipeListSFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeListSFork(inp <-chan []*list.List) (out1, out2 <-chan []*list.List) {
	cha1 := make(chan []*list.List)
	cha2 := make(chan []*list.List)
	go pipeListSFork(cha1, cha2, inp)
	return cha1, cha2
}

// ListSTube is the signature for a pipe function.
type ListSTube func(inp <-chan []*list.List, out <-chan []*list.List)

// ListSDaisy returns a channel to receive all inp after having passed thru tube.
func ListSDaisy(inp <-chan []*list.List, tube ListSTube) (out <-chan []*list.List) {
	cha := make(chan []*list.List)
	go tube(inp, cha)
	return cha
}

// ListSDaisyChain returns a channel to receive all inp after having passed thru all tubes.
func ListSDaisyChain(inp <-chan []*list.List, tubes ...ListSTube) (out <-chan []*list.List) {
	cha := inp
	for i := range tubes {
		cha = ListSDaisy(cha, tubes[i])
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
