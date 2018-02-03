// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bufio

// This file was generated with dotgo
// DO NOT EDIT - Improve the pattern!

import (
	bufio "bufio"
)

// MakeSplitFuncChan returns a new open channel
// (simply a 'chan bufio.SplitFunc' that is).
//
// Note: No 'SplitFunc-producer' is launched here yet! (as is in all the other functions).
//
// This is useful to easily create corresponding variables such as
//
//	var mySplitFuncPipelineStartsHere := MakeSplitFuncChan()
//	// ... lot's of code to design and build Your favourite "mySplitFuncWorkflowPipeline"
//	// ...
//	// ... *before* You start pouring data into it, e.g. simply via:
//	for drop := range water {
//		mySplitFuncPipelineStartsHere <- drop
//	}
//	close(mySplitFuncPipelineStartsHere)
//
// Hint: especially helpful, if Your piping library operates on some hidden (non-exported) type
// (or on a type imported from elsewhere - and You don't want/need or should(!) have to care.)
//
// Note: as always (except for PipeSplitFuncBuffer) the channel is unbuffered.
//
func MakeSplitFuncChan() (out chan bufio.SplitFunc) {
	return make(chan bufio.SplitFunc)
}

func sendSplitFunc(out chan<- bufio.SplitFunc, inp ...bufio.SplitFunc) {
	defer close(out)
	for i := range inp {
		out <- inp[i]
	}
}

// ChanSplitFunc returns a channel to receive all inputs before close.
func ChanSplitFunc(inp ...bufio.SplitFunc) (out <-chan bufio.SplitFunc) {
	cha := make(chan bufio.SplitFunc)
	go sendSplitFunc(cha, inp...)
	return cha
}

func sendSplitFuncSlice(out chan<- bufio.SplitFunc, inp ...[]bufio.SplitFunc) {
	defer close(out)
	for i := range inp {
		for j := range inp[i] {
			out <- inp[i][j]
		}
	}
}

// ChanSplitFuncSlice returns a channel to receive all inputs before close.
func ChanSplitFuncSlice(inp ...[]bufio.SplitFunc) (out <-chan bufio.SplitFunc) {
	cha := make(chan bufio.SplitFunc)
	go sendSplitFuncSlice(cha, inp...)
	return cha
}

func chanSplitFuncFuncNok(out chan<- bufio.SplitFunc, act func() (bufio.SplitFunc, bool)) {
	defer close(out)
	for {
		res, ok := act() // Apply action
		if !ok {
			return
		}
		out <- res
	}
}

// ChanSplitFuncFuncNok returns a channel to receive all results of act until nok before close.
func ChanSplitFuncFuncNok(act func() (bufio.SplitFunc, bool)) (out <-chan bufio.SplitFunc) {
	cha := make(chan bufio.SplitFunc)
	go chanSplitFuncFuncNok(cha, act)
	return cha
}

func chanSplitFuncFuncErr(out chan<- bufio.SplitFunc, act func() (bufio.SplitFunc, error)) {
	defer close(out)
	for {
		res, err := act() // Apply action
		if err != nil {
			return
		}
		out <- res
	}
}

// ChanSplitFuncFuncErr returns a channel to receive all results of act until err != nil before close.
func ChanSplitFuncFuncErr(act func() (bufio.SplitFunc, error)) (out <-chan bufio.SplitFunc) {
	cha := make(chan bufio.SplitFunc)
	go chanSplitFuncFuncErr(cha, act)
	return cha
}

func joinSplitFunc(done chan<- struct{}, out chan<- bufio.SplitFunc, inp ...bufio.SplitFunc) {
	defer close(done)
	for i := range inp {
		out <- inp[i]
	}
	done <- struct{}{}
}

// JoinSplitFunc sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinSplitFunc(out chan<- bufio.SplitFunc, inp ...bufio.SplitFunc) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinSplitFunc(cha, out, inp...)
	return cha
}

func joinSplitFuncSlice(done chan<- struct{}, out chan<- bufio.SplitFunc, inp ...[]bufio.SplitFunc) {
	defer close(done)
	for i := range inp {
		for j := range inp[i] {
			out <- inp[i][j]
		}
	}
	done <- struct{}{}
}

// JoinSplitFuncSlice sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinSplitFuncSlice(out chan<- bufio.SplitFunc, inp ...[]bufio.SplitFunc) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinSplitFuncSlice(cha, out, inp...)
	return cha
}

func joinSplitFuncChan(done chan<- struct{}, out chan<- bufio.SplitFunc, inp <-chan bufio.SplitFunc) {
	defer close(done)
	for i := range inp {
		out <- i
	}
	done <- struct{}{}
}

// JoinSplitFuncChan sends inputs on the given out channel and returns a done channel to receive one signal when inp has been drained
func JoinSplitFuncChan(out chan<- bufio.SplitFunc, inp <-chan bufio.SplitFunc) (done <-chan struct{}) {
	cha := make(chan struct{})
	go joinSplitFuncChan(cha, out, inp)
	return cha
}

func doitSplitFunc(done chan<- struct{}, inp <-chan bufio.SplitFunc) {
	defer close(done)
	for i := range inp {
		_ = i // Drain inp
	}
	done <- struct{}{}
}

// DoneSplitFunc returns a channel to receive one signal before close after inp has been drained.
func DoneSplitFunc(inp <-chan bufio.SplitFunc) (done <-chan struct{}) {
	cha := make(chan struct{})
	go doitSplitFunc(cha, inp)
	return cha
}

func doitSplitFuncSlice(done chan<- ([]bufio.SplitFunc), inp <-chan bufio.SplitFunc) {
	defer close(done)
	SplitFuncS := []bufio.SplitFunc{}
	for i := range inp {
		SplitFuncS = append(SplitFuncS, i)
	}
	done <- SplitFuncS
}

// DoneSplitFuncSlice returns a channel which will receive a slice
// of all the SplitFuncs received on inp channel before close.
// Unlike DoneSplitFunc, a full slice is sent once, not just an event.
func DoneSplitFuncSlice(inp <-chan bufio.SplitFunc) (done <-chan ([]bufio.SplitFunc)) {
	cha := make(chan ([]bufio.SplitFunc))
	go doitSplitFuncSlice(cha, inp)
	return cha
}

func doitSplitFuncFunc(done chan<- struct{}, inp <-chan bufio.SplitFunc, act func(a bufio.SplitFunc)) {
	defer close(done)
	for i := range inp {
		act(i) // Apply action
	}
	done <- struct{}{}
}

// DoneSplitFuncFunc returns a channel to receive one signal before close after act has been applied to all inp.
func DoneSplitFuncFunc(inp <-chan bufio.SplitFunc, act func(a bufio.SplitFunc)) (out <-chan struct{}) {
	cha := make(chan struct{})
	if act == nil {
		act = func(a bufio.SplitFunc) { return }
	}
	go doitSplitFuncFunc(cha, inp, act)
	return cha
}

func pipeSplitFuncBuffer(out chan<- bufio.SplitFunc, inp <-chan bufio.SplitFunc) {
	defer close(out)
	for i := range inp {
		out <- i
	}
}

// PipeSplitFuncBuffer returns a buffered channel with capacity cap to receive all inp before close.
func PipeSplitFuncBuffer(inp <-chan bufio.SplitFunc, cap int) (out <-chan bufio.SplitFunc) {
	cha := make(chan bufio.SplitFunc, cap)
	go pipeSplitFuncBuffer(cha, inp)
	return cha
}

func pipeSplitFuncFunc(out chan<- bufio.SplitFunc, inp <-chan bufio.SplitFunc, act func(a bufio.SplitFunc) bufio.SplitFunc) {
	defer close(out)
	for i := range inp {
		out <- act(i)
	}
}

// PipeSplitFuncFunc returns a channel to receive every result of act applied to inp before close.
// Note: it 'could' be PipeSplitFuncMap for functional people,
// but 'map' has a very different meaning in go lang.
func PipeSplitFuncFunc(inp <-chan bufio.SplitFunc, act func(a bufio.SplitFunc) bufio.SplitFunc) (out <-chan bufio.SplitFunc) {
	cha := make(chan bufio.SplitFunc)
	if act == nil {
		act = func(a bufio.SplitFunc) bufio.SplitFunc { return a }
	}
	go pipeSplitFuncFunc(cha, inp, act)
	return cha
}

func pipeSplitFuncFork(out1, out2 chan<- bufio.SplitFunc, inp <-chan bufio.SplitFunc) {
	defer close(out1)
	defer close(out2)
	for i := range inp {
		out1 <- i
		out2 <- i
	}
}

// PipeSplitFuncFork returns two channels to receive every result of inp before close.
//  Note: Yes, it is a VERY simple fanout - but sometimes all You need.
func PipeSplitFuncFork(inp <-chan bufio.SplitFunc) (out1, out2 <-chan bufio.SplitFunc) {
	cha1 := make(chan bufio.SplitFunc)
	cha2 := make(chan bufio.SplitFunc)
	go pipeSplitFuncFork(cha1, cha2, inp)
	return cha1, cha2
}

// SplitFuncTube is the signature for a pipe function.
type SplitFuncTube func(inp <-chan bufio.SplitFunc, out <-chan bufio.SplitFunc)

// SplitFuncDaisy returns a channel to receive all inp after having passed thru tube.
func SplitFuncDaisy(inp <-chan bufio.SplitFunc, tube SplitFuncTube) (out <-chan bufio.SplitFunc) {
	cha := make(chan bufio.SplitFunc)
	go tube(inp, cha)
	return cha
}

// SplitFuncDaisyChain returns a channel to receive all inp after having passed thru all tubes.
func SplitFuncDaisyChain(inp <-chan bufio.SplitFunc, tubes ...SplitFuncTube) (out <-chan bufio.SplitFunc) {
	cha := inp
	for i := range tubes {
		cha = SplitFuncDaisy(cha, tubes[i])
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
